package requests

import uuid "github.com/satori/go.uuid"

type Items struct {
	ProductID uuid.UUID `form:"productID" json:"productID" binding:"required"`
	Quantity  int64     `form:"quantity" json:"quantity" binding:"required"`
}
