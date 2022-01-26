package controllers

import (
	"net/http"
	"log"
	"strings"
	"wedding-app/models"

	"github.com/labstack/echo/v4"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetHealth(c echo.Context) error {
        return c.JSON(http.StatusOK, "healthy")
}

func GetRSVPs(c echo.Context) error {
	rsvps := []models.RSVP{}
	mgm.Coll(&models.RSVP{}).SimpleFind(&rsvps, bson.M{})
	return c.JSON(http.StatusOK, rsvps)
}

func PostRSVP(c echo.Context) error {
        name := c.FormValue("name")
        attending := c.FormValue("attending")

	dietNone := c.FormValue("none")
	dietVegetarian := c.FormValue("vegetarian")
	dietVegan := c.FormValue("vegan")
	var dietOther string
	if len(c.FormValue("other")) > 0 && len(c.FormValue("dietOther")) > 0 {
		dietOther = c.FormValue("dietOther")
	}

	dietArr := []string{}

	if len(dietNone) > 0 {
		dietArr = append(dietArr, dietNone)
	}

	if len(dietVegetarian) > 0 {
		dietArr = append(dietArr, dietVegetarian)
	}

	if len(dietVegan) > 0 {
		dietArr = append(dietArr, dietVegan)
	}

	if len(dietOther) > 0 {
		dietArr = append(dietArr, dietOther)
	}

	diet := strings.Join(dietArr, ", ")

	rsvp := models.NewRSVP(
                name,
                diet,
                len(attending) > 0,
        )

	err := mgm.Coll(rsvp).Create(rsvp)
	if err != nil {
		log.Fatal(err)
	}

        return c.String(http.StatusOK, "Thank you for RSVPing! name: "+name+" dietary restrictions: "+diet+" attending: "+attending)
}

