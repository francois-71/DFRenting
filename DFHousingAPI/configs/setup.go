package configs

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/gridfs"
)

var DB *mongo.Client
var GridFS *gridfs.Bucket

func ConnectDB() (*mongo.Client, error) {
    client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
    if err != nil {
        return nil, err
    }
  
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err = client.Connect(ctx); err != nil {
        return nil, err
    }

    // Ping the database
    if err = client.Ping(ctx, nil); err != nil {
        return nil, err
    }
    log.Println("Connected to MongoDB")

    // Set up GridFS
    dbName := "DFHousingAPI" // Replace with your actual database name
    GridFS, err = setupGridFSBucket(client, dbName)
    if err != nil {
        return nil, err
    }

    DB = client // Set the client to the global variable

    return client, nil
}

// GetCollection returns a MongoDB collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("DFHousingAPI").Collection(collectionName)
    return collection
}

// setupGridFSBucket sets up a GridFS bucket
func setupGridFSBucket(client *mongo.Client, dbName string) (*gridfs.Bucket, error) {
    // Get the database from the client using the provided dbName
    database := client.Database(dbName)

    // Create a GridFS bucket using the obtained database
    fs, err := gridfs.NewBucket(
        database, // Pass the database directly to gridfs.NewBucket
    )
    if err != nil {
        return nil, err
    }
    return fs, nil
}

func init() {
    var err error
    DB, err = ConnectDB()
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %s", err.Error())
    }
}
