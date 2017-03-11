package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.POST("/users", createUser)

	router.Run(":8080")

}

func getUsers(c *gin.Context) {

	var users []User

	db, err := gorm.Open("sqlite3", "./mailserver.db")

	if err != nil {
		log.Fatal(err)
	}

	db.Find(&users)

	responseObject := gin.H{
		"data": users,
	}
	c.IndentedJSON(http.StatusOK, responseObject)
}

func createUser(c *gin.Context) {

	db, err := gorm.Open("sqlite3", "./mailserver.db")
	if err != nil {
		log.Fatal(err)
	}

	var json UserForm

	if c.BindJSON(&json) == nil {
		user := User{
			Email:    json.Email,
			Password: json.Password,
			DomainID: json.DomainID,
		}

		db.Create(&user)
	}

}
