package main

import "github.com/jinzhu/gorm"

// User is a login configuration for the mail server
type User struct {
	gorm.Model

	Domain   Domain `gorm:"ForeignKey:DomainID"`
	DomainID int    `gorm:"index;not null"`

	Email    string
	Password string
}

func (User) TableName() string {
	return "virtual_users"
}
