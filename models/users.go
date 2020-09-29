package models

import (
	"errors"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"

	"../hash"

	"../rand"
	//This is
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	//ErrNotFound is returned when resouce can't be found in database
	ErrNotFound = errors.New("models: resourcess not found")

	//ErrInvalidID is returned when an invalid ID is provided
	//to a method like Delete.
	ErrInvalidID = errors.New("models:ID must be > 0")

	//ErrInvalidEmail is
	ErrInvalidEmail = errors.New("models:invalid email address provided")

	//ErrInvalidPassword is returned when an invalid password is used when attempting to authenticate a user.
	ErrInvalidPassword = errors.New("models: incorrect password provided")

	//ErrEmailRequired is returned when an email address is not provided when creating a user
	ErrEmailRequired = errors.New("models:Email address is required")

	//ErrEmailInvalid is returned when an email address provided doesn't meet or match any of our requirements.
	ErrEmailInvalid = errors.New("models:Email address is not valid")

	//ErrEmailTaken is returned when an update or create is attempted with the email address that is already in use
	ErrEmailTaken = errors.New("models:Email address is already taken")
)

const userPwPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

//User represents the user model stored in our database.This is used for user accounts for storing both an email address and a password so users can log in and again access to their accounts.
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

//UserDB is used to interact with the users database.For single users, any error but ErrNotFound should probably result in 500 error.
type UserDB interface {
	//Methods for Querying for Single User
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	//Methods for altering users
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	//Used to close a DB connection
	Close() error

	//Migration helpers
	AutoMigrate() error
	DestructiveReset() error
}

//UserService is a set of methods used to manipulate and work with the user model
type UserService interface {
	//Authenticate will verify the provided email address and password.If they are correct,the user corresponding to that email will be returned otherwise error will be returned.
	Authenticate(email, password string) (*User, error)
	UserDB
}

//NewUserService is
func NewUserService(connectionInfo string) (UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}
	hmac := hash.NewHMAC(hmacSecretKey)
	uv := newUserValidator(ug, hmac)

	return &userService{
		UserDB: uv,
	}, nil
}

var _ UserService = &userService{}

//UserService is
type userService struct {
	UserDB
}

//Authenticate can be used to authenticate the user with the provided email address and the password.
func (us *userService) Authenticate(email, password string) (*User, error) {
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

type userValidatorFunc func(*User) error

func runUserValFuncs(user *User, fns ...userValidatorFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

var _ UserDB = &userValidator{}

func newUserValidator(udb UserDB, hmac hash.HMAC) *userValidator {
	return &userValidator{
		UserDB:     udb,
		hmac:       hmac,
		emailRegex: regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2.16}$`),
	}
}

type userValidator struct {
	UserDB
	hmac       hash.HMAC
	emailRegex *regexp.Regexp
}

//ByEmail will normalize the email address before calling ByEmail on UserDB field.
func (uv *userValidator) ByEmail(email string) (*User, error) {
	user := User{
		Email: email,
	}
	if err := runUserValFuncs(&user, uv.normalizeEmail); err != nil {
		return nil, err
	}
	return uv.UserDB.ByEmail(user.Email)
}

//ByRemember will hash the remember token and then call ByRemember on the subsequent UserDB layer.
func (uv *userValidator) ByRemember(token string) (*User, error) {
	user := User{
		Remember: token,
	}
	if err := runUserValFuncs(&user, uv.hmacRemember); err != nil {
		return nil, err
	}
	return uv.UserDB.ByRemember(user.RememberHash)

}

//Create will create the provided user and backfill data
//like the id,created at,updated at,deleted at
func (uv *userValidator) Create(user *User) error {
	err := runUserValFuncs(user,
		uv.bcryptPassword,
		uv.setRememberIfUnset,
		uv.hmacRemember,
		uv.normalizeEmail,
		uv.requireEmail,
		uv.emailFormat,
		uv.emailIsAvail)
	if err != nil {
		return err
	}
	return uv.UserDB.Create(user)
}

//Update will hash a remember token if it is provided.
func (uv *userValidator) Update(user *User) error {
	err := runUserValFuncs(user,
		uv.bcryptPassword,
		uv.hmacRemember,
		uv.normalizeEmail,
		uv.requireEmail,
		uv.emailFormat,
		uv.emailIsAvail)
	if err != nil {
		return err
	}

	return uv.UserDB.Update(user)
}

//Delete will delete the user with the provided id
func (uv *userValidator) Delete(id uint) error {
	var user User
	user.ID = id
	err := runUserValFuncs(&user, uv.idGreaterThan(0))
	if err != nil {
		return err
	}
	return uv.UserDB.Delete(id)
}

//bcryptPassword will hash the user password with a predefined pepper(userPwPepper) and bcrypt if the password field is not the empty string.
func (uv *userValidator) bcryptPassword(user *User) error {
	if user.Password == "" {
		return nil
	}
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return nil
}

func (uv *userValidator) hmacRemember(user *User) error {
	if user.Remember == "" {
		return nil
	}
	user.RememberHash = uv.hmac.HASH(user.Remember)
	return nil

}

func (uv *userValidator) setRememberIfUnset(user *User) error {
	if user.Remember != "" {
		return nil
	}

	token, err := rand.RememberToken()
	if err != nil {
		return err
	}
	user.Remember = token
	return nil

}

func (uv *userValidator) idGreaterThan(n uint) userValidatorFunc {
	return userValidatorFunc(func(user *User) error {
		if user.ID <= n {
			return ErrInvalidID
		}
		return nil
	})
}

func (uv *userValidator) normalizeEmail(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	return nil

}

func (uv *userValidator) requireEmail(user *User) error {
	if user.Email == "" {
		return ErrEmailRequired

	}
	return nil

}

func (uv *userValidator) emailFormat(user *User) error {
	if user.Email == "" {
		return nil
	}
	if !uv.emailRegex.MatchString(user.Email) {
		return ErrEmailInvalid

	}
	return nil
}

func (uv *userValidator) emailIsAvail(user *User) error {
	existing, err := uv.ByEmail(user.Email)
	if err == ErrNotFound {
		//Email address is not taken
		return nil
	}
	if err != nil {
		return err

	}
	//We found a user with this email address
	//If the found user has the same ID as this user, it is
	//an update and this is the same user

	if user.ID != existing.ID {
		return ErrEmailTaken
	}
	return nil
}

var _ UserDB = &userGorm{}

func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &userGorm{
		db: db,
	}, nil
}

type userGorm struct {
	db *gorm.DB
}

//ByID will look up by the id provided
//case 1-user,error-nil
//case 2-nil, ErrNotFound
//case 3-nil,otherError
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id=?", id)
	err := first(db, &user)
	return &user, err
}

//ByEmail looks up a user with the given email address and
//returns that user
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email=?", email)
	err := first(db, &user)
	return &user, err

}

//ByRemember looks up a user with the given remember token and returns that user.This method expects remember token to be hashed already.Errors are same as ByEmail
func (ug *userGorm) ByRemember(rememberHash string) (*User, error) {
	var user User
	err := first(ug.db.Where("remember_hash=?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

//Create will create the provided user and backfill data
//like the id,created at,updated at,deleted at
func (ug *userGorm) Create(user *User) error {
	return ug.db.Create(user).Error

}

//Update will update the provided user with all of the data
//in the provided user object
func (ug *userGorm) Update(user *User) error {

	return ug.db.Save(user).Error

}

//Delete will delete the user with the provided id
func (ug *userGorm) Delete(id uint) error {
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error

}

//Close closes the user service database connection
func (ug *userGorm) Close() error {
	return ug.db.Close()

}

//DestructiveReset drops a table and rebuilds it
func (ug *userGorm) DestructiveReset() error {
	if err := ug.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return ug.AutoMigrate()
}

//AutoMigrate will attempt to automatically migrate the//users table
func (ug *userGorm) AutoMigrate() error {
	if err := ug.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil

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
