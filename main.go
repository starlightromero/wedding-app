package main

import (
	"log"
	"net/http"

	//	"github.com/globalsign/mgo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//var session, _ = mgo.Dial("127.0.0.1")
//var c = session.DB("WeddingDb").C("RSVP")

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	data, err := Asset("public/views/home.html")
	if err != nil {
		log.Fatal(err)
	}

	e.File("/", string(data))
	e.POST("/", postRSVP)
	e.GET("/health", getHealth)

	e.Logger.Fatal(e.Start(":8080"))
}

func getHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}

func postRSVP(c echo.Context) error {
	name := c.FormValue("name")
	dietaryRestrictions := c.FormValue("dietaryRestrictions")
	attending := c.FormValue("attending")
	return c.String(http.StatusOK, "Thank you for RSVPing! name: "+name+"dietary restrictions:  "+dietaryRestrictions+"attending: "+attending)
}
