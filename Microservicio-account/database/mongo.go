package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	client  *mongo.Client
	Account *mongo.Collection
)

/*
Init - Inicializa la conexión con la base de datos
@Params: uri string - URI de la base de datos
@Params: database string - Nombre de la base de datos
@Return: error - Error en caso de que la conexión falle
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

	Account = client.Database(database).Collection(databaseCollectionName)

	err = client.Database(database).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	return err
}

/*
Close - Cierra la conexión con la base de datos
@Return: error - Error en caso de que la conexión falle
*/
func Close() error {
	return client.Disconnect(context.Background())
}

/*
FindAccountById - Busca una cuenta por su ID
@Params: id string - ID de la cuenta
@Return: bson.M - Cuenta encontrada
@Return: error - Error en caso de que la conexión falle
*/
func FindAccountById(id string) (bson.M, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	account := bson.M{}
	err := Account.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&account)
	return account, err
}

/*
UpdateAccountById - Actualiza una cuenta por su ID
@Params: id string - ID de la cuenta
@Params: account bson.M - Cuenta a actualizar
*/
func UpdateAccountById(id string, account bson.M) {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := Account.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": account})
	if err != nil {
		log.Fatalf("Error updating account: %v", err)
	}
}
