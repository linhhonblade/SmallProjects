package usecase

type createOrderUC struct {
	orderCmdRepo OrderCommandRepository
}

func NewCreateOrderUC(orderCmdRepo OrderCommandRepository) *createOrderUC {
	return &createOrderUC{orderCmdRepo: orderCmdRepo}
}
