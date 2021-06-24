package domain

type CreateStruct struct {
	ID         		string   	`bson:"_id"`
	ItemName       	string   	`bson:"item_name"`
	Category 		string  	`bson:"category"`
	Amount			int64   	`bson:"amount"`
	Unit     		string   	`bson:"unit"`
	CostPerUnit		int64      	`bson:"cost_per_unit" `
	EXPDate     	int64   	`bson:"exp_date"`
	ImportDate      int64   	`bson:"import_date"`
	Supplier 		string 		`bson:"supplier"`
	TotalCost		int64      	`bson:"total_cost"`
	TotalAmount		int64      	`bson:"total_amount"`
	Status 			string		`bson:"status"`
}

type DeleteStruct struct {
	ID 		string	`bson:"_id" json:"id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type Updatestruct struct {
	ID      string  `bson:"_id" json:"id" validate:"required"`
	Amount	int64   `bson:"amount" json:"amount" validate:"required"`
	Status 	string	`bson:"status" json:"status" validate:"required"`
	Code 	int 	`json:"code"`
	Err 	error   `json:"err"`
}

type ReadIDStruct struct {
	ID 		string  `bson:"_id" json:"id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type ReadCategoryByPageStruct struct {
	Category 	string 	`bson:"category" json:"category"`
	PerPage 	int		`bson:"per_page" json:"per_page"`
	Page    	int		`bson:"page" json:"page"`
	Code 		int 	`json:"code"`
	Err 		error 	`json:"err"`

}

type ReadNameByPageStruct struct {
	ItemName 	string 	`bson:"item_name" json:"item_name"`
	PerPage 	int		`bson:"per_page" json:"per_page"`
	Page    	int		`bson:"page" json:"page"`
	Code 		int 	`json:"code"`
	Err 		error 	`json:"err"`

}

type FilterStruct struct {
	SortBy	 	string	`bson:"sort_by" json:"sort_by"`
	Ascending	bool	`bson:"ascending" json:"ascending"`
	Active 		bool	`bson:"active" json:"active"`
}

type Ingredient struct{
	IngredientName      string   `bson:"ingredient_name" json:"ingredient_name"`
	Amount      		int64    `bson:"amount" json:"amount"`
}

type CalculateCost struct{
	ItemName         	string   `bson:"item_name"`
	CostPerUnit      	int64    `bson:"cost_per_unit"`
}