package main

import "gorm.io/gorm"

type User struct {
	Name string
	Age  int
	gorm.Model
}
