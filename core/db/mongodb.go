package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	Dbtype   string
	Host     string
	Port     int
	Database string
}

var Config DBConfig

func Connect() {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", Config.Host, Config.Port))
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

func init() {
	jsonFile, err := os.Open("dbConfig.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
		panic(err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("Cannot get configuration from file")
		panic(err)
	}
	fmt.Println("Config:")
	fmt.Println(Config)
}

func Insert(collectionName string, obj any) {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", Config.Host, Config.Port))

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Collection handle
	collection := client.Database(Config.Database).Collection(collectionName)

	// Insert a single document
	//student := Student{"Tom", 18, "male", "Beijing"}
	result, err := collection.InsertOne(ctx, obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted data with ID:", result.InsertedID)

	// Insert multiple documents
	// students := []interface{}{
	// 	Student{"Lucy", 17, "female", "Shanghai"},
	// 	Student{"Jerry", 19, "male", "Guangzhou"},
	// }
	// result, err = collection.InsertMany(ctx, students)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println("Inserted documents with IDs:", result.InsertedIDs)
}

func Search[T any](collectionName string) []T {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", Config.Host, Config.Port))

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Collection handle
	collection := client.Database(Config.Database).Collection(collectionName)

	// // Find a single document
	// var result T
	// err = collection.FindOne(ctx, bson.M{"name": "Tom"}).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Find multiple documents
	var results []T
	cur, err := collection.Find(ctx, bson.M{"age": bson.M{"$gt": 17}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var elem T
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

/*
查询最后一个区块
*/
func SearchLastOne[T any](collectionName string, orderCol string, org T) *T {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", Config.Host, Config.Port))

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Collection handle
	collection := client.Database(Config.Database).Collection(collectionName)
	l := options.Find().SetLimit(1)
	s := options.Find().SetSort(bson.D{{orderCol, -1}}) //"blockheader.timestamp"

	var result T
	cur, err := collection.Find(ctx, bson.M{}, l, s)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	if cur.Next(ctx) {
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return &result
}
