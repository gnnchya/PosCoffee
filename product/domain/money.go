package domain

type CreateMoneyStruct struct {
	ID 			string		`bson:"-id"`
	Value   	int64   	`bson:"value"`
	Amount 		int64  		`bson:"amount"`
	Currency	string   	`bson:"currency"`
}

type DeleteMoneyStruct struct {
	ID 		string	`bson:"-id" json:"_id"`
	Value   int64   `bson:"value" json:"value"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateMoneyStruct struct {
	ID 		string	`bson:"-id" json:"_id"`
	Value   int64   `bson:"value" json:"value"`
	Amount 	int64  	`bson:"amount" json:"amount"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}


type ReadMoneyStruct struct {
	ID 		string	`bson:"-id" json:"_id"`
}

