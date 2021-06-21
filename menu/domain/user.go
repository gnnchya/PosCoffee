package domain

type User struct {
	Name string `bson:"name"`
}

type SearchValue struct {
	Value string `bson:"value"`
}
