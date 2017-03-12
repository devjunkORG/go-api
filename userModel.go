package main

import (
	"math/rand"

	"github.com/jinzhu/gorm"
	"github.com/kless/osutil/user/crypt/sha512_crypt"
)

// User is a login configuration for the mail server
type User struct {
	gorm.Model

	Domain   Domain `gorm:"ForeignKey:DomainID"`
	DomainID int    `gorm:"index;not null"`

	Email    string
	Password string
}

// TableName sets table name
func (User) TableName() string {
	return "virtual_users"
}

// BeforeCreate callback
func (u *User) BeforeCreate() (err error) {
	u.Password = encryptPassword(u.Password)
	return
}

func generateSalt() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return "$6$" + string(b)
}

func encryptPassword(password string) string {
	c := sha512_crypt.New()
	hash, err := c.Generate([]byte(password), []byte(generateSalt()))
	if err != nil {
		panic(err)
	}
	return "{SHA512-CRYPT}" + hash
}

// UserForm : form binding
type UserForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	DomainID int    `form:"domain_id" json:"domain_id" binding:"required"`
}
