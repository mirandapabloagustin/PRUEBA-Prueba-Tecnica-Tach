package routes

import (
	"main/controller"
	"github.com/streadway/amqp"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, channel *amqp.Channel) {
	routes := app.Group("/apiAccount/accounts")

	routes.Get("/",controller.GetAccounts)

	routes.Get("/:id",controller.GetAccountById)
	
	routes.Post("/add",controller.AddAccount)

	routes.Put("/:id",controller.UpdateAccountById)

	routes.Delete("/:id",controller.DeleteAccountById)

	routes.Post("/:id/send", func(contex *fiber.Ctx) error{
		return controller.SendTranferMessage(channel,contex)
	})

}