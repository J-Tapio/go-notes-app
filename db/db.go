package db

import (
	"context"
	"log"
	"os"
	"time"

	//"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type NoteFile struct {
	Title string   `bson:"title, omitempty"`
	File  []byte   `bson:"file, omitempty"`
	Tags  []string `bson:"tags, omitempty"`
}

type Note struct {
	Title string   `bson:"title"`
	Tags  []string `bson:"tags"`
}

type Notes []Note

var dbClient *mongo.Client
var dbCollection *mongo.Collection
var AvailableNotes Notes

func InitDBConnection() {
	log.Println("Connecting to Mongo Database")

	/* //Load from local .env file. Only for development purposes.
	err := godotenv.Load()
	if err != nil {
		log.Println("Error while loading .env file")
	} */

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("ATLAS_URI")))

	if err != nil {
		log.Printf("Error while connecting to Mongo database, error:\n%s", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		dbClient = nil
		dbCollection = nil
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return
	}

	dbClient = client
	dbCollection = client.Database("markdown-notes").Collection("notes")
}

func Disconnect() {
	log.Println("Disconnecting from Mongo Database")

	if dbClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		if err := dbClient.Disconnect(ctx); err != nil {
			log.Printf("Error with disconnecting from Mongo Database: %s\n", err)
		}
		cancel()
	}
}

func GetAvailableNotes() {
	log.Println("Fetching available documents")

	// If initial db connection did not fail
	if dbClient != nil {

		opts := options.Find().SetProjection(bson.D{{Key: "title", Value: 1}, {Key: "tags", Value: 1}})

		cur, err := dbCollection.Find(context.TODO(), bson.D{}, opts)
		cur.All(context.TODO(), &AvailableNotes)

		if err != nil {
			log.Println(err)
			AvailableNotes = []Note{}
		}
	}
}

func GetDocument(title string) (bool, NoteFile) {
	log.Printf("Fetching document: %s\n", title)

	var document NoteFile

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err := dbCollection.FindOne(ctx, bson.D{primitive.E{Key: "title", Value: title}}).Decode(&document)

	if err != nil {
		log.Printf("Something went wrong with finding document: '%s', Error:  %s\n", title, err)
		cancel()
		return false, document
	} else {
		cancel()
		return true, document
	}
}
