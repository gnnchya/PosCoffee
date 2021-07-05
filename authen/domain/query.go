package domain

type User struct{
	Username		string
	Password 		string
	UID  			string
	MetaData		metadata
	RoleID			[]string
	CreateAt 		int64
	UpdateAt 		int64
	DeleteAt 		int64
}

type metadata struct{
	Prefix 			Language
	Name 			Language
	Lastname 		Language
	Email 			string
	MobileNumber	string
	LineID			string
	Facebook 		string
	BankAccount		string
	Gender 			string
	BirthDay		int64
}

type Language struct{
	TH			string
	EN 			string
}

