package tokenin

type RefreshInput struct {
	ClientID     string `json:"client_id" form:"client_id" binding:"required"`
	UID       	 string `json:"uid" form:"uid" binding:"required"`
	ClientSecret string `json:"client_secret" form:"client_secret" binding:"required"`
	GrantType    string `json:"grant_type" form:"grant_type" binding:"required"`
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
	Scope        string `json:"scope" form:"scope"`
}
