package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/is-ashish/mongo-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://is-ashish:jnp18986@cluster0.01n4k.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// Most Imp
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client Option -> Apply URI
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection success")
	// collection instance
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

// MongoDB Helper ----- File
// Insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted One Movie in DB with ID=", inserted.InsertedID)

}

// update 1 record
func updateOneMoview(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // converting _ID value
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Updated count", result.ModifiedCount)
}

// Deletion of record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("One Movie got deleted with Moview count", deleteCount.DeletedCount)
}

// delete all record
func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of moview deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all Movies

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}
