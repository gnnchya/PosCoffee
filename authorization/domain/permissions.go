package domain

type Permissions struct {
	ID        string `bson:"_id,omitempty"`
	Method    string `bson:"method"`
	Endpoint  string `bson:"endpoint"`
	CreatedAt int64  `bson:"createdAt,omitempty"`
	UpdatedAt int64  `bson:"updatedAt,omitempty"`
	DeletedAt *int64 `bson:"deletedAt,omitempty"`
}
