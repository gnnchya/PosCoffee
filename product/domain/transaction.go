package domain

type Menu struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Code 			int 	 `json:"code"`
	Err 			error 	 `json:"err"`
}

type Cart struct{
	Menu		Menu 	`bson:"menu" json:"menu"`
	Amount 		int64   `bson:"amount" json:"amount"`
	Option 		string  `bson:"option" json:"option"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

type CreateOrderStruct struct {
	ID         		string   		`bson:"_id" json:"id"`
	CustomerID 		string  		`bson:"customer_id" json:"customer_id"`
	Cart			[]Cart   	`bson:"amount" json:"amount"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		GeoJson      	`bson:"destination" json:"destination"`
	Time			int64      		`bson:"date_time" json:"date_time"`
	Code 			int 			`json:"code"`
	Err 			error 			`json:"err"`
}

type DeleteOrderStruct struct {
	ID 		string	`bson:"_id" json:"id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateOrderStruct struct {
	ID         		string   		`bson:"_id" json:"id"`
	CustomerID 		string  		`bson:"customer_id" json:"customer_id"`
	Cart			[]MenuStruct   	`bson:"amount" json:"amount"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		GeoJson      	`bson:"destination" json:"destination"`
	Time			int64      		`bson:"date_time" json:"date_time"`
	Code 			int 			`json:"code"`
	Err 			error 			`json:"err"`
}


type ReadOrderStruct struct {
	ID string `bson:"_id" json:"id"`

}

type ReadOrderByPageStruct struct {
	PerPage int
	Page    int
}
