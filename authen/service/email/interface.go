package email

type Service interface {
	SendVerifyUrl(email string) error
}
