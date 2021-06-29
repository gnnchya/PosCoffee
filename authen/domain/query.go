package domain

type DeleteStruct struct {
	ID 		string 	`bson:"_id" json:"id"`
}

type UpdateStruct struct {
	ID 				string  `bson:"_id" json:"-"`
	CustomerID 		string  `bson:"customer_id" json:"customer_id"`
	TotalPrice		int64	`bson:"total_price" json:"total_price"`
}

type ReadStruct struct {
	ID string `bson:"_id" json:"id"`
}

type ReadAllStruct struct {
	PerPage int
	Page    int
}

