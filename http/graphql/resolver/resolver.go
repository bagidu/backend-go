package resolver

import (
	"github.com/bagiduid/backend/services/mail"
	"github.com/bagiduid/backend/services/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	UserService user.Service
	MailService mail.Service
}
