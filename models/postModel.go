package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}

type Product struct {
	gorm.Model
	Name  string
	Desc  string
	Price int64
}
