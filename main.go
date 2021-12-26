package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Meepogeroi/VacancyProject/db"
	"github.com/Meepogeroi/VacancyProject/handlers/User"
	vacancy "github.com/Meepogeroi/VacancyProject/handlers/Vacancy"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = *flag.String("port", "6001", "server port")
	}

	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()
	conn := db.Connect()
	apiPublic := echo.New()

	apiPublic.GET("/", home)

	//User
	apiPublic.POST("/user/create", user.CreateUser(conn))
	apiPublic.GET("/users", user.GetAll(conn))
	apiPublic.GET("/user", user.GetUser(conn))

	//Vacancy
	apiPublic.POST("/vac/create", vacancy.Create(conn))
	apiPublic.GET("/vacs", vacancy.GetAll(conn))
	apiPublic.GET("/vac", vacancy.GetVac(conn))

	apiPublic.Logger.Fatal(apiPublic.Start(":" + port))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Ilim Backend\nVersion: 0.0.1\n")
}
