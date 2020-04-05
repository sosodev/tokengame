package highscore

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tokengame/backend/database"
	"github.com/tokengame/backend/database/models"
)

// Highscores for multiple highscore JSON response
type Highscores struct {
	Highscores []models.Highscore `json:"highscores"`
}

// GetHighscores returns the highscores
func GetHighscores(c echo.Context) error {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	var highscores []models.Highscore
	if err = db.Find(&highscores).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Highscores{
		Highscores: highscores,
	})
}

// PostHighscore posts a highscore to the highscores table
// Takes the nickname, score, and challenge ID as parameters
func PostHighscore(c echo.Context) error {
	// bind the highscore JSON data to the struct
	highscore := &models.Highscore{}
	if err := c.Bind(highscore); err != nil {
		return err
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Store it
	if err = db.Create(highscore).Error; err != nil {
		return err
	}

	// Return the new data as JSON for verification
	return c.JSON(http.StatusOK, *highscore)
}
