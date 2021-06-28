package domain

type CreateMoneyStruct struct {
	ID 			string		`bson:"_id" json:"-id"`
	Value   	int64   	`bson:"value" json:"value"`
	Amount 		int64  		`bson:"amount" json:"amount"`
	Currency	string   	`bson:"currency" json:"currency"`
}

type DeleteMoneyStruct struct {
	ID 		string	`bson:"_id" json:"_id"`
	Value   int64   `bson:"value" json:"value"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateMoneyStruct struct {
	ID 		string	`bson:"_id" json:"_id"`
	Value   int64   `bson:"value" json:"value"`
	Amount 	int64  	`bson:"amount" json:"amount"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}


type ReadMoneyStruct struct {
	ID 		string	`bson:"_id" json:"_id"`
}

type ChangeStruct struct{
	Value 	int64	`bson:"value" json:"value"`
	Amount 	int64	`bson:"amount" json:"amount"`
}