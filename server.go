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

	router.Run(":8080")

}

func getUsers(context *gin.Context) {

	var users []User

	db, err := gorm.Open("sqlite3", "./mailserver.db")

	if err != nil {
		log.Fatal(err)
	}

	db.Find(&users)

	responseObject := gin.H{
		"data": users,
	}
	context.IndentedJSON(http.StatusOK, responseObject)
}
