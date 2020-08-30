package models

import (
	"errors"

	"github.com/jinzhu/gorm"

	"../hash"
	"../rand"

	//This is
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	//ErrNotFound is returned when resouce can't be found in database
	ErrNotFound = errors.New("models: resource not found")

	//ErrInvalidID is returned when an invalid ID is provided
	//to a method like Delete.
	ErrInvalidID = errors.New("models:ID must be > 0")

	//ErrInvalidEmail is
	ErrInvalidEmail = errors.New("models:invalid email address provided")

	//ErrInvalidPassword is returned when an invalid password is used when attempting to authenticate a user.
	ErrInvalidPassword = errors.New("models: incorrect password provided")
)

const userPwPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

//NewUserService is
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	hmac := hash.NewHMAC(hmacSecretKey)
	return &UserService{
		db:   db,
		hmac: hmac,
	}, nil
}

//UserService is
type UserService struct {
	db   *gorm.DB
	hmac hash.HMAC
}

//ByID will look up by the id provided
//case 1-user,error-nil
//case 2-nil, ErrNotFound
//case 3-nil,otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id=?", id)
	err := first(db, &user)
	return &user, err
}

//ByEmail looks up a user with the given email address and
//returns that user
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email=?", email)
	err := first(db, &user)
	return &user, err

}

//ByRemember looks up a user with the given remember token and returns that user.This method will handle hashing the token for us.Errors are the same as ByEmail and BYID
func (us *UserService) ByRemember(token string) (*User, error) {
	var user User
	rememberHash := us.hmac.HASH(token)
	err := first(us.db.Where("remember_hash=?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

//Authenticate can be used to authenticate the user with the provided email address and the password.
func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))

	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		default:
			return nil, err
		}

	}

	return foundUser, nil

}

//First will query using the provided gorm.DB and it will
//get the first item returned and place it into dst,
//if nothing is found in the query it will return ErrNotFound

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err

}

//Create will create the provided user and backfill data
//like the id,created at,updated at,deleted at
func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
	}
	user.RememberHash = us.hmac.HASH(user.Remember)

	return us.db.Create(user).Error

}

//Update will update the provided user with all of the data
//in the provided user object
func (us *UserService) Update(user *User) error {
	if user.Remember != "" {
		user.RememberHash = us.hmac.HASH(user.Remember)
	}
	return us.db.Save(user).Error

}

//Delete will delete the user with the provided id
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error

}

//Close closes the user service database connection
func (us *UserService) Close() error {
	return us.db.Close()

}

//DestructiveReset drops a table and rebuilds it
func (us *UserService) DestructiveReset() error {
	if err := us.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return us.AutoMigrate()
}

//AutoMigrate will attempt to automatically migrate the
//users table
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil

}

//User is
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}
