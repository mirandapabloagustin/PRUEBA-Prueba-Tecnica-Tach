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

func main(){
	//Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	//Obtener variables de entorno
	var(
		dataBaseUri = os.Getenv("DATABASE_URI")
		dataBaseName = os.Getenv("DATABASE_NAME")
		PORT = os.Getenv("PORT")
		rabbitMQ = os.Getenv("RABBITMQ_URI")
	)

	//Inicializar la base de datos
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

	defer func() {
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()


	// Inicialización de RabbitMQ
	conexion, err := amqp.Dial(rabbitMQ)
	rabbit_mq.FailOnError(err, "No se pudo conectar a RabbitMQ")
	defer conexion.Close()

	// Crear un canal de comunicación
	channel, err := conexion.Channel()
	rabbit_mq.FailOnError(err, "No se pudo abrir un canal de comunicación")
	defer channel.Close()

	// CREAR CONSUMER QUEUE - declarar la cola "deposit"
	err = rabbit_mq.DeclareQueue(channel, "deposit")
	rabbit_mq.FailOnError(err, "No se pudo declarar la QUEUE")

	// CREAR PRODUCER QUEUE - declarar la cola "withdraw"
	err = rabbit_mq.DeclareQueue(channel, "withdraw")
	rabbit_mq.FailOnError(err, "No se pudo declarar la QUEUE")

	// CONSUMER QUEUE - consumir mensajes de la cola "deposit"
	chanelRapi, err := rabbit_mq.ListenToMassage(channel, "deposit", "")
	rabbit_mq.FailOnError(err, "No se pudo escuchar la Queue")


	// Crear una instancia de Fiber
	app := fiber.New()

	// Procesar mensajes
	go func() {
		for msg := range chanelRapi {
			controller.CreateTransferFromMessage(channel,msg.Body)
		}
	}()




	// Middleware
	app.Use(recover.New()) 
	app.Use(logger.New())  

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.UserRoutes(app)

	log.Fatal(app.Listen(`:` + PORT))
}