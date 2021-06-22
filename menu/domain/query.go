package domain

type CreateStruct struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type DeleteStruct struct {
	ID string `bson:"_id" json:"id"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type UpdateStruct struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient"`
	Price      		int64    `bson:"price" json:"price"`
	Available 		bool	 `bson:"available" json:"available"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type ReadStruct struct {
	ID string `bson:"_id" json:"id"`

}

type ReadAllStruct struct {
	PerPage int
	Page    int
}