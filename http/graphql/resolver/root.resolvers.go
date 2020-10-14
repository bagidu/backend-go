package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bagiduid/backend/http/graphql/generated"
	"github.com/bagiduid/backend/models"
	"github.com/bagiduid/backend/services/mail"
	"github.com/bagiduid/backend/services/user"
)

func (r *mutationResolver) Register(ctx context.Context, input models.RegisterUser) (*models.User, error) {
	u := user.User{
		Name:     input.Name,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}
	if err := r.UserService.Create(&u); err != nil {
		return nil, err
	}

	r.MailService.Send(&mail.Mail{
		To:      u.Email,
		Subject: "Email Verification",
		Text:    fmt.Sprintf("Hello %s please verify your mail by click this link: %s", u.Name, "https://bagidu.id/verify/codeddddddddddd"),
	})

	return &models.User{
		ID:       u.ID.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Username: u.Username,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context, limit int, offset int) ([]*models.User, error) {
	res, err := r.UserService.All(limit, offset)
	if err != nil {
		return nil, err
	}
	var users []*models.User

	for _, u := range res {
		users = append(users, &models.User{
			ID:       u.ID.Hex(),
			Email:    u.Email,
			Name:     u.Name,
			Username: u.Username,
		})
	}
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
