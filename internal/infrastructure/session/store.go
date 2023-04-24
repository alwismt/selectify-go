package session

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis"
	"github.com/joho/godotenv"
)

var store = session.New(session.Config{
	Expiration:     3 * time.Hour,
	Storage:        NewStore(),
	KeyLookup:      "cookie:session",
	CookieHTTPOnly: true,
	KeyGenerator:   utils.UUIDv4,
})

func NewStore() *redis.Storage {

	// Load env: since global variable "middleware/session" initialized before the main function so it won't read the env from main
	err := godotenv.Load("./.env")
	if err != nil {
		panic("error loading .env file " + err.Error())
	}

	var port int
	var dbIndex int = 1

	redisPort := os.Getenv("REDIS_PORT")

	if redisPort == "" {
		panic("REDIS is not set")
	} else {
		port, err = strconv.Atoi(redisPort)
		if err != nil {
			panic("REDIS is not set between 1-15")
		}

	}
	cachedb := os.Getenv("REDIS_INDEX_ADMIN")
	if cachedb == "" {
		panic("REDIS_INDEX_ADMIN is not set")
	} else {
		// convert cachedb to int
		dbIndex, err = strconv.Atoi(cachedb)
		if err != nil {
			panic("REDIS_INDEX_ADMIN is not set between 1-15")
		}

	}
	var (
		redisHost     = os.Getenv("REDIS_HOST")
		redisPassword = os.Getenv("REDIS_PASSWORD")
		redisUsername = os.Getenv("REDIS_USERNAME")
	)

	return redis.New(redis.Config{
		Host:      redisHost,
		Port:      port,
		Username:  redisUsername,
		Password:  redisPassword,
		Database:  dbIndex,
		Reset:     false,
		TLSConfig: nil,
	})
}
