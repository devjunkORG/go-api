package main

import "github.com/jinzhu/gorm"

// Domain represents a domain configured to work with the mail server
type Domain struct {
	gorm.Model

	ID   int    `gorm:"primary_key"`
	name string `gorm:"not null"`

	Users   []User
	Aliases []Alias
}
