package authenticationin

type LoginInput struct {
	Username string `json:"username"  validate:"required"`
	Password string `json:"implement"  validate:"required"`
} // @Name LoginInput

type VerifyInput struct {
	Code    string `json:"code" validate:"required"`
	RefCode string `json:"ref_code" validate:"required"`
	Contact string `json:"contact" validate:"required"`
} // @Name VerifyInput
