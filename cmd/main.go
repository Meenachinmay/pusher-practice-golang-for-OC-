package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pusher/pusher-http-go/v5"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"pusher-practice/config"
	"pusher-practice/handlers"
	"pusher-practice/internal/database"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}

	// initialize the redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Failed to load dburl")
		return
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect with database")
		return
	}

	var testQuery int
	err = conn.QueryRow("SELECT 1").Scan(&testQuery)
	if err != nil {
		log.Fatal("Failed on the test query, check database connection")
	} else {
		log.Println("Connection test query passed. Connection is stable")
	}

	// Initialize pusher here
	pusherClient := &pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: "ap3",
		Secure:  true,
	}

	apiConfig := &config.ApiConfig{
		DB:           database.New(conn),
		RedisClient:  redisClient,
		PusherClient: pusherClient,
	}

	localApiConfig := &handlers.LocalApiConfig{
		apiConfig,
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/healthcheck", localApiConfig.HandlerReadiness)
	router.GET("/check-ws", localApiConfig.HandlerWs)
	router.POST("/send-message", localApiConfig.HandlersSendMessage)

	log.Fatal(router.Run(":8080"))
}
