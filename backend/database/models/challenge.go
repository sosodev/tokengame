package models

import "github.com/jinzhu/gorm"

// Challenge database model
type Challenge struct {
	gorm.Model
	Title string `json:"title"`
}
