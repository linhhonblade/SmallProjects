package domain

import (
	"github.com/google/uuid"
)

type OrderLine struct {
	id        uuid.UUID
	orderId   uuid.UUID
	productId uuid.UUID
	quantity  float32
}

func (o OrderLine) Id() uuid.UUID {
	return o.id
}

func (o OrderLine) OrderId() uuid.UUID {
	return o.orderId
}

func (o OrderLine) ProductId() uuid.UUID {
	return o.productId
}

func (o OrderLine) Quantity() float32 {
	return o.quantity
}

func NewOrderLine(id uuid.UUID, orderId uuid.UUID, productId uuid.UUID, quantity float32) (*OrderLine, error) {
	// TODO: Add validation
	return &OrderLine{id: id, orderId: orderId, productId: productId, quantity: quantity}, nil
}
