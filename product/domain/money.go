package domain

type CreateMoneyStruct struct {
	Value   		int64   		`bson:"value" json:"value"`
	Amount 			int64  			`bson:"amount" json:"amount"`
	Currency		string   		`bson:"currency" json:"currency"`
	Code 			int 			`json:"code"`
	Err 			error 			`json:"err"`
}

type DeleteMoneyStruct struct {
	Value   int64   `bson:"value" json:"value"`
	Code 	int 	`json:"code"`
	Err 	error 	`json:"err"`
}

type UpdateMoneyStruct struct {
	Value   		int64   		`bson:"value" json:"value"`
	Amount 			int64  			`bson:"amount" json:"amount"`
	Code 			int 			`json:"code"`
	Err 			error 			`json:"err"`
}


type ReadMoneyStruct struct {
	Value   int64   `bson:"value" json:"value"`

}

type ReadMoneyByPageStruct struct {
	PerPage int
	Page    int
}
