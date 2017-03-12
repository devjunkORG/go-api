package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	db, err := gorm.Open("sqlite3", "./mailserver.db")

	if err != nil {
		log.Fatal(err)
	}

	router.GET("/users", getUsers(db))
	router.POST("/users", createUser(db))

	router.Run(":8080")

}

func getUsers(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var users []User

		db.Find(&users)

		responseObject := gin.H{
			"data": users,
		}
		c.IndentedJSON(http.StatusOK, responseObject)
	}

	return gin.HandlerFunc(fn)
}

func createUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var formData UserForm

		requestError := c.Bind(&formData)

		if requestError == nil {
			user := User{
				Email:    formData.Email,
				Password: formData.Password,
				DomainID: formData.DomainID,
			}

			db.Create(&user)

			jsonResponse, err := json.Marshal(user)

			if err != nil {
				log.Fatal(err)
			}

			responseObject := gin.H{
				"data": jsonResponse,
			}
			c.IndentedJSON(http.StatusOK, responseObject)
		} else {
			responseObject := gin.H{
				"error": requestError.Error(),
			}
			c.IndentedJSON(http.StatusBadRequest, responseObject)
		}
	}

	return gin.HandlerFunc(fn)
}
