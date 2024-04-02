package models

import (
	"github.com/google/uuid"
)

// Account - Estructura de la cuenta
type Transaction struct {
	IdTransaction   uuid.UUID `json:"idTransaction" bson:"idTransaction" form:"idTransaction" query:"idTransaction"`
	SenderAccount   string    `json:"senderAccount" bson:"senderAccount" form:"senderAccount" query:"senderAccount"`
	ReceiverAccount string    `json:"receiverAccount" bson:"receiverAccount" form:"receiverAccount" query:"receiverAccount"`
	Amount          float64   `json:"amount" bson:"amount" form:"amount" query:"amount" validate:"required,min=0,max=1000000000"`
	Date            string    `json:"date" bson:"date" form:"date" query:"date" validate:"required"`
	Time            string    `json:"time" bson:"time" form:"time" query:"time" validate:"required"`
}
// TrasactionDTO - Estructura de la transferencia enviada
type TransactionDTO struct {
	SenderAccount   string  `json:"senderAccount" bson:"senderAccount" form:"senderAccount" query:"senderAccount"`
	ReceiverAccount string  `json:"receiverAccount" bson:"receiverAccount" form:"receiverAccount" query:"receiverAccount"`
	TypeTransaction string  `json:"typeTransaction" bson:"typeTransaction" validate:"required,oneof=transferencia deposito retiro"`
	Amount          float64 `json:"amount" bson:"amount" form:"amount" query:"amount" validate:"required,min=0.00,max=1000000000.00"`
}
// AccountDTO - Estructura de la cuenta enviada
type AccountDTO struct {
	Bank        string  `json:"bank" bson:"bank" form:"bank" query:"bank" validate:"required"`
    TypeAccount string  `json:"typeAccount" bson:"typeAccount" form:"typeAccount" query:"typeAccount" validate:"required"`
    FullName    string  `json:"fullname" bson:"fullname" form:"fullname" query:"fullname" validate:"required"`
    Dni         string  `json:"dni" bson:"dni" form:"dni" query:"dni" validate:"required,min=0,max=999999999999999"`
    Email       string  `json:"email" bson:"email" form:"email" query:"email" validate:"required,email"`
    Phone       string  `json:"phone" bson:"phone" form:"phone" query:"phone" validate:"required"`
    Balance     float64 `json:"balance" bson:"balance" form:"balance" query:"balance" validate:"required,min=0,max=1000000000",float64`
}
