package services

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"time"

	authLogEntity "github.com/alwismt/selectify/internal/adminApp/domains/shared/entities"
	authLogRepo "github.com/alwismt/selectify/internal/adminApp/domains/shared/repositories"
	transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"
	"github.com/alwismt/selectify/internal/infrastructure/messagebroker/queue"
	mongodb "github.com/alwismt/selectify/internal/infrastructure/persistence/mongoDB"

	"github.com/google/uuid"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthLogService interface {
	NewAuthEntry(ctx context.Context, data *transferobjects.EventAuthDTO) error
}

type authLogService struct {
	db     *mongo.Database
	aLrepo authLogRepo.AuthLogRepository
	event  queue.RabbitMQPublisher
}

func NewAuthLogService(mdb *mongo.Database) AuthLogService {
	if mdb == nil {
		mdb = mongodb.MongoDB
	}
	event := queue.NewRabbitMQPublisher(nil)
	return &authLogService{db: mdb, aLrepo: authLogRepo.NewAuthLogRepository(mdb), event: event}
}

func (q *authLogService) NewAuthEntry(ctx context.Context, data *transferobjects.EventAuthDTO) error {

	// Create a channel to receive the location and ISP information
	locChan := make(chan string)
	latChan := make(chan float64)
	longChan := make(chan float64)
	ispChan := make(chan string)

	// Parse the user agent string
	uaString := data.UserAgent
	ua := user_agent.New(uaString)

	// Extract device and browser information
	browser, _ := ua.Browser()

	if data.IP == "127.0.0.1" {
		data.IP = "112.135.217.248"
		// data.IP = "168.138.104.28"
	}
	// Start a goroutine to lookup the geolocation information for the IP address
	go func() {
		db, err := geoip2.Open("./data/geoLite2/GeoLite2-City.mmdb")
		if err != nil {
			log.Println(err)
			return
		}
		defer db.Close()
		record, err := db.City(net.ParseIP(data.IP))
		if err != nil {
			log.Println(err)
			return
		}
		location := fmt.Sprintf("%s, %s", record.City.Names["en"], record.Country.Names["en"])
		latitude := record.Location.Latitude
		longitude := record.Location.Longitude
		locChan <- location
		latChan <- latitude
		longChan <- longitude
	}()

	// Start another goroutine to lookup the ISP information for the IP address
	go func() {
		ispdb, err := geoip2.Open("./data/geoLite2/GeoLite2-ASN.mmdb")
		if err != nil {
			log.Println(err)
			return
		}
		defer ispdb.Close()

		asn, err := ispdb.ASN(net.ParseIP(data.IP))
		if err != nil {
			log.Println(err)
			return
		}
		isprec := asn.AutonomousSystemOrganization
		ispChan <- isprec
	}()

	// Receive the results from the channels
	location := <-locChan
	isp := <-ispChan
	latitude := <-latChan
	longitude := <-longChan

	userId, _ := uuid.Parse(data.UserID)
	sesId, _ := uuid.Parse(data.SessionID)
	log := &authLogEntity.AuthLog{
		ID:              uuid.New(),
		UserID:          userId,
		SessionID:       sesId,
		IP:              data.IP,
		Location:        location,
		Platform:        ua.Platform(),
		OperatingSystem: ua.OS(),
		Model:           ua.Model(),
		Browser:         browser,
		UserAgent:       data.UserAgent,
		ISP:             isp,
		Latitude:        latitude,
		Longitude:       longitude,
		LoggedTime:      data.Timestamp,
		CreatedAt:       time.Now(),
	}

	// get previous data
	previousData, err := q.aLrepo.GetAuthLogs(context.Background(), userId)
	if err != nil {
		return err
	}
	if previousData == nil {
		// save to db
		if err = q.aLrepo.AddAuthLog(context.Background(), log); err != nil {
			return err
		}
		return nil
	}
	distSus, NewDevice := detectSuspicious(log, previousData)
	if distSus {
		susLogService := NewSuspiciousLoginService(nil)
		err := susLogService.NewSuspiciousLogin(ctx, log)
		if err != nil {
			return err
		}
		// log.Suspicious = true
		// log.SuspiciousReason = "Suspicious login location"
		fmt.Println("Suspicious login location")
		return nil
	}
	if NewDevice {
		email := &transferobjects.EmailDTO{
			Name: "selectify_email",
			Type: "newDevice",
			Data: log,
		}

		err = q.event.QueueEvent(email)
		if err != nil {
			return err
		}

		// log.Suspicious = true
		// log.SuspiciousReason = "New device"
		fmt.Println("New device")
	}

	// save to db
	if err = q.aLrepo.AddAuthLog(context.Background(), log); err != nil {
		return err
	}
	return nil
}

func detectSuspicious(newLoginData *authLogEntity.AuthLog, previousData []authLogEntity.AuthLog) (bool, bool) {
	const threshold = 500.0 // threshold distance in kilometers
	closestDist := math.MaxFloat64
	newDevice := true
	for _, data := range previousData {
		dist := distance(newLoginData, &data)
		if dist < closestDist {
			closestDist = dist
		}
		device := oldDevice(newLoginData, &data)
		if device {
			newDevice = false
		}

	}
	distSus := closestDist > threshold
	if distSus {
		return distSus, newDevice
	}
	if newDevice {
		return distSus, newDevice
	}
	return false, false
}

func oldDevice(newLoginData *authLogEntity.AuthLog, previousData *authLogEntity.AuthLog) bool {
	if newLoginData.Browser == previousData.Browser && newLoginData.OperatingSystem == previousData.OperatingSystem && newLoginData.Platform == previousData.Platform && newLoginData.Model == previousData.Model {
		return true
	}
	return false
}

// calculates the great-circle distance between two points using the Haversine formula
func distance(p1, p2 *authLogEntity.AuthLog) float64 {
	const radius = 6371.0 // Earth's radius in kilometers
	lat1, lon1 := toRadians(p1.Latitude), toRadians(p1.Longitude)
	lat2, lon2 := toRadians(p2.Latitude), toRadians(p2.Longitude)
	dlat, dlon := lat2-lat1, lon2-lon1
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radius * c
}

// converts degrees to radians
func toRadians(deg float64) float64 {
	return deg * math.Pi / 180
}
