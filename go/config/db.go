package config

import (
	"log"
	"fmt"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/jdotc2/whisper/go/controllers"
)

// Connect Database
func Connect() {
	// Database Config
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//Cancel context to avoid memory leak
	defer cancel()
	
	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)	
	} else {
		log.Println("Connected!")
	}
	db := client.Database("whisper")
	controllers.UserCollection(db)

	return
	}

// Test Connection to database 
func Test() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/test?w=majority"))
		if err != nil {
			panic(err)
		}
	
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	
		// Ping the primary
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			panic(err)
		}
	
		fmt.Println("Successfully connected and pinged.")
}