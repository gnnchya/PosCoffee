package domain

type CreateStockStruct struct {
	ID         		string   	`bson:"_id" json:"id"`
	ItemName       	string   	`bson:"item_name" json:"item_name" validate:"required"`
	Category 		string  	`bson:"category" json:"category" validate:"required"`
	Amount			int64   	`bson:"amount" json:"amount" validate:"required"`
	Unit     		string   	`bson:"unit" json:"unit" validate:"required"`
	CostPerUnit		int64      	`bson:"cost_per_unit" json:"cost_per_unit"`
	EXPDate     	int64   	`bson:"exp_date" json:"exp_date" validate:"required"`
	ImportDate      int64   	`bson:"import_date" json:"import_date" validate:"required"`
	Supplier 		string 		`bson:"supplier" json:"supplier"`
	TotalCost		int64      	`bson:"total_cost" json:"total_cost" validate:"required"`
	TotalAmount		int64      	`bson:"total_amount" json:"total_amount" validate:"required"`
	Code 			int 		`json:"code"`
	Err 			error 		`json:"err"`
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