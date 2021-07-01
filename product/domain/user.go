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
	Format string
}

type ReportOrder struct{
	Field string
	Order string
	Format string
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