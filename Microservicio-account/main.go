package main

import (
	"fmt"
	"log"
	"main/controller"
	"main/database"
	"main/rabbit_mq"
	"main/routes"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Variables de entorno
	var (
		dataBaseUri  = os.Getenv("DATABASE_URI")
		dataBaseName = os.Getenv("DATABASE_NAME")
		PORT         = os.Getenv("PORT")
		rabbitMQ     = os.Getenv("RABBITMQ_URI")
	)

	// Inicialización de la base de datos con reintento en caso de error
	var err error
	for {
		err = database.Init(dataBaseUri, dataBaseName)
		if err != nil {
			log.Printf("Error initializing database: %v. Retrying in 2 seconds...\n", err)
			time.Sleep(10 * time.Second) // Esperar antes de volver a intentar
			continue
		}
		break
	}
	
	// Cierre de la base de datos al finalizar
	defer func() {
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Inicialización de RabbitMQ
	conexion, err := amqp.Dial(rabbitMQ)
	rabbit_mq.FailOnError(err, "No se pudo establecer conexión con RabbitMQ")
	defer conexion.Close()

	// Crear un CANAL de comunicación
	channel, err := conexion.Channel()
	rabbit_mq.FailOnError(err, "No se pudo abrir el canal de comunicación")
	defer channel.Close()
	
	// CREAR PRODUCER QUEUE - declarar la cola "deposit"
	err = rabbit_mq.DeclareQueue(channel, "deposit")
	rabbit_mq.FailOnError(err, "No se pudo declarar la Queue")

	// CREAR CONSUMER QUEUE - escuchar mensajes en la cola "deposit"
	channelRappi,err := rabbit_mq.ListenToMassage(channel, "withdraw","")
	rabbit_mq.FailOnError(err, "No se pudo escuchar la Queue")

	//Configuracion de Fiber
	app := fiber.New()

	go func ()  {
		for msg := range channelRappi {
			fmt.Println("[ACCOUNT]--> Mensaje recibido: ", string(msg.Body))
			controller.AccionMessage(msg.Body)
		}
	}()

	// Middleware
	app.Use(recover.New()) // Middleware de recuperación
	app.Use(logger.New())  // Middleware de registro

	// Ruta de estado
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Microservicio de cuentas escuchando en el puerto " + PORT)
	})

	// Configuración de las rutas
	routes.UserRoutes(app, channel)

	// Iniciar la aplicación en el puerto especificado
	log.Fatal(app.Listen(`:` + PORT))

}
