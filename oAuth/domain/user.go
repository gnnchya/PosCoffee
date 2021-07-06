package domain

type ConsumerStruct struct {
	ID 				string 	`bson:"_id" json:"id"`
	ClientID		string 	`bson:"client_id" json:"client_id"`
	ClientSecret	string 	`bson:"client_secret" json:"client_secret"`
	Scope           string 	`bson:"scope" json:"scope"`
	CreatedAt       int64  	`bson:"createdAt" json:"createdAt"`
}

type TokenStruct struct {
	ID 				string 	`bson:"_id" json:"id"`
	UID 			string	`bson:"uid" json:"uid"`
	AccessToken 	string	`bson:"access_token" json:"access_token"`
	Expire 			int64	`bson:"expire" json:"expire"`
	RefreshToken	string	`bson:"refresh_token" json:"refresh_token"`
	RefreshExpire	int64	`bson:"refresh_expire" json:"refresh_expire"`
}

