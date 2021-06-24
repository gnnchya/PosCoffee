package domain

type CreateStockStruct struct {
	ID         		string   	`bson:"_id"`
	ItemName       	string   	`bson:"item_name"`
	Category 		string  	`bson:"category"`
	Amount			int64   	`bson:"amount"`
	Unit     		string   	`bson:"unit"`
	CostPerUnit		int64      	`bson:"cost_per_unit"`
	EXPDate     	int64   	`bson:"exp_date"`
	ImportDate      int64   	`bson:"import_date"`
	Supplier 		string 		`bson:"supplier"`
	TotalCost		int64      	`bson:"total_cost"`
	TotalAmount		int64      	`bson:"total_amount"`
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