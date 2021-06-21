package domain

type CreateStruct struct {
	ID         		string   	`bson:"_id" json:"id"`
	ItemName       	string   	`bson:"item_name" json:"item_name"`
	Category 		string  	`bson:"category" json:"category"`
	Amount			int64   	 	`bson:"amount" json:"amount"`
	Unit     		string   	`bson:"unit" json:"unit"`
	CostPerUnit		int      	`bson:"cost_per_unit" json:"cost_per_unit"`
	EXPDate     	int64   	`bson:"exp_date" json:"exp_date"`
	ImportDate      int64   	`bson:"import_date" json:"import_date"`
	Supplier 		string 		`bson:"supplier" json:"supplier"`
	TotalCost		int64      	`bson:"total_cost" json:"total_cost"`
	TotalAmount		int64      	`bson:"total_amount" json:"total_amount"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type DeleteStruct struct {
	ID 		string	`bson:"_id" json:"id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateStruct struct {
	ID         		string   	`bson:"_id" json:"id"`
	ItemName       	string   	`bson:"item_name" json:"item_name"`
	Category 		string  	`bson:"category" json:"category"`
	Amount			int64   	`bson:"amount" json:"amount"`
	Unit     		string   	`bson:"unit" json:"unit"`
	CostPerUnit		int      	`bson:"cost_per_unit" json:"cost_per_unit"`
	EXPDate     	int64   	`bson:"exp_date" json:"exp_date"`
	ImportDate      int64   	`bson:"import_date" json:"import_date"`
	Supplier 		string 		`bson:"supplier" json:"supplier"`
	TotalCost		int64      	`bson:"total_cost" json:"total_cost"`
	TotalAmount		int64      	`bson:"total_amount" json:"total_amount"`
	Code int `json:"code"`
	Err error `json:"err"`
}


type ReadStruct struct {
	ID string `bson:"_id" json:"id"`

}

type ReadByPageStruct struct {
	PerPage int
	Page    int
}
