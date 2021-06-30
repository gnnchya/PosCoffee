package domain

type Ingredient struct{
	IngredientName    string   `bson:"item_name" json:"item_name"`
	Amount      		int64    `bson:"amount" json:"amount"`
}

type CreateStruct struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string  `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
}

type DeleteStruct struct {
	ID string `bson:"_id" json:"_id"`
	Code int  `json:"code"`
	Err error `json:"err"`
}

type UpdateStruct struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
}

type ReadStruct struct {
	ID string `bson:"_id" json:"_id"`
}

type ReadAllStruct struct {
	PerPage int
	Page    int
}