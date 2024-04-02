package routes

import (
	"main/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	routes := app.Group("/apiTransaction/transaction")

	routes.Get("/", controller.GetTransactions)

	routes.Get("/:id", controller.GetTransactionById)

	routes.Delete("/:id", controller.DeleteAccountById)

}