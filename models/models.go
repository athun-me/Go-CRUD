package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Place string
	Age   int64
}
