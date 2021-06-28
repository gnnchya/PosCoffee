package domain

type CreateStockStruct struct {
	ID         		string   	`bson:"_id" json:"_id"`
	ItemName       	string   	`bson:"item_name" json:"item_name"`
	Category 		string  	`bson:"category" json:"category"`
	Amount			int64   	`bson:"amount" json:"amount"`
	Unit     		string   	`bson:"unit" json:"unit"`
	CostPerUnit		int64      	`bson:"cost_per_unit" json:"cost_per_unit"`
	EXPDate     	int64   	`bson:"exp_date" json:"exp_date"`
	ImportDate      int64   	`bson:"import_date" json:"import_date"`
	Supplier 		string 		`bson:"supplier" json:"supplier"`
	TotalCost		int64      	`bson:"total_cost" json:"total_cost"`
	TotalAmount		int64      	`bson:"total_amount" json:"total_amount"`
	Status 			string		`bson:"status" json:"status"`
}

type DeleteStockStruct struct {
	ID 		string	`bson:"_id" json:"id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateStockStruct struct {
	ID      string  `bson:"_id" json:"id" validate:"required"`
	Amount	int64   `bson:"amount" json:"amount" validate:"required"`
	Code 	int 	`json:"code"`
	Err 	error   `json:"err"`
}

type ReportStruct struct{
	Transaction		[]CreateOrderStruct		`bson:"transaction" json:"transaction"`
	UsedStock		[]CreateStockStruct			`bson:"used_stock" json:"used_stock"`
	OutOfStock		[]CreateStockStruct			`bson:"out_of_stock" json:"out_of_stock"`
	TotalIncome		int64					`bson:"total_income" json:"total_income"`
	TotalCost		int64					`bson:"total_cost" json:"total_cost"`
	TotalProfit		int64					`bson:"total_profit" json:"total_profit"`
}