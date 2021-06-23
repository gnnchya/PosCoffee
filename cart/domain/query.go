package domain

type Menu struct{
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
	Code 			int 	 `bson:"code" json:"code"`
	Err 			error 	 `bson:"err" json:"err"`
}

type CreateStruct struct {
	ID 				string  `bson:"_id" json:"_id"`
	CustomerID 		string  `bson:"customer_id" json:"customer_id"`
	Menu			[]Menu 	`bson:"menu" json:"menu"`
	Code 			int 	 `bson:"code" json:"code"`
	Err 			error 	 `bson:"err" json:"err"`
}

type MenuTest struct{
	ID         		string   `bson:"_id"`
	Category       	[]string   `bson:"category"`
	Name 			string   `bson:"name"`
	Ingredient 		[]string `bson:"ingredient"`
	Price      		int64    `bson:"price"`
	Available 		bool	 `bson:"available"`
	Amount 			int64    `bson:"amount"`
	Option 			string   `bson:"option"`
}

type CreateStructTest struct {
	ID 				string  `bson:"_id"`
	CustomerID 		string  `bson:"customer_id"`
	Menu			[]MenuTest 	`bson:"menu"`
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
	Code 			int 	`json:"code"`
	Err 			error	`json:"err"`
}

type ReadStruct struct {
	ID string `bson:"_id" json:"id"`
}

type ReadAllStruct struct {
	PerPage int
	Page    int
}

