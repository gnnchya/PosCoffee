package domain

type Users struct {
	ID             string       `bson:"_id,omitempty"`
	Prefix         *Lang        `bson:"prefix"`
	FirstName      *Lang        `bson:"firstName"`
	LastName       *Lang        `bson:"lastName"`
	SignUpChannel  string       `bson:"signUpChannel"`
	Email          string       `bson:"email"`
	MobileNumber   string       `bson:"mobileNumber"`
	IdentifyType   string       `bson:"identifyType"`
	IdentifyNumber string       `bson:"identifyNumber"`
	Password       string       `bson:"password"`
	PassCode       string       `bson:"passCode"`
	MetaData       *MetaData    `bson:"metaData"`
	RoleID         []string     `bson:"roleID"`
	MemberGroup    *MemberGroup `bson:"memberGroup"`
	CreatedAt      int64        `bson:"createdAt"`
	UpdatedAt      int64        `bson:"updatedAt"`
	DeletedAt      *int64       `bson:"deletedAt"`
}

type Lang struct {
	Th string `bson:"th"`
	En string `bson:"en"`
}

type MetaData struct {
	Gender    string  `bson:"gender"`
	BirthDate int     `bson:"birthDate"`
	YearOnly  bool    `bson:"yearOnly"`
	Weight    float64 `bson:"weight"`
	Height    float64 `bson:"height"`
}

type Roles struct {
	ID         string        `bson:"_id,omitempty"`
	Name       *Lang         `bson:"name"`
	Permission []*Permission `bson:"permission"`
}

type Permission struct {
	ID       string `bson:"_id,omitempty"`
	Method   string `bson:"method"`
	Endpoint string `bson:"endpoint"`
}

type UsersSoftDelete struct {
	DeletedAt int64 `bson:"deletedAt"`
}

type MemberGroup struct {
	Name    string   `bson:"name"`
	Members []string `bson:"members"`
}

type ForgotPassword struct {
	Password  string `bson:"password"`
	UpdatedAt int64  `bson:"updatedAt"`
}

type SearchValue struct {
	Value string `bson:"value"`
}
