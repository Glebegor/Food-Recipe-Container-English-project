package main

import (
	"fmt"
	"time"

	"backend/api/routers"
	"backend/domain"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for testing
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	conf := domain.Config{
		DBName:     "postgres",
		DBUser:     "postgres",
		DBPassword: "123321",
		DBHost:     "localhost",
		DBPort:     "5432",
		SERVERHost: "localhost",
		SERVERPort: "8080",
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", conf.DBUser, conf.DBPassword, conf.DBName, conf.DBHost, conf.DBPort))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to connect to database")
	}

	if err = db.Ping(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to ping database")
	}

	router := gin.Default()
	router.Use(CORS())
	timeout := 5 * time.Second

	routers.SetupRouter(router, db, conf, timeout)
	router.Run(fmt.Sprintf("%s:%s", conf.SERVERHost, conf.SERVERPort))
}
