package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alwismt/selectify/internal/adminApp/domains/Auth/entities"
	"github.com/alwismt/selectify/internal/infrastructure/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func ConnectToMongoDB() {
	dsn, err := utils.ConnectionURLBuilder("mongo")
	if err != nil {
		// do something later
		fmt.Println("Failed to env")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI(dsn).SetMaxPoolSize(1000))
	if err != nil {
		panic(err)
	}

	// connect to the MongoDB server
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	// MongoClient = client

	database := os.Getenv("MONGODB_NAME")
	if database == "" {
		panic("MONGODB_NAME is not set")
	}
	MongoDB = client.Database(database)

	addCollectionIndexes()
	admin := os.Getenv("ADMIN_SEED")
	if admin == "true" {
		seedAdmin()
	}
}

func addCollectionIndexes() {
	adminCollectionName := "admins"
	adminCollectionExists, err := collectionExists(MongoDB, adminCollectionName)
	if err != nil {
		log.Fatal(err)
	}

	if !adminCollectionExists {

		err = MongoDB.CreateCollection(context.Background(), adminCollectionName)
		if err != nil {
			log.Fatal(err)
		}
	}

	adminCollection := MongoDB.Collection(adminCollectionName)
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err = adminCollection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
}

func collectionExists(db *mongo.Database, collectionName string) (bool, error) {
	collections, err := db.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		return false, err
	}
	for _, collection := range collections {
		if collection == collectionName {
			return true, nil
		}
	}
	return false, nil
}

func seedAdmin() {
	// Connect to mongo Database *mongo.Database
	db := MongoDB
	name := os.Getenv("ADMIN_NAME")
	if name == "" {
		name = "Super Admin"
	}
	email := os.Getenv("ADMIN_EMAIL")
	if email == "" {
		panic("ADMIN_EMAIL is not set")
	}
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		panic("ADMIN_PASSWORD is not set")
	}

	seed := entities.Admin{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  utils.GeneratePassword(password),
		Type:      1,
		Status:    1,
		LastLogin: time.Time{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// Check if admin exists
	adminCollection := db.Collection("admins")
	filter := bson.M{"email": email}
	var admin entities.Admin
	err := adminCollection.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Insert admin
			_, err := adminCollection.InsertOne(context.Background(), seed)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			// write .env file to ADMIN_SEED=false
			if err := utils.UpdateEnvFile("ADMIN_SEED", "false"); err != nil {
				log.Fatalln(err)
			}
		}
	}
}
