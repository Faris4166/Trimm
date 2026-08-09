package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ShorthlyLink struct {
	gorm.Model
	OriginalURL string `gorm:"unique"`
	ShortURL    string `gorm:"unique"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment instead")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword +
		" dbname=" + dbDatabase + " port=" + dbPort + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&ShorthlyLink{})

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/shorten", func(c *gin.Context) {
		var data struct {
			URL string `json:"url" binding:"required"`
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var link ShorthlyLink
		result := db.Where("original_url = ?", data.URL).First(&link)

		if result.Error != nil {
			if result.Error != gorm.ErrRecordNotFound {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			link = ShorthlyLink{OriginalURL: data.URL, ShortURL: generateShortURL()}
			if err := db.Create(&link).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"short_url": link.ShortURL})
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		shortURL := c.Param("shortURL")
		var link ShorthlyLink
		result := db.Where("short_url = ?", shortURL).First(&link)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			}
			return
		}
		c.Redirect(http.StatusMovedPermanently, link.OriginalURL)
	})

	r.Run(":8080")
}

func generateShortURL() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	var shortURL string
	for i := 0; i < 6; i++ {
		shortURL += string(chars[rand.Intn(len(chars))])
	}

	return shortURL
}
