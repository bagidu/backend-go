package mail

import (
	"context"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type mailgunService struct {
	mg *mailgun.MailgunImpl
}

// NewMailgunService ...
func NewMailgunService() Service {
	key := os.Getenv("MAILGUN_KEY")
	if key == "" {
		panic("Mailgun key not provided, set it in env var: MAILGUN_KEY")
	}
	mg := mailgun.NewMailgun("mail.bagidu.id", key)

	return &mailgunService{mg}
}

func (s *mailgunService) Send(m *Mail) error {
	msg := s.mg.NewMessage("no-reply@bagidu.id", m.Subject, m.Text, m.To)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := s.mg.Send(ctx, msg)
	return err
}
