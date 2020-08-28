package controllers

import (
	"net/http"
	"log"
	"time"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "github.com/jdotc2/blue-apricot/server/config"
)

// User Object
type User struct {
	ID        int    `json:"id"`
	Name			string		`json:"name"`
	Images	  struct {
		Main			string		`json:"main"`
		Map				string		`json:"map"`
	}
	Exclusive struct {
		Pokemon		[]int			`json:"pokemon"`
	}
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


// DATABASE INSTANCE
var collection *mongo.Collection

// UserCollection Client Connection
func UserCollection(c *mongo.Database) {
	collection = c.Collection("user")
}

func (user *User) setID(ID int) {
	user.ID = ID
}

// GetAllUsers Get all
func GetAllUsers(c *gin.Context) {
	users := []User{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	id := 0

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
			var user User
			cursor.Decode(&user)
			user.setID(id)
			id++
			users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    users,
	})
	return
}
