package domain

type InsertQ struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	string   `bson:"category" json:"category"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type DeleteQ struct {
	ID string `bson:"_id" json:"id"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type UpdateQ struct {
	ID         		string   `bson:"_id" json:"id" validate:"required"`
	Category       	string   `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]string `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
	Code int `json:"code"`
	Err error `json:"err"`
}

type ViewQ struct {
	ID string `bson:"_id" json:"id"`

}

type ViewByPageQ struct {
	PerPage int
	Page    int
}