package main

import "github.com/jinzhu/gorm"

// Alias represents a mail alias
type Alias struct {
	gorm.Model

	source      string `gorm:"not null"`
	destination string `gorm:"not null"`

	Domain   Domain
	DomainID int `gorm:"index;not null"`
}

func (Alias) TableName() string {
	return "virtual_aliases"
}
