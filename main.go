package main

import (
	"fmt"
	"go_test/global"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func main() {
	var tip = "hello this is go build test"
	fmt.Println(tip)
	fmt.Println("GWAF_RELEASE:" + global.GWAF_RELEASE)
	fmt.Println("GWAF_RELEASE_VERSION_NAME:" + global.GWAF_RELEASE_VERSION_NAME)
	fmt.Println("GWAF_RELEASE_VERSION:" + global.GWAF_RELEASE_VERSION)

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{})
	db.Delete(&User{})

	user := User{Name: "Jinzhu", Age: 18}
	db.Create(&user)
	newUser := User{
		Name: "adsf ",
		Age:  0,
	}
	db.First(&newUser)
	fmt.Println(newUser)

}
