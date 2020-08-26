package main

import (
	"fmt"

	"../rand"
)

func main() {
	fmt.Println(rand.String(10))
	fmt.Println(rand.RememberToken())
}

// import (
// 	"fmt"

// 	"../models"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// )

// // //User is
// // type User struct {
// // 	gorm.Model
// // 	Namee  string
// // 	Email  string `gorm:"not null;unique_index"`
// // 	Color  string
// // 	Orders []Order
// // }

// // //Order is
// // type Order struct {
// // 	gorm.Model
// // 	UserID      uint
// // 	Amount      int
// // 	Description string
// // }

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "milanpoudel"
// 	dbname   = "lenslocked_dev"
// )

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	us, err := models.NewUserService(psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer us.Close()
// 	us.DestructiveReset()
// 	user := models.User{
// 		Name:  "Milan Poudel",
// 		Email: "milan@gmail.com",
// 	}
// 	if err := us.Create(&user); err != nil {
// 		panic(err)
// 	}

// 	if err := us.Delete(user.ID); err != nil {
// 		panic(err)
// 	}

// 	// user.Email = "milanacademia@yahoo.com"
// 	// if err := us.Update(&user); err != nil {
// 	// 	panic(err)
// 	// }

// 	userByID, err := us.ByID(user.ID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(userByID)
// }

// // db, err := gorm.Open("postgres", psqlInfo)
// // if err != nil {
// // 	panic(err)
// // }
// // defer db.Close()
// // db.LogMode(true)
// // db.AutoMigrate(&User{}, &Order{})

// // var u User = User{
// // 	Color: "Purple",
// // 	Email: "kshitiz@yahoo.com",
// // }

// // var u User
// // if err := db.First(&u).Error; err != nil {
// // 	// 	panic(err)
// // 	// }
// // 	// fmt.Printf("the user is :%+v\n", u)
// // 	// fmt.Println(u)
// // 	// fmt.Println(u.Orders)
// // 	// createOrder(db, u, 1001, "Fake description1")
// // 	// createOrder(db, u, 9999, "Fake description2")
// // 	// createOrder(db, u, 100, "Fake description3")

// // 	//Error handling with GORM
// // 	db = db.Where("email=?", "ngngn@yahoo.com").First(&u)
// // 	if err := db.Where("email=?", "ngngn@yahoo.com").First(&u).Error; err != nil {
// // 		switch err {
// // 		case gorm.ErrRecordNotFound:
// // 			fmt.Println("user not found")
// // // 		default:
// // // 			panic(err)
// // // 		}
// // // 	}
// // // }

// // func createOrder(db *gorm.DB, user User, amount int, desc string) {
// // 	err := db.Create(&Order{
// // 		UserID:      user.ID,
// // 		Amount:      amount,
// // 		Description: desc,
// // 	}).Error
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // }

// // if db.RecordNotFound() {
// // 	fmt.Println("user not found!")

// // } else if db.Error != nil {
// // 	fmt.Println(db.Error)
// // } else {
// // 	fmt.Println(u)
// // }

// // errors := db.GetErrors()
// // if len(errors) > 0 {
// // 	fmt.Println(errors)
// // 	os.Exit(1)
// // }

// // newDB := db.Where("email=?", "blahbh@yahoo.com").First(&u)
// // // newDB = newDB.Or("color=?", "Purple")
// // // newDB = newDB.First(&u)
// // if newDB.Error != nil {
// // 	panic(newDB.Error)
// // }

// //Finding multiple records
// // db.Find(&users)
// // fmt.Println(len(users))
// // fmt.Println(users)

// // //db.First(&u)
// // //db.Last(&u, 12)
// // //newDB := db.Where("id=? AND color=?", 16, "Purple")
// // db.Where(u).First(&u)
// // //newDB.First(&u)
// // fmt.Println("user is:", u)

// // name, email, color := getInfo()
// // u := User{
// // 	Namee: name,
// // 	Email: email,
// // 	Color: color,
// // }
// // if err = db.Create(&u).Error; err != nil {
// // 	panic(err)
// // }
// // fmt.Printf("%+v\n", u)

// // func getInfo() (name, email, color string) {
// // 	reader := bufio.NewReader(os.Stdin)
// // 	fmt.Println("What is your name?")
// // 	name, _ = reader.ReadString('\n')
// // 	fmt.Println("What is your email address?")
// // 	email, _ = reader.ReadString('\n')
// // 	fmt.Println("What is your favorite color?")
// // 	color, _ = reader.ReadString('\n')
// // 	name = strings.TrimSpace(name)
// // 	email = strings.TrimSpace(email)
// // 	color = strings.TrimSpace(color)
// // 	return name, email, color
// // }

// // 	var name string
// // 	var email string

// // 	rows, err := db.Query(`
// //   SELECT users.name,users.email
// //   FROM users
// //   INNER JOIN orders ON users.id=orders.user_id
// //   `)

// // 	for rows.Next() {
// // 		rows.Scan(&name, &email)

// // 		print("Name:", name, "Email:", email, "\n")
// // 	}

// // 	if rows.Err() != nil {
// // 		panic(rows.Err())
// // 	}

// // }

// // for i := 1; i <= 6; i++ {

// // 	userID := 1
// // 	if i < 3 {
// // 		userID = 3
// // 	}
// // 	amount := i * 100
// // 	description := fmt.Sprintf("USB-C Adapter x%d", i)

// // 	rows, err := db.Query(

// // 		`SELECT users.id, users.email, users.name,
// // 		orders.id AS order_id, orders.amount,orders.description
// // 		FROM users
// // 		INNER JOIN orders ON users.id=orders.user_id`)
// // 	// `SELECT * FROM users
// // 	// INNER JOIN orders ON users.id=orders.user_id`)

// // 	// `SELECT * FROM users
// // 	// INNER JOIN orders ON users.id=orders.user_id`)

// // 	fmt.Println(amount, description, rows, err, userID)

// // _, err = db.Exec(`
// // INSERT INTO orders(user_id,amount,description)
// // VALUES($1,$2,$3)`, userID, amount, description)
// // fmt.Println("Code in the middle")
// // if err != nil {
// // 	panic(err)
// // }

// // type User struct {
// // 	Name  string
// // 	Email string
// // }

// // var users []User
// // var name string
// // var email string

// // err = db.QueryRow(`
// //  INSERT INTO users(name,email)
// //  VALUES($1, $2)
// //  RETURNING id`,
// // 	"milanpoudel",
// // 	"milanpoerp@mail.com",
// // ).Scan(&id)
// // if err != nil {
// // 	panic(err)
// // }
// // fmt.Println("id is.....", id)

// // rows := db.QueryRow(`SELECT name,email FROM users WHERE name=$1`, "milanpoungngl")
// // if err != nil {
// // 	panic(err)
// // }
// // err = rows.Scan(&name, &email)
// // if err != nil {
// // 	if err == sql.ErrNoRows {
// // 		fmt.Println("no rows")
// // 	} else {
// // 		panic(err)
// // 	}

// // }

// // rows, err := db.Query(`SELECT name, email FROM users`)
// // if err != nil {
// // 	panic(err)
// // }
// // defer rows.Close()
// // for rows.Next() {
// // 	var user User
// // 	err := rows.Scan(&user.Name, &user.Email)
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	users = append(users, user)
// // }
// // if rows.Err() != nil {
// // 	panic(err)

// // }
// // for _, value := range users {
// // 	fmt.Printf("the name is %s and the email is %s\n", value.Name, value.Email)
// // }
// // fmt.Println(users)
