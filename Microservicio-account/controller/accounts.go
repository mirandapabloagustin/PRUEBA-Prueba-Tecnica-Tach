package controller

import (
	"log"
	"main/database"
	"main/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	GetAccounts - Función que se encarga de obtener todas las cuentas
	@Param ctx *fiber.Ctx - Contexto de la petición
	@Return error - Error en caso de que la petición falle
*/
func GetAccounts(ctx *fiber.Ctx) error {
	cursor, err := database.Account.Find(ctx.Context(), bson.M{})
	//verificar si hay un error al obtener los datos
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error getting accounts",
			"error":   err,
		})
	}

	var accounts []bson.M
	//Parsear los datos obtenidos al arreglo de transactions
	if err = cursor.All(ctx.Context(), &accounts); err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error getting accounts",
			"error":   err,
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Accounts found",
		"data":    accounts,
	})
}

/*
	AddAccount - Función que se encarga de agregar una cuenta
	@Param ctx *fiber.Ctx - Contexto de la petición
	@Return error - Error en caso de que la petición falle
*/
func AddAccount(ctx *fiber.Ctx) error {
	account := new(models.Account)
	//Asignamos un UUID
	account.IdAccount = uuid.New()
	account.CbuAccount = uuid.New()

	//Parseamos a la estructura
	if err := ctx.BodyParser(account); err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing data",
			"error":   err,
		})
	}

	//Creamos un validator y validamos la Estructura
	validator := validator.New()
	if err := validator.Struct(account); err != nil { //<- Validamos la estructura de la cuenta
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error validating data",
		})
	}
	//Insertamos los datos
	_, err := database.Account.InsertOne(ctx.Context(), account)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error inserting account",
			"error":   err,
		})
	}

	log.Println(account.IdAccount)

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Account inserted",
		"data":    account,
	})
}

/*
	GetAccountById - Función que se encarga de obtener una cuenta por su ID
	@Param ctx *fiber.Ctx - Contexto de la petición
	@Return error - Error en caso de que la petición falle
*/
func GetAccountById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Convertir la cadena en un ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			// Manejar el error si no se puede convertir la cadena en ObjectId
			"status":  400,
			"message": "Invalid ID",
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

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Account found",
		"data":    account,
	})
}

/*
	UpdateAccountById - Función que se encarga de actualizar una cuenta por su ID
	@Param ctx *fiber.Ctx - Contexto de la petición
	@Return error - Error en caso de que la petición falle
*/
func UpdateAccountById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			// Manejar el error si no se puede convertir la cadena en ObjectId
			"status":  400,
			"message": "Invalid ID",
			"error":   err.Error(),
		})
	}

	account := new(models.Account)

	if err := ctx.BodyParser(account); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing data",
			"error":   err,
		})
	}

	_, err = database.Account.UpdateOne(ctx.Context(), bson.M{"_id": objID}, bson.D{{"$set", account}})

	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":   fiber.StatusInternalServerError,
			"message":  "Error updating account",
			"fielData": account,
			"error":    err,
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Account updated",
		"data":    account,
	})
}

/*
	DeleteAccountById - Función que se encarga de eliminar una cuenta por su ID
	@Param ctx *fiber.Ctx - Contexto de la petición
	@Return error - Error en caso de que la petición falle
*/
func DeleteAccountById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			// Manejar el error si no se puede convertir la cadena en ObjectId
			"status":  400,
			"message": "Invalid ID",
			"error":   err.Error(),
		})
	}

	_, err = database.Account.DeleteOne(ctx.Context(), bson.M{"_id": objID})

	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error deleting account",
			"error":   err,
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Account deleted",
	})
}



