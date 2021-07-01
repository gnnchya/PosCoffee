package domain

type User struct {
	Name string `bson:"name"`
}

type SearchValue struct {
	Value string `bson:"value"`
	Type  string `bson:"type"`
}

type ReportValue struct{
	From int64
	Until int64
}

type ReportOrder struct{
	Field string
	Order string
}