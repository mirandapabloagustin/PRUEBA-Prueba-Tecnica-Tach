package controller

import (
	"context"
	"log"
	"main/database"
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
GetTransactions - Función que obtiene todas las transacciones
@Param ctx Contexto de la petición
@Return error
*/
func GetTransactions(ctx *fiber.Ctx) error {
	cursor, err := database.Transaction.Find(ctx.Context(), bson.M{})
	//verificar si hay un error al obtener los datos
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error getting transactions",
			"error":   err,
		})
	}

	var transactions []bson.M
	//Parsear los datos obtenidos al arreglo de transactions
	if err = cursor.All(ctx.Context(), &transactions); err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error getting transactions",
			"error":   err,
		})
	}

	//Iterar sobre las transacciones para mostrar las transacciones
	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Transactions found",
		"data":    transactions,
	})
}


/*
GetTransactionById - Función que obtiene una transacción por su id
@Param ctx Contexto de la petición
@Return error
*/
func GetTransactionById(ctx *fiber.Ctx) error {
	//Obtener el id de la transacción
	id := ctx.Params("id")
	
	//Convertir el id a un ObjectID
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing id",
			"error":   err,
		})
	}
	
	//Obtener la transacción
	var transaction models.Transaction
	err = database.Transaction.FindOne(ctx.Context(), bson.M{"_id": objId}).Decode(&transaction)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error getting transaction",
			"error":   err,
		})
	}
	
	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Transaction found",
		"data":    transaction,
	})
}


/*
CreateTransaction - Función que crea una transacción
@Param ctx Contexto de la petición
@Return error
*/
func DeleteAccountById (ctx *fiber.Ctx) error {
	//Obtener el id de la transacción
	id := ctx.Params("id")
	
	//Convertir el id a un ObjectID
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing id",
			"error":   err,
		})
	}
	
	//Eliminar la transacción
	_, err = database.Transaction.DeleteOne(ctx.Context(), bson.M{"_id": objId})
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error deleting transaction",
			"error":   err,
		})
	}
	
	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Transaction deleted",
	})
}

/*
AddTransaction - Función que agrega una transacción a la base de datos
@Param channel Canal de comunicación con RabbitMQ
@Param dtoTransfer Estructura de la transferencia
*/
func AddTransaction(channel *amqp.Channel,dtoTransfer models.TransactionDTO) {

	newTransfer := models.Transaction{
		IdTransaction:   uuid.New(),
		SenderAccount:   dtoTransfer.SenderAccount,
		ReceiverAccount: dtoTransfer.ReceiverAccount,
		Amount:          dtoTransfer.Amount,
		Date:            time.Now().Format("2006-01-02"),
		Time:            time.Now().Format("15:04:05"),
	}

	_, err := database.Transaction.InsertOne(context.TODO(), newTransfer)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[TRANSFET]---> Mensaje recibido y procesado: %v", dtoTransfer)

	MessageToQueueConsumer(channel, dtoTransfer)
}