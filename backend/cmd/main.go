package main

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
        DBName: "food-test",
        DBUser: "postgres",
        DBPassword: "123321",
        DBHost: "localhost",
        DBPort: "5436",
        SERVERHost: "localhost",
        SERVERPort: "8080",
    }



    return;
}