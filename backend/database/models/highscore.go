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
