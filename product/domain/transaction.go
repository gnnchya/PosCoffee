package domain

type Ingredient struct{
	IngredientName    string   `bson:"item_name" json:"item_name"`
	Amount      		int64    `bson:"amount" json:"amount"`
}

type Menu struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string  `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Amount 			int64    `bson:"amount" json:"amount"`
	Option 			string   `bson:"option" json:"option"`
}

type Cart struct{
	ID 			string  `bson:"_id" json:"_id"`
	CustomerID 	string  `bson:"customer_id" json:"customer_id"`
	Menu		[]Menu 	`bson:"menu" json:"menu"`
	TotalPrice	int64	`bson:"total_price"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

type CreateOrderStruct struct {
	ID         		string   		`bson:"_id" json:"_id"`
	Cart			Cart   			`bson:"cart" json:"cart"`
	Finished		bool     		`bson:"finished" json:"finished"`
	Price	     	int64   		`bson:"price" json:"price"`
	PaymentMethod	string			`bson:"payment_method" json:"payment_method"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		GeoJson      	`bson:"destination" json:"destination"`
	Time			int64      		`bson:"time" json:"time"`
	TotalCost 		int64			`bson:"total" json:"total"`
}

type DeleteOrderStruct struct {
	ID 		string	`bson:"_id" json:"_id"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateOrderStruct struct {
	ID         		string   		`bson:"_id" json:"_id"`
	Finished		bool     		`bson:"finished" json:"finished"`
	TypeOfOrder 	string 			`bson:"type" json:"type"`
	Destination		GeoJson      	`bson:"destination" json:"destination"`
}


type ReadOrderStruct struct {
	ID string `bson:"_id" json:"_id"`

}

type ReadOrderByPageStruct struct {
	PerPage int
	Page    int
}

type CalculateCost struct{
	ItemName         	string   `bson:"item_name"`
	CostPerUnit      	int64    `bson:"cost_per_unit"`
}

type TotalSale struct{
	ID struct{
		ID 		string		`bson:"_id" json:"_id"`
		Name 	string		`bson:"name" json:"name"`
		Price 	int64		`bson:"price" json:"price"`
	}	`bson:"_id" json:"_id"`
	TotalAmount int64		`bson:"total_amount" json:"total_amount"`
	TotalSales 	int64		`bson:"total_sales" json:"total_sales"`
}

type OldMenu struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string  `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
}