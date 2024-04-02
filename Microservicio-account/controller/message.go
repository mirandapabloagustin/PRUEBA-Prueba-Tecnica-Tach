package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"main/database"
	"main/models"
	"main/rabbit_mq"
	"os"
)

/*
SendTranferMessage - Función que se encarga de enviar un mensaje de transferencia
@Param channel *amqp.Channel - Canal de comunicación con RabbitMQ
@Param ctx *fiber.Ctx - Contexto de la petición
@Return error - Error en caso de que la petición falle
*/
func SendTranferMessage(channel *amqp.Channel, ctx *fiber.Ctx) error {
	r_key := os.Getenv("RABBITMQ_ROUTING_KEY")
	id := ctx.Params("id")

	// Convertir la cadena en un ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid ID",
			"error":   err,
		})
	}

	var transaction models.TransactionDTO
	if err := ctx.BodyParser(&transaction); err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing data request",
			"error":   err,
		})
	}

	var account models.Account
	err = database.Account.FindOne(ctx.Context(), bson.M{"_id": objID}).Decode(&account)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error getting account",
			"error":   err,
		})
	}

	transaction.SenderAccount = id

	//Transformamos el mensaje a JSON
	message, _ := json.Marshal(transaction)

	errSend := rabbit_mq.SendMessage(channel, "", r_key, string(message))
	if err != nil {
		rabbit_mq.FailOnError(errSend, "No se pudo enviar el mensaje")
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error sending message",
			"error":   err,
		})
	} else {
		log.Println("[ACCOUNT]--> Mensaje enviado: ", transaction)
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Message sent",
			"data":    transaction,
		})
	}
}

/*
AccionMessage - Función que se encarga de realizar la acción de un mensaje
@Param msgBytes []byte - Mensaje en bytes
*/
func AccionMessage(msgBytes []byte) {
	var dataTrasfer models.TransactionDTO

	err := json.Unmarshal(msgBytes, &dataTrasfer)
	if err != nil {
		log.Fatalf("Error al decodificar el mensaje: %v", err)
	}

	// Buscar la cuenta del Recepetor
	accoun_R, err := database.FindAccountById(dataTrasfer.ReceiverAccount)
	if err != nil {
		log.Fatalf("Error al buscar la cuenta del receptor: %v", err)
	}

	balance_R, ok := accoun_R["balance"].(float64)
	if !ok {
		log.Fatalf("El balance no es un número")
	}
	balance_R += dataTrasfer.Amount
	accoun_R["balance"] = balance_R

	//Busca la cuenta del Emisor
	account_S, err := database.FindAccountById(dataTrasfer.SenderAccount)

	balance_S, ok := account_S["balance"].(float64)
	if !ok {
		log.Fatalf("El balance no es un número")
	}
	balance_S -= dataTrasfer.Amount
	account_S["balance"] = balance_S

	// Actualiza la cuenta
	database.UpdateAccountById(dataTrasfer.ReceiverAccount, accoun_R)
	database.UpdateAccountById(dataTrasfer.SenderAccount, account_S)
}
