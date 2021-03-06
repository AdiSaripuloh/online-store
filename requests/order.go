package requests

import uuid "github.com/satori/go.uuid"

type CreateOrder struct {
	UserID     uuid.UUID `form:"userID"`
	GrandTotal float64   `form:"grantTotal"`
	Items      []Items   `form:"items"`
}

type PayOrder struct {
	OrderID uuid.UUID `form:"orderID"`
	Amount  float64   `form:"amount"`
}
