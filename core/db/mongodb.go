package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func Insert(obj any) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Collection handle
	collection := client.Database("test").Collection("student")

	// Insert a single document
	//student := Student{"Tom", 18, "male", "Beijing"}
	result, err := collection.InsertOne(ctx, obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted student with ID:", result.InsertedID)

	// Insert multiple documents
	// students := []interface{}{
	// 	Student{"Lucy", 17, "female", "Shanghai"},
	// 	Student{"Jerry", 19, "male", "Guangzhou"},
	// }
	// result, err = collection.InsertMany(ctx, students)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Inserted documents with IDs:", result.InsertedIDs)
}

func Search() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Collection handle
	collection := client.Database("test").Collection("student")

	// Find a single document
	var result Student
	err = collection.FindOne(ctx, bson.M{"name": "Tom"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found student with name:", result.Name)

	// Find multiple documents
	var results []Student
	cur, err := collection.Find(ctx, bson.M{"age": bson.M{"$gt": 17}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found multiple students with age > 17: %+v\n", results)
}
