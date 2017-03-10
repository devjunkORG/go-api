package main

import "github.com/jinzhu/gorm"

// User is a login configuration for the mail server
type User struct {
	gorm.Model

	ID int `gorm:"primary_key"`

	Domain   Domain
	DomainID int `gorm:"index;not null"`

	Email    string
	Password string
}

// TableName set proper tablename
func (User) TableName() string {
	return "virtual_users"
}
