package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	//This is
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	//ErrNotFound is returned when resouce can't be found in database
	ErrNotFound = errors.New("models: resource not found")
)

//NewUserService is
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

//UserService is
type UserService struct {
	db *gorm.DB
}

//ByID will look up by the id provided
//case 1-user,error-nil
//case 2-nil, ErrNotFound
//case 3-nil,otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id=?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err

	}

}

//DestrutiveReset drops a table and rebuilds it
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

//Close closes the user service database connection
func (us *UserService) Close() error {
	return us.db.Close()

}

//User is
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
