package rabbit_mq

import (
	"log"

	"github.com/streadway/amqp"
)

/*
	DeclareExchangeChanell - Función para declarar un canal de intercambio
	@Param channel *amqp.Channel - Canal de comunicación
	@Param nameExchage string - Nombre del intercambio
	@Param exchange string - Tipo de intercambio
	@Return error - Error
*/
func DeclareExchageChanell(channel *amqp.Channel, nameExchage string, exchange string) error {
	err := channel.ExchangeDeclare(
		nameExchage, // Nombre del intercambio
		exchange,    // Tipo de intercambio
		false,       // Duradero
		false,       // Autoeliminable
		false,       // Interno
		false,       // No se espera confirmación
		nil,         // Argumentos adicionales
	)
	if err != nil {
		log.Fatalf("No se pudo declarar el intercambio: %v", err)
	}
	return err
}

/*
	DeclareQueue - Función para declarar una cola
	@Param channel *amqp.Channel - Canal de comunicación
	@Param nameQueue string - Nombre de la cola
	@Return error - Error
*/
func DeclareQueue(channel *amqp.Channel, nameQueue string) error {
	_, err := channel.QueueDeclare(
		nameQueue, // Nombre de la cola
		false,      // Duradero
		false,     // Autoeliminable
		false,     // Exclusivo
		false,     // No esperar confirmación
		nil,       // Argumentos adicionales
	)
	if err != nil {
		log.Fatalf("No se pudo declarar la cola: %v", err)
	}
	return err
}

/*
	BindQueue - Función para enlazar una cola
	@Param channel *amqp.Channel - Canal de comunicación
	@Param nameQueue string - Nombre de la cola
	@Param nameExchage string - Nombre del intercambio
	@Param routingKey string - Clave de enrutamiento
	@Return error - Error
*/
func SendMessage(channel *amqp.Channel, nameExchage, routingKey, message string) error {
	err := channel.Publish(
		nameExchage, // Nombre del intercambio
		routingKey,  // Clave de enrutamiento
		false,       // Mandar a confirmar
		false,       // Publicación inmediata
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	return err
}

/* 
	ListenToMassage - Función para escuchar mensajes
	@Param channel *amqp.Channel - Canal de comunicación
	@Param nameQueue string - Nombre de la cola
	@Param nameConsumer string - Nombre del consumidor
	@Return <-chan amqp.Delivery - Canal de mensajes
	@Return error - Error
*/
func ListenToMassage(channel *amqp.Channel, nameQueue,nameConsumer string) ( <-chan amqp.Delivery, error) {
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
	@Param err error - Error
	@Param msg string - Mensaje
*/
func FailOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
  }