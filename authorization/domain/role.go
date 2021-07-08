package domain

type Roles struct {
	ID          string    `bson:"_id,omitempty"`
	Name        *Lang     `bson:"name"`
	Permissions []*string `bson:"permissions"`
	CreatedAt   int64     `bson:"createdAt,omitempty"`
	UpdatedAt   int64     `bson:"updatedAt,omitempty"`
	DeletedAt   *int64    `bson:"deletedAt,omitempty"`
}

type Lang struct {
	Th string `bson:"th"`
	En string `bson:"en"`
}