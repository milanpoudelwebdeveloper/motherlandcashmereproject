//Embedding Interfaces and Chaining
package main

import "fmt"

type User struct {
	Name string
}
type UserReader interface {
	ByID(id uint) *User
}

type UserService struct {
	UserReader
}

type UserValidator struct {
	userReader UserReader
}

type UserCache struct {
	UserReader
}

type UserGorm struct {
}

func (ug UserGorm) ByID(id uint) *User {
	return &User{
		Name: "Milan Poudel",
	}
}

func main() {
	uv := UserValidator{
		userReader: UserGorm{},
	}
	us := UserService{
		UserReader: uv.userReader,
	}
	fmt.Println(us.ByID(2))

}

//Interfaces and chaining
// package main

// import "fmt"

// type Dog struct {
// }

// func (d Dog) Speak() {
// 	fmt.Println("Bark")
// }

// type Cat struct {
// }

// func (c Cat) Speak() {
// 	fmt.Println("Meow")
// }

// type Speaker struct {
// 	s SpeakerInterface
// }

// type SpeakerPrefixer struct {
// 	sp SpeakerInterface
// }

// func (sp SpeakerPrefixer) Speak() {
// 	fmt.Print("Prefix:")
// 	sp.sp.Speak()

// }

// type SpeakerInterface interface {
// 	Speak()
// }

// func main() {
// 	s := Speaker{s: SpeakerPrefixer{sp: Cat{}}}
// 	s.s.Speak()
// }

//Chaining and interfaces

// package main

// import "fmt"

// //Interfaces and Embedding
// type Speaker interface {
// 	Speak() string
// }

// type Animal struct {
// 	Speaker
// }

// type Dog struct {
// 	name string
// }

// func (d Dog) Speak() string {

// 	return d.name

// }

// type Cat struct {
// 	name string
// }

// func (c Cat) Speak() string {
// 	return c.name

// }

// func main() {
// 	c := Cat{name: "Straw Cat"}
// 	//d := Dog{name: "Bull Dog"

// 	animal := Animal{c}
// 	fmt.Println(animal.Speak())

// }

// package main

// import "fmt"

// type Cat struct {
// }

// func (c Cat) Speak() {
// 	fmt.Println("meow")
// }

// type Dog struct{}

// func (d Dog) Speak() {
// 	fmt.Println("woof")
// }

// type Husky struct {
// 	Speaker
// }

// type Speaker interface {
// 	Speak()
// }

// func main() {
// 	husky := Husky{Cat{}}
// 	husky.Speak()
// }

// package main

// import (
// 	"fmt"

// 	"../models"
// )

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
// 	//us.AutoMigrate()
// 	user := models.User{
// 		Name:     "Milan Poudel",
// 		Email:    "milan@yahoo.com",
// 		Password: "milan",
// 		Remember: "abc123",
// 	}
// 	err = us.Create(&user)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", user)
// 	user2, err := us.ByRemember("abc123")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", *user2)

// 	// toHash := []byte("this is my string to hash")
// 	// h := hmac.New(sha256.New, []byte("my-secret-key"))
// 	// h.Write(toHash)
// 	// b := h.Sum(nil)
// 	// fmt.Println(base64.URLEncoding.EncodeToString(b))
// 	// hmac := hash.NewHMAC("my-secret-key")
// 	// fmt.Println(hmac.HASH("this is my string to hash"))

// }

// // import (
// // 	"fmt"

// // 	"../models"
// // 	_ "github.com/jinzhu/gorm/dialects/postgres"
// // )

// // // //User is
// // // type User struct {
// // // 	gorm.Model
// // // 	Namee  string
// // // 	Email  string `gorm:"not null;unique_index"`
// // // 	Color  string
// // // 	Orders []Order
// // // }

// // // //Order is
// // // type Order struct {
// // // 	gorm.Model
// // // 	UserID      uint
// // // 	Amount      int
// // // 	Description string
// // // }

// // const (
// // 	host     = "localhost"
// // 	port     = 5432
// // 	user     = "postgres"
// // 	password = "milanpoudel"
// // 	dbname   = "lenslocked_dev"
// // )

// // func main() {
// // 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// // 	us, err := models.NewUserService(psqlInfo)
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	defer us.Close()
// // 	us.DestructiveReset()
// // 	user := models.User{
// // 		Name:  "Milan Poudel",
// // 		Email: "milan@gmail.com",
// // 	}
// // 	if err := us.Create(&user); err != nil {
// // 		panic(err)
// // 	}

// // 	if err := us.Delete(user.ID); err != nil {
// // 		panic(err)
// // 	}

// // 	// user.Email = "milanacademia@yahoo.com"
// // 	// if err := us.Update(&user); err != nil {
// // 	// 	panic(err)
// // 	// }

// // 	userByID, err := us.ByID(user.ID)
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	fmt.Println(userByID)
// // }

// // // db, err := gorm.Open("postgres", psqlInfo)
// // // if err != nil {
// // // 	panic(err)
// // // }
// // // defer db.Close()
// // // db.LogMode(true)
// // // db.AutoMigrate(&User{}, &Order{})

// // // var u User = User{
// // // 	Color: "Purple",
// // // 	Email: "kshitiz@yahoo.com",
// // // }

// // // var u User
// // // if err := db.First(&u).Error; err != nil {
// // // 	// 	panic(err)
// // // 	// }
// // // 	// fmt.Printf("the user is :%+v\n", u)
// // // 	// fmt.Println(u)
// // // 	// fmt.Println(u.Orders)
// // // 	// createOrder(db, u, 1001, "Fake description1")
// // // 	// createOrder(db, u, 9999, "Fake description2")
// // // 	// createOrder(db, u, 100, "Fake description3")

// // // 	//Error handling with GORM
// // // 	db = db.Where("email=?", "ngngn@yahoo.com").First(&u)
// // // 	if err := db.Where("email=?", "ngngn@yahoo.com").First(&u).Error; err != nil {
// // // 		switch err {
// // // 		case gorm.ErrRecordNotFound:
// // // 			fmt.Println("user not found")
// // // // 		default:
// // // // 			panic(err)
// // // // 		}
// // // // 	}
// // // // }

// // // func createOrder(db *gorm.DB, user User, amount int, desc string) {
// // // 	err := db.Create(&Order{
// // // 		UserID:      user.ID,
// // // 		Amount:      amount,
// // // 		Description: desc,
// // // 	}).Error
// // // 	if err != nil {
// // // 		panic(err)
// // // 	}
// // // }

// // // if db.RecordNotFound() {
// // // 	fmt.Println("user not found!")

// // // } else if db.Error != nil {
// // // 	fmt.Println(db.Error)
// // // } else {
// // // 	fmt.Println(u)
// // // }

// // // errors := db.GetErrors()
// // // if len(errors) > 0 {
// // // 	fmt.Println(errors)
// // // 	os.Exit(1)
// // // }

// // // newDB := db.Where("email=?", "blahbh@yahoo.com").First(&u)
// // // // newDB = newDB.Or("color=?", "Purple")
// // // // newDB = newDB.First(&u)
// // // if newDB.Error != nil {
// // // 	panic(newDB.Error)
// // // }

// // //Finding multiple records
// // // db.Find(&users)
// // // fmt.Println(len(users))
// // // fmt.Println(users)

// // // //db.First(&u)
// // // //db.Last(&u, 12)
// // // //newDB := db.Where("id=? AND color=?", 16, "Purple")
// // // db.Where(u).First(&u)
// // // //newDB.First(&u)
// // // fmt.Println("user is:", u)

// // // name, email, color := getInfo()
// // // u := User{
// // // 	Namee: name,
// // // 	Email: email,
// // // 	Color: color,
// // // }
// // // if err = db.Create(&u).Error; err != nil {
// // // 	panic(err)
// // // }
// // // fmt.Printf("%+v\n", u)

// // // func getInfo() (name, email, color string) {
// // // 	reader := bufio.NewReader(os.Stdin)
// // // 	fmt.Println("What is your name?")
// // // 	name, _ = reader.ReadString('\n')
// // // 	fmt.Println("What is your email address?")
// // // 	email, _ = reader.ReadString('\n')
// // // 	fmt.Println("What is your favorite color?")
// // // 	color, _ = reader.ReadString('\n')
// // // 	name = strings.TrimSpace(name)
// // // 	email = strings.TrimSpace(email)
// // // 	color = strings.TrimSpace(color)
// // // 	return name, email, color
// // // }

// // // 	var name string
// // // 	var email string

// // // 	rows, err := db.Query(`
// // //   SELECT users.name,users.email
// // //   FROM users
// // //   INNER JOIN orders ON users.id=orders.user_id
// // //   `)

// // // 	for rows.Next() {
// // // 		rows.Scan(&name, &email)

// // // 		print("Name:", name, "Email:", email, "\n")
// // // 	}

// // // 	if rows.Err() != nil {
// // // 		panic(rows.Err())
// // // 	}

// // // }

// // // for i := 1; i <= 6; i++ {

// // // 	userID := 1
// // // 	if i < 3 {
// // // 		userID = 3
// // // 	}
// // // 	amount := i * 100
// // // 	description := fmt.Sprintf("USB-C Adapter x%d", i)

// // // 	rows, err := db.Query(

// // // 		`SELECT users.id, users.email, users.name,
// // // 		orders.id AS order_id, orders.amount,orders.description
// // // 		FROM users
// // // 		INNER JOIN orders ON users.id=orders.user_id`)
// // // 	// `SELECT * FROM users
// // // 	// INNER JOIN orders ON users.id=orders.user_id`)

// // // 	// `SELECT * FROM users
// // // 	// INNER JOIN orders ON users.id=orders.user_id`)

// // // 	fmt.Println(amount, description, rows, err, userID)

// // // _, err = db.Exec(`
// // // INSERT INTO orders(user_id,amount,description)
// // // VALUES($1,$2,$3)`, userID, amount, description)
// // // fmt.Println("Code in the middle")
// // // if err != nil {
// // // 	panic(err)
// // // }

// // // type User struct {
// // // 	Name  string
// // // 	Email string
// // // }

// // // var users []User
// // // var name string
// // // var email string

// // // err = db.QueryRow(`
// // //  INSERT INTO users(name,email)
// // //  VALUES($1, $2)
// // //  RETURNING id`,
// // // 	"milanpoudel",
// // // 	"milanpoerp@mail.com",
// // // ).Scan(&id)
// // // if err != nil {
// // // 	panic(err)
// // // }
// // // fmt.Println("id is.....", id)

// // // rows := db.QueryRow(`SELECT name,email FROM users WHERE name=$1`, "milanpoungngl")
// // // if err != nil {
// // // 	panic(err)
// // // }
// // // err = rows.Scan(&name, &email)
// // // if err != nil {
// // // 	if err == sql.ErrNoRows {
// // // 		fmt.Println("no rows")
// // // 	} else {
// // // 		panic(err)
// // // 	}

// // // }

// // // rows, err := db.Query(`SELECT name, email FROM users`)
// // // if err != nil {
// // // 	panic(err)
// // // }
// // // defer rows.Close()
// // // for rows.Next() {
// // // 	var user User
// // // 	err := rows.Scan(&user.Name, &user.Email)
// // // 	if err != nil {
// // // 		panic(err)
// // // 	}
// // // 	users = append(users, user)
// // // }
// // // if rows.Err() != nil {
// // // 	panic(err)

// // // }
// // // for _, value := range users {
// // // 	fmt.Printf("the name is %s and the email is %s\n", value.Name, value.Email)
// // // }
// // // fmt.Println(users)
