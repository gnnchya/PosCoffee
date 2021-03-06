package domain

type UserStruct struct{
	ID  			string				`bson:"_id" json:"id"`
	UID  			string				`bson:"uid" json:"uid"`
	Username		string				`bson:"username" json:"username"`
	Password 		string				`bson:"password" json:"password"`
	MetaData		*MetaDataStruct		`bson:"meta_data" json:"meta_data"`
	RoleID			[]string		`bson:"role_id" json:"role_id"`
	Verify 			bool 				`bson:"verify" json:"verify"`
	CreatedAt 		int64				`bson:"created_at" json:"created_at"`
	UpdatedAt 		int64				`bson:"updated_at" json:"updated_at"`
	DeletedAt 		int64				`bson:"deleted_at" json:"deleted_at"`
}

type MetaDataStruct struct{
	Prefix 			*LanguageStruct		`bson:"prefix" json:"prefix"`
	Name 			*LanguageStruct		`bson:"name" json:"name"`
	Lastname 		*LanguageStruct		`bson:"lastname" json:"lastname"`
	Email 			string				`bson:"email" json:"email"`
	MobileNumber	string				`bson:"mobile_number" json:"mobile_number"`
	LineID			string				`bson:"line_id" json:"line_id"`
	Facebook 		string				`bson:"facebook" json:"facebook"`
	BankAccount		string				`bson:"bank_account" json:"bank_account"`
	Gender 			string				`bson:"gender" json:"gender"`
	BirthDate		int64				`bson:"birth_date" json:"birth_date"`
}

type LanguageStruct struct{
	TH			string		`bson:"th" json:"th"`
	EN 			string		`bson:"en" json:"en"`
}

type ConsumerStruct struct {
	ID 				string 	`bson:"_id" json:"id"`
	ClientID		string 	`bson:"client_id" json:"client_id"`
	ClientSecret	string 	`bson:"client_secret" json:"client_secret"`
	RedirectUri     string `bson:"redirectUri"`
	Scope           string 	`bson:"scope" json:"scope"`
	CreatedAt       int64  	`bson:"createdAt" json:"createdAt"`
}

type TokenStruct struct {
	ID 				string 	`bson:"_id" json:"id"`
	UID 			string	`bson:"uid" json:"uid"`
	AccessToken 	string	`bson:"access_token" json:"access_token"`
	AccessExpire 	int64	`bson:"access_expire" json:"access_expire"`
	RefreshToken	string	`bson:"refresh_token" json:"refresh_token"`
	RefreshExpire	int64	`bson:"refresh_expire" json:"refresh_expire"`
	CreatedAt 		int64	`bson:"created_at" json:"created_at"`
}


type PermissionStruct struct{
	ID			string
	Method 		string
	End			string
}

type RoleStruct struct{
	ID 				string 			`bson:"_id" json:"id"`
	Name 			LanguageStruct		`bson:"name" json:"name"`
	Permissions 	[]PermissionStruct	`bson:"permissions" json:"permissions"`
}


type UsersSoftDeleteStruct struct {
	DeletedAt int64 `bson:"deleted_at" json:"deleted_at"`
}

type MemberGroupStruct struct {
	Name    string   `bson:"name" json:"name"`
	Members []string `bson:"members" json:"members"`
}

type ForgotPasswordStruct struct {
	Password  string `bson:"password" json:"password"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
}

type SearchValueStruct struct {
	Value string `bson:"value" json:"value"`
}

type Paginator struct {
	Total   int
	CurPage int
	PerPage int
	HasMore bool
}

type PageOption struct {
	Page    int      `form:"page" validate:"min=0"`
	PerPage int      `form:"per_page" validate:"min=0"`
	Filters []string `form:"filters"`
	Sorts   []string `form:"sorts"`
}

type SetOpParam struct {
	Filters      []string
	SetFieldName string
	Item         interface{}
}

type TokenView struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expired      int64  `json:"expired"`
}