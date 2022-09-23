package models

type Order struct {
	OrderId            int32  `db:"orderid"`
	SellerId           int32  `db:"sellerid"`
	SellerTicker       string `db:"sellerticker"`
	SellerAmount       string `db:"selleramount"`
	SellerVoteComplete bool   `db:"sellervotecomplete"`
	BuyerId            int32  `db:"buyerid"`
	BuyerTicker        string `db:"buyerticker"`
	BuyerAmount        string `db:"buyeramount"`
	BuyerVoteComplete  bool   `db:"buyervotecomplete"`
	IsCompleted        bool   `db:"iscompleted"`
}
