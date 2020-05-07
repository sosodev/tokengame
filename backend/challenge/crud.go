package challenge

import (
	"net/http"
	"strconv"

	"github.com/tokengame/backend/database/models"

	"github.com/labstack/echo/v4"
	"github.com/tokengame/backend/database"
)

// Challenges is a struct to make the JSON more standard
// So the top level is an object
type Challenges struct {
	Challenges []models.Challenge `json:"challenges"`
}

// GetChallenges is an echo request handler
// That returns all of the challenges as a JSON response
func GetChallenges(c echo.Context) error {
	// connect to the database
	db, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close() // defer closing the database until the function returns

	// query for the challenges
	var challenges []models.Challenge
	if err = db.Find(&challenges).Error; err != nil {
		return err
	}

	// return the JSON of a complete Challenges struct
	return c.JSON(http.StatusOK, Challenges{
		Challenges: challenges,
	})
}

// GetChallenge returns json for a single challenge
// The challenge is found by the id route parameter
func GetChallenge(c echo.Context) error {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// convert the id param to an integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	challenge := &models.Challenge{}
	// Get the first challenge with the matching ID
	if err = db.Preload("Highscores").First(challenge, id).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, *challenge)
}

// CreateChallenge let's you create a challenge via the API
func CreateChallenge(c echo.Context) error {
	challenge := &models.Challenge{}
	if err := c.Bind(challenge); err != nil {
		return err
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Create(challenge)
	if db.Error != nil {
		return db.Error
	}

	return c.JSON(http.StatusOK, *challenge)
}

// DeleteChallenge let's you delete a challenge by ID
func DeleteChallenge(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// load the challenge by ID
	challenge := &models.Challenge{}
	if err = db.First(challenge, id).Error; err != nil {
		return err
	}
	// delete it
	if err = db.Delete(challenge).Error; err != nil {
		return err
	}

	return c.String(http.StatusOK, "")
}
