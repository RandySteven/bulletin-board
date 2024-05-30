package email

type IEmail interface {
	SendEmailRegister() error
	SendEmailTest() error
}
