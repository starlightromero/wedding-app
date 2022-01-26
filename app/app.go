package app

import (
	"os"
	"log"
	"wedding-app/handlers"

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
	e.POST("/", handlers.CreateRSVP)
	e.GET("/rsvp", handlers.GetAllRSVPs)
	e.GET("/rsvp/:id", handlers.GetOneRSVP)
	e.Static("/rsvp/update", "/public/views/update.html")
	e.PATCH("/rsvp/:id", handlers.UpdateOneRSVP)
	e.Static("/rsvp/delete", "/public/views/delete.html")
	e.DELETE("/rsvp/:id", handlers.DeleteOneRSVP)
	e.GET("/health", handlers.GetHealth)

	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	e.Logger.Fatal(e.Start(":8080"))
}

