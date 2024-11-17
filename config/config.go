// config/config.go
package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"donasiPohon/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var DB *gorm.DB
var Gemini *genai.Client

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database! Error: %v", err)
	}

	// Perform automatic migrations
	err = DB.AutoMigrate(&models.User{}, &models.Komunitas{}, &models.Campaign{}, &models.Donation{})
	if err != nil {
		log.Fatalf("Failed to migrate database! Error: %v", err)
	}
}

func InitGemini() {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatalf("error when initiate gemini client: %v", err)
	}

	log.Println("gemini client connected")
	Gemini = client
}
