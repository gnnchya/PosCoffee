package domain

type Cart struct{
	Menu		Menu 	`bson:"menu" json:"menu"`
	Amount 		int64   `bson:"amount" json:"amount"`
	Option 		string  `bson:"option" json:"option"`
}

type Menu struct{
	ID         		string   `bson:"_id" json:"id"`
	Category       	string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Code 			int 	 `json:"code"`
	Err 			error 	 `json:"err"`
}
type CreateStruct struct {
	ID         		string   	`bson:"_id" json:"id"`
	CustomerID 		string   	`bson:"customer_id" json:"customer_id"`
	Cart  			[]Cart  	`bson:"cart" json:"cart"`
	Purchased     	bool   		`bson:"purchased" json:"purchased"`
	Price  			int64    	`bson:"price" json:"price"`
	TypeOfOrder 	string 		`bson:"type_of_order" json:"type_of_order"`
	Destination    	string     	`bson:"destination" json:"destination"`
	Time			int64 		`bson:"time" json:"time"`
	Code 			int 		`json:"code"`
	Err 			error		`json:"err"`
}

type DeleteStruct struct {
	ID 		string 	`bson:"_id" json:"id"`
	Code	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateStruct struct {
	ID         		string   `bson:"_id" json:"id"`
	CustomerID 		string   `bson:"customer_id" json:"customer_id"`
	Cart  			[]struct{
			Menu	struct{
							ID         		string   `bson:"_id" json:"id"`
							Category       	string   `bson:"category" json:"category"`
							Name 			string   `bson:"name" json:"name" validate:"required"`
							Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
							Price      		int64    `bson:"price" json:"price"`
							Available 		bool	 `bson:"available" json:"available"`
							Code 			int 	 `json:"code"`
							Err				error 	 `json:"err"`
						   }`bson:"menu" json:"menu"`
			Amount 	int64   `bson:"amount" json:"amount"`
			Option 	string  `bson:"option" json:"option"`
	}   					`bson:"cart" json:"cart"`
	Purchased     	bool   	`bson:"purchased" json:"purchased"`
	Price  			int64   `bson:"price" json:"price"`
	TypeOfOrder 	string 	`bson:"type_of_order" json:"type_of_order"`
	Destination    	string  `bson:"destination" json:"destination"`
	Time			int64 	`bson:"time" json:"time"`
	Code 			int 	`json:"code"`
	Err 			error 	`json:"err"`
}
type ReadStruct struct {
	ID string `bson:"_id" json:"id"`

}

type ReadAllStruct struct {
	PerPage int
	Page    int
}

