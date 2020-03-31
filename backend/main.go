package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tokengame/backend/database"
)

func main() {
	// run the database migrations
	// fail immediately on error
	err := database.RunDatabaseMigrations()
	if err != nil {
		log.Fatal(err)
	}

	// create the server object
	e := echo.New()

	// enable static asset serving on the server
	// but also route 404's to index.html
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		Index: "index.html",
		HTML5: true,
	}))

	// start the server
	e.Logger.Fatal(e.Start(":8080"))
}
