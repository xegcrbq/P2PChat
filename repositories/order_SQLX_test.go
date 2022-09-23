package repositories

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/db"
	"github.com/xegcrbq/P2PChat/models/cmd"
	"testing"
)

var orderRepo = NewOrderRepoSQLX(db.ConnectSQLXTest())

func TestOrderRepoSQLX(t *testing.T) {
	fmt.Println(orderRepo.ReadOrderByOrderId(&cmd.ReadOrderByOrderId{OrderId: 1}))
	fmt.Println(orderRepo.ReadOrdersBySellerId(&cmd.ReadOrdersBySellerId{SellerId: 1}))
	fmt.Println(orderRepo.ReadOrdersBySellerId(&cmd.ReadOrdersBySellerId{SellerId: 2}))
	fmt.Println(orderRepo.ReadOrdersByBuyerId(&cmd.ReadOrdersByBuyerId{BuyerId: 1}))
	fmt.Println(orderRepo.ReadOrdersByBuyerId(&cmd.ReadOrdersByBuyerId{BuyerId: 2}))
}
