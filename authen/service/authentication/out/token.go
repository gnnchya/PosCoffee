package out

type RespBody struct {
	Data Token `json:"data"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Expired      int64  `json:"expired,omitempty"`
}

type RepsValidate struct {
	Data ValidateToken `json:"data"`
}

type ValidateToken struct {
	UserId  string `json:"user_id"`
	Expired int64  `json:"expired"`
}
