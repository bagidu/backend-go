package resolver

import (
	"context"
	"errors"

	"github.com/bagiduid/backend/services/user"
	"github.com/go-chi/jwtauth"
)

// GetUser helper
func GetUser(ctx context.Context, service user.Service) (*user.User, error) {
	_, c, e := jwtauth.FromContext(ctx)
	if e != nil {
		return nil, errors.New("Unable to parse token")
	}

	id, _ := c["uid"].(string)
	if id == "" {
		return nil, errors.New("User id not found in token")
	}

	return service.FindOne(id)
}
