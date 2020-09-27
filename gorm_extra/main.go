package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//User is
type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Color  string
	Orders []Order
}

//Order is
type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "milanpoudel"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var u []User

	if err := db.Preload("Orders").Find(&u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("the user is %+v\n", u)
	// createOrder(db, u, 1001, "Fake description 5")
	// createOrder(db, u, 1002, "Fake decription 6")
	// createOrder(db, u, 1003, "Fake description 7")
	// db = db.Where("email=?", "maikskb@uahoo.com").First(&u)
	// if err := db.Where("email=?", "maikskb@uahoo.com").First(&u).Error; err != nil {
	// 	switch err {
	// 	case gorm.ErrRecordNotFound:
	// 		fmt.Println("User not found")
	// 	default:
	// 		panic(err)
	// 	}

	// }
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}

}
