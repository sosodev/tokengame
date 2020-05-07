package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Challenge database model
type Challenge struct {
	gorm.Model
	Title      string         `json:"title"`                                 // Title of the challenge
	Head       string         `json:"head"`                                  // Initial code block before the user's code (e.g. function declaration)
	Foot       string         `json:"foot"`                                  // Code block after the user's code (e.g. return statement)
	Tokens     pq.StringArray `json:"tokens" gorm:"type:varchar(1000)[]"`    // The pool of tokens for the challenge
	Testcases  pq.StringArray `json:"testcases" gorm:"type:varchar(1000)[]"` // A set of tests, as functions, to be run after the user's code, should return a boolean indicating success
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
			")",
			"1",
			"}",
			"{",
			"<",
			">",
			";",
			"=",
		},
		Testcases: []string{
			"function test1() { return JSON.stringify(bubble_sort([2, 3, 1])) === JSON.stringify([1, 2, 3]) }; test1();",
		},
	})
	if db.Error != nil {
		return err
	}

	db.Create(&Challenge{
		Title: "Fibonacci Sequence",
		Head: `/* Write an algorithm that will generate the fibonacci numbers up to the nth number in the sequence.

ie. 0, 1, 1, 2, 3, 5, 8, ..., n.

Return the n index number, fib(4) should return 3
*/function fib(n) {
`,
		Foot: "}",
		Tokens: []string{
			"return current",
			"return n;",
			"return ",
			"n >= 0",
			"n <= 1",
			"while(",
			"next = 1",
			"current = 0",
			"var",
			"temp",
			"next",
			"current",
			"if(",
			"fib(",
			"--n",
			"n--",
			"n",
			";",
			")",
			"{",
			"}",
			"=",
			"+",
			"-1",
			"-2",
		},
		Testcases: []string{
			"function test1() { return fib(4) === 3};test1();",
		},
	})
	if db.Error != nil {
		return err
	}

	db.Create(&Challenge{
		Title: "FizzBuzz",
		Head: `/*Write a program that adds to the list the numbers from 1 to 100. 
		But for multiples of three add "fizz" instead of the number 
		and for the multiples of five add "buzz". 
		For numbers which are multiples of both three and five add "fizzbuzz".*/
		function fizzBuzz(out) {`,
		Foot: "return out;\n}",
		Tokens: []string{
			"var i = 1;",
			"for(;",
			"100",
			"99",
			"<",
			";",
			"i++",
			"<=",
			"if(",
			"%",
			"==",
			"&&",
			"5",
			"3",
			")",
			"else",
			"i",
			"out.push(\"fizzbuzz\");",
			"out.push(\"fizz\");",
			"out.push(\"buzz\");",
			"out.push(i);",
			"{",
			"}",
		},
		Testcases: []string{
			"function test1() {var out=[];var i = 1;for (;i<=100;i++){if (i % 5 == 0 && i % 3 == 0){out.push(\"fizzbuzz\");}if (i % 3 == 0){out.push(\"fizz\");}else if (i % 5 == 0){out.push(\"buzz\");}else{out.push(i);}}var user_array =  fizzBuzz([]);if(user_array.length !== 99){return false;}else{for (i = 0; i < 100; i++){if(user_array[i] != out[i]){return false;}}return true;} }test1();",
		},
	})

	return db.Error
}
