package controller

import (
	"encoding/json"
	"log"
	"main/models"
	"main/rabbit_mq"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

func NewMessage(c *fiber.Ctx) error {
	return c.SendString("New message")
}

/*
CreateTransferFromMessage - Función que recibe un mensaje de RabbitMQ y lo procesa
@Param ctx Contexto de la petición
@Return error
*/
func CreateTransferFromMessage(channel *amqp.Channel, msgInBytes []byte) {

	// Crear una nueva transferencia
	var transfer models.TransactionDTO
	err := json.Unmarshal(msgInBytes, &transfer)
	if err != nil {
		log.Printf("Error al deserializar el mensaje: %v", err)
	}

	// Agregar la transferencia a la base de datos
	AddTransaction(channel, transfer)

}

/*
	MessageToQueueConsumer - Función que envía un mensaje a la cola de RabbitMQ
	@Param channel Canal de comunicación con RabbitMQ
	@Param dataTransfer Estructura de la transferencia
*/
func MessageToQueueConsumer (channel *amqp.Channel, dataTransfer models.TransactionDTO) {
	
	dataTransfer.TypeTransaction = "retiro"

	msg,_ := json.Marshal(dataTransfer)

	err := rabbit_mq.SendMessage(channel, "","withdraw",string(msg))
	if err != nil {
		log.Printf("Error al enviar el mensaje: %v", err)
	}

	log.Printf("[TRANSFET]---> Mensaje enviado: %v", dataTransfer)

}