package app

import (
	"os"
	"log"
	"wedding-app/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/kamva/mgm/v3"
        "go.mongodb.org/mongo-driver/mongo/options"
)

func Run() {
	err := mgm.SetDefaultConfig(nil, "admin", options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "/public/views")
	e.POST("/", controllers.PostRSVP)
	e.GET("/rsvp", controllers.GetRSVPs)
	e.GET("/health", controllers.GetHealth)

	e.Logger.Fatal(e.Start(":8080"))
}

