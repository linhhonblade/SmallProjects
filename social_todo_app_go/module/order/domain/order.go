package domain

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	id         uuid.UUID
	userId     uuid.UUID
	expireAt   time.Time
	status     OrderStatus
	orderLines []OrderLine
}

func NewOrder(id uuid.UUID, userId uuid.UUID, expireAt time.Time, status OrderStatus, orderLines []OrderLine) (*Order, error) {
	// TODO: Add validation
	return &Order{id: id, userId: userId, expireAt: expireAt, status: status, orderLines: orderLines}, nil
}

func (o Order) Id() uuid.UUID {
	return o.id
}

func (o Order) UserId() uuid.UUID {
	return o.userId
}

func (o Order) ExpireAt() time.Time {
	return o.expireAt
}

func (o Order) Status() OrderStatus {
	return o.status
}

func (o Order) OrderLines() []OrderLine {
	return o.orderLines
}

type OrderStatus int

const (
	OrderStatusDraft OrderStatus = iota
	OrderStatusConfirmed
	OrderStatusCancelled
	OrderStatusCompleted
)

func (s OrderStatus) String() string {
	return [...]string{"Draft", "Confirmed", "Cancelled", "Completed"}[s]
}

func GetOrderStatus(status string) OrderStatus {
	switch status {
	case "Draft":
		return OrderStatusDraft
	case "Confirmed":
		return OrderStatusConfirmed
	case "Cancelled":
		return OrderStatusCancelled
	case "Completed":
		return OrderStatusCompleted
	default:
		return OrderStatusDraft
	}
}
