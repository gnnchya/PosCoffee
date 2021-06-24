package domain

type User struct {
	Name string `bson:"name"`
}

type SearchValue struct {
	Type 	string	`bson:"type"`
	Value 	string `bson:"value"`
}
