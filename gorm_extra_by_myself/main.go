package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Human is
type Human struct {
	gorm.Model
	HumanName string
	Email     string `gorm:"not null;unique_index"`
	Color     string
	Hobbies   string
	Country   string
	Porders   []Porder
	City      string
}

//Porder is
type Porder struct {
	gorm.Model
	HumanID     uint
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
	db.LogMode(true)
	db.AutoMigrate(&Human{}, &Porder{})

	var h []Human
	if err := db.Preload("Porders").Find(&h).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("User not found")
		default:
			panic(err)
		}
	}
	fmt.Println("Hey, the human record is", h)
}

func createMenu(db *gorm.DB, human Human, amount int, desc string) {
	db.Create(&Porder{
		HumanID:     human.ID,
		Amount:      amount,
		Description: desc,
	})

}
