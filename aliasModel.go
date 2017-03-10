package main

import "github.com/jinzhu/gorm"

// Alias represents a mail alias
type Alias struct {
	gorm.Model

	ID          int    `gorm:"primary_key"`
	source      string `gorm:"not null"`
	destination string `gorm:"not null"`

	Domain   Domain
	DomainID int `gorm:"index;not null"`
}
