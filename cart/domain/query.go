package domain

type Ingredient struct{
	ItemName    string   `bson:"item_name" json:"item_name"`
	Amount      		int64    `bson:"amount" json:"amount"`
}

type Menu struct{
	ID         		string   `bson:"_id" json:"_id"`
	Category       	[]string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
}

type CreateStruct struct {
	ID 				string  `bson:"_id" json:"_id"`
	CustomerID 		string  `bson:"customer_id" json:"customer_id"`
	Menu			[]Menu	`bson:"menu" json:"menu"`
	TotalPrice		int64	`bson:"total_price" json:"total_price"`
}

type DeleteStruct struct {
	ID 		string 	`bson:"_id" json:"id"`
}

type UpdateStruct struct {
	ID 				string  `bson:"_id" json:"_id"`
	CustomerID 		string  `bson:"customer_id" json:"customer_id"`
	Menu			[]Menu 	`bson:"menu" json:"menu"`
	TotalPrice		int64	`bson:"total_price" json:"total_price"`
}

type ReadStruct struct {
	ID string `bson:"_id" json:"id"`
}

type ReadAllStruct struct {
	PerPage int
	Page    int
}

