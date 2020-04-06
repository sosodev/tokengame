package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // enables the use of Postgresql with Gorm
	"github.com/tokengame/backend/database/models"
)

// GetDatabaseConnection returns a gorm.DB pointer or error if one occurred
func GetDatabaseConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=database port=5432 user=postgres password=TokenTournamentSecretPW dbname=postgres sslmode=disable")
}

// RunDatabaseMigrations run's GORM automigrations on the database models
// ensuring that the database table matches the struct fields
func RunDatabaseMigrations() error {
	db, err := GetDatabaseConnection()
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Challenge{}, &models.Highscore{})
	return nil
}

// SeedDatabase runs all of the models seed functions
// Returns the first error if one occurred
func SeedDatabase() error {
	db, err := GetDatabaseConnection()
	if err != nil {
		return err
	}

	err = models.SeedChallenges(db)
	if err != nil {
		return err
	}

	return nil
}
