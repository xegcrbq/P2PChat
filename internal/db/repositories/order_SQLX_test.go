package repositories

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"testing"
)

var orderRepo = NewOrderRepoSQLX(db.ConnectSQLXTest())

func TestOrderRepoSQLX(t *testing.T) {
	fmt.Println(orderRepo.ReadOrderByOrderId(&commands.ReadOrderByOrderId{OrderId: 1}))
	fmt.Println(orderRepo.ReadOrdersBySellerId(&commands.ReadOrdersBySellerId{SellerId: 1}))
	fmt.Println(orderRepo.ReadOrdersBySellerId(&commands.ReadOrdersBySellerId{SellerId: 2}))
	fmt.Println(orderRepo.ReadOrdersByBuyerId(&commands.ReadOrdersByBuyerId{BuyerId: 1}))
	fmt.Println(orderRepo.ReadOrdersByBuyerId(&commands.ReadOrdersByBuyerId{BuyerId: 2}))
}
