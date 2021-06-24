package domain

type Ingredient struct{
	IngredientName      string   `bson:"ingredient_name"`
	Amount      		int64    `bson:"amount"`
}

type Menu struct{
	ID         		string   `bson:"_id"`
	Category       	[]string   `bson:"category"`
	Name 			string   `bson:"name"`
	Ingredient 		[]Ingredient `bson:"ingredient"`
	Price      		int64    `bson:"price"`
	Available 		bool	 `bson:"available"`
	Amount 			int64    `bson:"amount"`
	Option 			string   `bson:"option"`
}

type CreateStruct struct {
	ID 				string  `bson:"_id"`
	CustomerID 		string  `bson:"customer_id"`
	Menu			[]Menu	`bson:"menu"`
	TotalPrice		int64	`bson:"total_price"`
}

type DeleteStruct struct {
	ID 		string 	`bson:"_id" json:"id"`
	Code	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateStruct struct {
	ID 				string  `bson:"_id" json:"_id"`
	CustomerID 		string  `bson:"customer_id" json:"customer_id"`
	Menu			[]Menu 	`bson:"menu" json:"menu"`
	TotalPrice		int64	`bson:"total_price"`
}

type ReadStruct struct {
	ID string `bson:"_id" json:"id"`
}

type ReadAllStruct struct {
	PerPage int
	Page    int
}

