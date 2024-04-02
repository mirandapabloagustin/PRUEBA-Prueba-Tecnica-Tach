package database

import (
	"os"
	"log"
	"context"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	Transaction *mongo.Collection
)

/*
Init - Función que inicializa la conexión con la base de datos
@Param uri URI de la base de datos
@Param database Nombre de la base de datos
@Return error
*/
func Init(uri string, database string) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseCollectionName := os.Getenv("DATABASE_COLLECTION")

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	
	options := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)
	
	localClient, err := mongo.Connect(context.Background(), options)
	if err != nil {
		return err
	}
	log.Println("Connecting to MongoDB")

	client = localClient

	Transaction = client.Database(database).Collection(databaseCollectionName)

	err = client.Database(database).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	return err
}

/*
Close - Función que cierra la conexión con la base de datos
@Return error
*/
func Close() error {
	return client.Disconnect(context.Background())
}
