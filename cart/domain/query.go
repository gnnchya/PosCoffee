package domain

type InsertQ struct {
	ID         		string   `bson:"_id" json:"id"`
	CustomerName    string   `bson:"customer_name" json:"customer_name" validate:"required"`
	CustomerID 		string   `bson:"customer_id" json:"customer_id"`
	Cart  			string   `bson:"actual_lastname" json:"actual_lastname"`
	Status     		string   `bson:"status" json:"status"`
	Price  			int64    `bson:"birth_date" json:"birth_date"`
	Branch     		string      `bson:"branch" json:"branch"`
	TypeOfOrder 	string `bson:"type_of_order" json:"type_of_order"`
	Destination    	string     `bson:"destination" json:"destination"`
	Code int `json:"code"`
	Err error `json:"err"`

}

type DeleteQ struct {
	ID string `bson:"_id" json:"id"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type UpdateQ struct {
	ID         		string   `bson:"_id" json:"id"`
	CustomerName    string   `bson:"customer_name" json:"customer_name" validate:"required"`
	CustomerID 		string   `bson:"customer_id" json:"customer_id"`
	Cart  			string   `bson:"actual_lastname" json:"actual_lastname"`
	Status     		string   `bson:"status" json:"status"`
	Price  			int64    `bson:"birth_date" json:"birth_date"`
	Branch     		string      `bson:"branch" json:"branch"`
	TypeOfOrder 	string `bson:"type_of_order" json:"type_of_order"`
	Destination    	string     `bson:"destination" json:"destination"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type ViewQ struct {
	ID string `bson:"_id" json:"id"`

}

type ViewByPageQ struct {
	PerPage int
	Page    int
}

