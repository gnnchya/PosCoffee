package domain

type Consumer struct {
	ID 				string 	`bson:"_id" json:"id"`
	ClientID		string 	`bson:"client_id" json:"client_id"`
	ClientSecret	string 	`bson:"client_secret" json:"client_secret"`
}

type Tokens struct {
	ID 				string 	`bson:"_id" json:"id"`
	UID 			string	`bson:"uid" json:"uid"`
	Token 			string	`bson:"token" json:"token"`
	Expire 			string	`bson:"expire" json:"expire"`
	RefreshToken	string	`bson:"refresh_token" json:"refresh_token"`
	RefreshExpire	string	`bson:"refresh_expire" json:"refresh_expire"`
}

