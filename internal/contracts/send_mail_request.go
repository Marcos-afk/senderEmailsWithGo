package contracts

type SendMailRequest struct {
	To      []string
	Subject string
	Message string
}