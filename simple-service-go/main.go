package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// CREATE TABLE "public"."res_users" (
// "id" int4 NOT NULL DEFAULT nextval('res_users_id_seq'::regclass),
// "login" bpchar NOT NULL,
// "password" bpchar,
// PRIMARY KEY ("id")
// );

type User struct {
	Id       int    `json:"id,omitempty" gorm:"column:id"`
	Login    string `json:"login" gorm:"column:login"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "res_users"
}

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db, err)

	// insert new user
	//newUser := User{Login: "demo", Password: "demo"}
	//if err := db.Create(&newUser).Error; err != nil {
	//	fmt.Println(err)
	//}

	// select
	var users []User
	db.Where("").Find(&users)
	fmt.Println(users)

	var user User
	if err := db.Where("id = 1").First(&user).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(user)

	// delete
	//db.Table(User{}.TableName()).Where("id = 2").Delete(nil)
	db.Table(User{}.TableName()).Where("id = 1").Delete(nil)
}
