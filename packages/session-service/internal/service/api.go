package service


type api struct {
	session *gocql.Session
}

type MessageDatastore interface {
	CreateMessage(i CreateMessageInput) error
	SendMessages(magicNumber int, c *config.SmtpConfig) error
	GetMessagesByEmail(i GetMessagesByEmailInput) ([]*Message, string)
}

func NewAPI(session *gocql.Session) *api {
	return &api{session: session}
}