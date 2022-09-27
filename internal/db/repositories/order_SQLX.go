package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
)

type OrderRepoSQLX struct {
	db *sqlx.DB
}

func NewOrderRepoSQLX(db *sqlx.DB) *OrderRepoSQLX {
	return &OrderRepoSQLX{
		db: db,
	}
}

func (r *OrderRepoSQLX) CreateOrderByOrder(cmd *commands.CreateOrderByOrder) error {
	order := cmd.Order
	_, err := r.db.Exec(`
		insert into 
			orders(sellerid, sellerticker, selleramount, buyerid, buyerticker, buyeramount) 
		VALUES($1, $2, $3, $4, $5, $6);`,
		order.SellerId, order.SellerTicker, order.SellerAmount, order.BuyerId, order.BuyerTicker, order.BuyerAmount)
	return err
}

func (r *OrderRepoSQLX) ReadOrderByOrderId(cmd *commands.ReadOrderByOrderId) (*models.Order, error) {
	var order models.Order
	err := r.db.Get(&order,
		`select * from orders where orderid=$1;`, cmd.OrderId)
	return &order, err
}

func (r *OrderRepoSQLX) ReadOrdersByBuyerId(cmd *commands.ReadOrdersByBuyerId) (*[]models.Order, error) {
	var orders []models.Order
	err := r.db.Select(&orders,
		`select * from orders where buyerid=$1;`, cmd.BuyerId)
	return &orders, err
}

func (r *OrderRepoSQLX) ReadOrdersBySellerId(cmd *commands.ReadOrdersBySellerId) (*[]models.Order, error) {
	var orders []models.Order
	err := r.db.Select(&orders,
		`select * from orders where sellerid=$1;`, cmd.SellerId)
	return &orders, err
}
