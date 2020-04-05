package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Challenge database model
type Challenge struct {
	gorm.Model
	Title      string         `json:"title"`                                // Title of the challenge
	Head       string         `json:"head"`                                 // Initial code block before the user's code (e.g. function declaration)
	Foot       string         `json:"foot"`                                 // Code block after the user's code (e.g. return statement)
	Tokens     pq.StringArray `json:"tokens" gorm:"type:varchar(100)[]"`    // The pool of tokens for the challenge
	Testcases  pq.StringArray `json:"testcases" gorm:"type:varchar(100)[]"` // A set of tests, as functions, to be run after the user's code, should return a boolean indicating success
	Highscores []Highscore    `json:"highscores"`
}

// SeedChallenges seeds the challenges table with some basic data
func SeedChallenges(db *gorm.DB) error {
	// see if any challenges exist already
	firstChallenge := &Challenge{}

	err := db.First(firstChallenge).Error
	if !gorm.IsRecordNotFoundError(err) {
		return err // return the error if it's something other than record not found
	}
	if firstChallenge.Title != "" {
		return nil // end execution if there is already challenge data
	}

	// Create the seed data
	db.Create(&Challenge{
		Title: "Bubble Sort",
		Head:  "/* sort the given integer array (arr) into ascending order */function bubble_sort(arr) {",
		Foot:  "return arr; }",
		Tokens: []string{
			"var len =",
			"var i =",
			"var j =",
			"arr.length",
			"i <=",
			"= temp;",
			"for (",
			"j < len",
			"if (",
			"arr[j]",
			"arr[j + 1]",
			"var temp =",
			"j",
			"0",
			"len",
			"i",
			"++",
			"-",
			"j",
			")",
			"1",
			"}",
			"{",
			">",
			";",
			"=",
		},
		Testcases: []string{
			"function test1() { return JSON.stringify(bubble_sort([2, 3, 1])) === JSON.stringify([1, 2, 3]) }",
		},
	})

	return db.Error
}
