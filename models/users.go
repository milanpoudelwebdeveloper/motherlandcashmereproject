package models

import (
	"github.com/jinzhu/gorm"
	//This is
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//User is
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
