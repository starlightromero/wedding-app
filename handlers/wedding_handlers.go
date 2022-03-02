package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"wedding-app/models"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
)

func GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
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
