package domain

type UserStruct struct{
	ID  			string				`bson:"_id" json:"id"`
	UID  			string				`bson:"uid" json:"uid"`
	Username		string				`bson:"username" json:"username"`
	Password 		string				`bson:"password" json:"password"`
	MetaData		metadataStruct		`bson:"meta_data" json:"meta_data"`
	Role			[]RoleStruct		`bson:"role_id" json:"role_id"`
	CreateAt 		int64				`bson:"create_at" json:"create_at"`
	UpdateAt 		int64				`bson:"update_at" json:"update_at"`
	DeleteAt 		int64				`bson:"delete_at" json:"delete_at"`
}

type metadataStruct struct{
	Prefix 			LanguageStruct		`bson:"prefix" json:"prefix"`
	Name 			LanguageStruct		`bson:"name" json:"name"`
	Lastname 		LanguageStruct		`bson:"lastname" json:"lastname"`
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
}

type TokenStruct struct {
	ID 				string 	`bson:"_id" json:"id"`
	UID 			string	`bson:"uid" json:"uid"`
	Token 			string	`bson:"token" json:"token"`
	Expire 			string	`bson:"expire" json:"expire"`
	RefreshToken	string	`bson:"refresh_token" json:"refresh_token"`
	RefreshExpire	string	`bson:"refresh_expire" json:"refresh_expire"`
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
