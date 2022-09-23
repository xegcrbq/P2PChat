package cmd

import "github.com/xegcrbq/P2PChat/models"

type CreateOrderByOrder struct {
	Order *models.Order
}
type ReadOrderByOrderId struct {
	OrderId int32
}
type ReadOrdersByBuyerId struct {
	BuyerId int32
}
type ReadOrdersBySellerId struct {
	SellerId int32
}
