package handlers

import (
	"fmt"
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

func GetAllRSVPs(c echo.Context) error {
	rsvps := []models.RSVP{}
	mgm.Coll(&models.RSVP{}).SimpleFind(&rsvps, bson.M{})
	return c.JSON(http.StatusOK, rsvps)
}

func GetOneRSVP(c echo.Context) error {
	id := c.Param("id")
	rsvp := &models.RSVP{}
	mgm.Coll(rsvp).FindByID(id, rsvp)
	return c.JSON(http.StatusOK, rsvp)
}

func CreateRSVP(c echo.Context) error {
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

func UpdateOneRSVP(c echo.Context) error {
	id := c.Param("id")
        rsvp := &models.RSVP{}
        mgm.Coll(rsvp).FindByID(id, rsvp)

	// rsvp.Name = "Moulin Rouge!"
	err := mgm.Coll(rsvp).Update(rsvp)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, "")
}

func DeleteOneRSVP(c echo.Context) error {
	id := c.Param("id")
        rsvp := &models.RSVP{}
        mgm.Coll(rsvp).FindByID(id, rsvp)

	err := mgm.Coll(rsvp).Delete(rsvp)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, rsvp)
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("/public/views/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
}
