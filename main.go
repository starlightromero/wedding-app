package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var session, _ = mgo.Dial("127.0.0.1")
var c = session.DB("WeddingDb").C("RSVP")

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/views/home.html")
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
