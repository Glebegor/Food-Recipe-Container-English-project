package main
import (
    "fmt"

    "backend/api/routers"
    "github.com/sirupsen/logrus"
    "github.com/jmoiron/sqlx"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)
type Config struct {
    DBName string
    DBUser string
    DBPassword string
    DBHost string
    DBPort string
    SERVERPort string
    SERVERHost string
}


func main() {
    logrus.SetFormatter(&logrus.JSONFormatter{})

    // Config
    conf := Config{
        DBName: "postgres",
        DBUser: "postgres",
        DBPassword: "123321",
        DBHost: "localhost",
        DBPort: "5432",
        SERVERHost: "localhost",
        SERVERPort: "8080",
    }

    // Connect to database
    var db *sqlx.DB
    db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", conf.DBUser, conf.DBPassword, conf.DBName, conf.DBHost, conf.DBPort))
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "error": err,
        }).Fatal("Failed to connect to database")
    }

   // Ping
    if err = db.Ping(); err != nil {
        logrus.WithFields(logrus.Fields{
            "error": err,
        }).Fatal("Failed to ping database")
    }

    gin := gin.Default()
    timeout := 5 * time.Second

    routers.SetupRouter(gin, db, conf)
    git.Run(fmt.Sprintf("%s:%s", conf.SERVERHost, conf.SERVERPort))
    return;
}