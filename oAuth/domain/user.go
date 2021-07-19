package domain

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

