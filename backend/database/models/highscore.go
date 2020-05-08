package models

import (
	"github.com/jinzhu/gorm"
)

// Highscore database model
type Highscore struct {
	gorm.Model
	Nickname    string `json:"nickname"`
	Score       int    `json:"score"`
	ChallengeID uint   `json:"challenge_id"`
}

// SeedHighscores seeds the highscores table
func SeedHighscores(db *gorm.DB) error {
	highscore := &Highscore{}

	err := db.First(highscore).Error
	if !gorm.IsRecordNotFoundError(err) {
		return err
	}

	if highscore.Nickname != "" {
		return nil
	}

	db.Create(&Highscore{
		Nickname:    "Cool Guy",
		Score:       8420,
		ChallengeID: 2,
	})
	if db.Error != nil {
		return db.Error
	}

	db.Create(&Highscore{
		Nickname:    "Chet",
		Score:       5356,
		ChallengeID: 1,
	})
	if db.Error != nil {
		return db.Error
	}

	db.Create(&Highscore{
		Nickname:    "Chet but fast",
		Score:       7560,
		ChallengeID: 1,
	})
	if db.Error != nil {
		return db.Error
	}

	return db.Error
}
