package rabbit_mq

import (
	"log"

	"github.com/streadway/amqp"
)

/*
Connect - Función para conectar a RabbitMQ
@Param uri: URI de conexión
@return *amqp.Connection: Conexión
@return error: Error
*/
func DeclareExchageChanell(channel *amqp.Channel, nameExchage string, exchange string) error {
	err := channel.ExchangeDeclare(
		nameExchage,
		exchange,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("No se pudo declarar el intercambio: %v", err)
	}
	return err
}

/*
DeclareQueue - Función para declarar una cola
@Param channel: Canal de comunicación
@Param nameQueue: Nombre de la cola
@return error: Error
*/
func DeclareQueue(channel *amqp.Channel, nameQueue string) error {
	_, err := channel.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("No se pudo declarar la cola: %v", err)
	}
	return err
}

/*
SendMessage - Función para enviar mensajes
@param channel: Canal de comunicación
@param nameExchage: Nombre del intercambio
@param routingKey: Clave de enrutamiento
@param message: Mensaje
@return error: Error
*/
func SendMessage(channel *amqp.Channel, nameExchage, routingKey, message string) error {
	err := channel.Publish(
		nameExchage,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Error al enviar mensaje a RabbitMQ: %v", err)
		return err
	}
	return err
}

/*
ListenToMassage - Función para escuchar mensajes
@param channel: Canal de comunicación
@param nameQueue: Nombre de la cola
@param nameConsumer: Nombre del consumidor
@return <-chan amqp.Delivery: Canal de entrega
@return error: Error
*/
func ListenToMassage(channel *amqp.Channel, nameQueue, nameConsumer string) (<-chan amqp.Delivery, error) {
	//channel Rappi, te llegan los mensajes mas rapido que la comida
	chanelRappi, err := channel.Consume(
		nameQueue,
		nameConsumer,
		true,
		false,
		false,
		false,
		nil,
	)
	return chanelRappi, err
}

/*
FailOnError - Función para manejar errores
@Param err: Error
@Param msg: Mensaje
*/
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
