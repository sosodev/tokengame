package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tokengame/backend/challenge"
	"github.com/tokengame/backend/database"
	"github.com/tokengame/backend/highscore"
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

	// seed the database with starting data
	err = database.SeedDatabase()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// create an api route group
	api := e.Group("api")
	api.GET("/challenges", challenge.GetChallenges) // routes "/api/challenges" to our challenge.GetChallenges handler
	api.GET("/challenges/:id", challenge.GetChallenge)
	api.GET("/highscores", highscore.GetHighscores)
	api.POST("/highscores/new", highscore.PostHighscore)

	// create the admin api route subgroup
	admin := e.Group("admin")
	// requests can only succeed when the user inputs the correct username and password
	admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "beargang" {
			return true, nil
		}

		return false, nil
	}))
	admin.POST("/challenges", challenge.CreateChallenge) // POST /admin/challenges
	admin.DELETE("/highscores/:id", highscore.RemoveHighscore)
	admin.DELETE("/challenges/:id", challenge.DeleteChallenge)

	// enable static asset serving on the server
	// but also route 404's to index.html
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		Index: "index.html",
		HTML5: true,
	}))

	// start the server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
