package domain

type User struct {
	Name string `bson:"name"`
}

type SearchValue struct {
	Value string `bson:"value"`
	Type  string `bson:"type"`
}

type ReportValue struct{
	From 	int64	`bson:"from" json:"from"`
	Until 	int64	`bson:"until" json:"until"`
	Format 	string	`bson:"format" json:"format"`
}

type ReportOrder struct{
	Field 	string	`bson:"field" json:"field"`
	Order 	string	`bson:"order" json:"order"`
	Format 	string	`bson:"format" json:"format"`
}

type ReadCategoryByPageStruct struct {
	Category 	string 	`bson:"category" json:"category"`
	PerPage 	int		`bson:"per_page" json:"per_page"`
	Page    	int		`bson:"page" json:"page"`
}

type ReadNameByPageStruct struct {
	ItemName 	string 	`bson:"item_name" json:"item_name"`
	PerPage 	int		`bson:"per_page" json:"per_page"`
	Page    	int		`bson:"page" json:"page"`
}