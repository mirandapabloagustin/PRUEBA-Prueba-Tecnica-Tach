package models

import (
	"github.com/google/uuid"
)

/*
Account - Estructura que define la cuenta
*/
type Account struct {
	IdAccount   uuid.UUID `json:"idAccount" bson:"idAccount" form:"idAccount" query:"idAccount" validate:"required"` //<- Implementamos la validación de campos
	CbuAccount  uuid.UUID `json:"cbu" bson:"cbu" form:"cbu" query:"cbu" validate:"required"`
	Bank        string    `json:"bank" bson:"bank" form:"bank" query:"bank" validate:"required"`
	TypeAccount string    `json:"typeAccount" bson:"typeAccount" form:"typeAccount" query:"typeAccount" validate:"required"`
	FullName    string    `json:"fullname" bson:"fullname" form:"fullname" query:"fullname" validate:"required"`
	Dni         string    `json:"dni" bson:"dni" form:"dni" query:"dni" validate:"required,min=0,max=999999999999999"`
	Email       string    `json:"email" bson:"email" form:"email" query:"email" validate:"required,email"`
	Phone       string    `json:"phone" bson:"phone" form:"phone" query:"phone" validate:"required"`
	Balance     float64   `json:"balance" bson:"balance" form:"balance" query:"balance" validate:"required,min=0,max=1000000000",float64`
}

/*
Transaction - Estructura que define la transacción
*/

type AccountDTO struct {
	Bank        string  `json:"bank" bson:"bank" form:"bank" query:"bank" validate:"required"`
	TypeAccount string  `json:"typeAccount" bson:"typeAccount" form:"typeAccount" query:"typeAccount" validate:"required"`
	FullName    string  `json:"fullname" bson:"fullname" form:"fullname" query:"fullname" validate:"required"`
	Dni         string  `json:"dni" bson:"dni" form:"dni" query:"dni" validate:"required,min=0,max=999999999999999"`
	Email       string  `json:"email" bson:"email" form:"email" query:"email" validate:"required,email"`
	Phone       string  `json:"phone" bson:"phone" form:"phone" query:"phone" validate:"required"`
	Balance     float64 `json:"balance" bson:"balance" form:"balance" query:"balance" validate:"required,min=0,max=1000000000",float64`
}

/*
TransactionDTO - Estructura que define la transacción dto
*/
type TransactionDTO struct {
	SenderAccount   string  `json:"senderAccount" bson:"senderAccount" form:"senderAccount" query:"senderAccount"`
	ReceiverAccount string  `json:"receiverAccount" bson:"receiverAccount" form:"receiverAccount" query:"receiverAccount"`
	TypeTransaction string  `json:"typeTransaction" bson:"typeTransaction" validate:"required,oneof=transferencia deposito retiro"`
	Amount          float64 `json:"amount" bson:"amount" form:"amount" query:"amount" validate:"required,min=0.00,max=1000000000.00"`
}
