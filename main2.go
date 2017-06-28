package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id      int
	Name    string
	Address Address
}

type Address struct {
	Id       int
	Address1 string
	UserId   int `sql:"type:integer REFERENCES users(id)"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "testdb.db")
	db.Exec("PRAGMA foreign_keys = ON")
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Address{})
	fmt.Println(db.Save(&User{Name: "jinzhu", Address: Address{Address1: "address 1"}}).Error)
	fmt.Println(db.Save(&Address{Address1: "address 1", UserId: 9999}).Error)
}
