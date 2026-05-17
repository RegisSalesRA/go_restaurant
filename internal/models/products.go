package models

type Product struct {
	ID            int
	CategoryId    int
	Name          string
	Price         int
	StockQuantity int
	FinalCounted  int
	Active        bool
}
