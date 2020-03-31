package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tokengame/backend/database"
)

/*
	Some useful documentation links :)

	Go: https://golang.org/doc/
	Echo (backend framework): https://echo.labstack.com/guide
	GORM (simple ORM for Go): https://gorm.io/docs/
*/

func main() {
	// create the server object
	e := echo.New()

	// run the database migrations
	// fail immediately on error
	err := database.RunDatabaseMigrations()
	if err != nil {
		e.Logger.Fatal(err)
	}

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
