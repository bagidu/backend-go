package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bagiduid/backend/http/graphql/generated"
	"github.com/bagiduid/backend/models"
	"github.com/bagiduid/backend/services/mail"
	"github.com/bagiduid/backend/services/user"
	jwt "github.com/dgrijalva/jwt-go"
)

func (r *mutationResolver) Register(ctx context.Context, input models.RegisterUser) (*models.User, error) {
	u := user.User{
		Name:      input.Name,
		Email:     input.Email,
		Username:  input.Username,
		Password:  input.Password,
		CreatedAt: time.Now(),
	}
	log.Printf("User before save: %v", u)
	if err := r.UserService.Create(&u); err != nil {
		return nil, err
	}
	log.Printf("User after save: %v", u)

	_, code, err := r.JWT.Encode(jwt.MapClaims{"type": "email_verification", "uid": u.ID.Hex()})
	if err != nil {
		return nil, err
	}
	r.MailService.Send(&mail.Mail{
		To:      u.Email,
		Subject: "Email Verification",
		Text:    fmt.Sprintf("Hello %s please verify your mail by click this link: https://bagidu.id/verify/%s", u.Name, code),
	})

	return &models.User{
		ID:         u.ID.Hex(),
		Name:       u.Name,
		Email:      u.Email,
		Username:   u.Username,
		VerifiedAt: u.VerifiedAt,
		CreatedAt:  u.CreatedAt,
	}, nil
}

func (r *mutationResolver) VerifyEmail(ctx context.Context, code string) (*models.User, error) {
	t, err := r.JWT.Decode(code)
	if err != nil {
		return nil, err
	}

	claims, _ := t.Claims.(jwt.MapClaims)

	uid := fmt.Sprintf("%s", claims["uid"])
	u, err := r.UserService.FindOne(uid)
	if err != nil {
		return nil, err
	}

	if u.VerifiedAt != nil {
		return nil, errors.New("Email already verified")
	}

	now := time.Now()
	u.VerifiedAt = &now

	if err := r.UserService.Update(u); err != nil {
		return nil, errors.New("Unable to verify user (update error)")
	}

	return &models.User{
		ID:         u.ID.Hex(),
		Email:      u.Email,
		Name:       u.Name,
		Username:   u.Username,
		VerifiedAt: u.VerifiedAt,
		CreatedAt:  u.CreatedAt,
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
			ID:         u.ID.Hex(),
			Email:      u.Email,
			Name:       u.Name,
			Username:   u.Username,
			VerifiedAt: u.VerifiedAt,
			CreatedAt:  u.CreatedAt,
		})
	}
	return users, nil
}

func (r *queryResolver) User(ctx context.Context, filter models.UserFilter) (*models.User, error) {
	if filter.ID != nil {
		u, e := r.UserService.FindOne(*filter.ID)
		if e != nil {
			return nil, e
		}

		return &models.User{
			ID:         u.ID.Hex(),
			Name:       u.Name,
			Email:      u.Email,
			Username:   u.Username,
			VerifiedAt: u.VerifiedAt,
			CreatedAt:  u.CreatedAt,
		}, nil
	} else if filter.Username != nil {
		return nil, errors.New("Currently unable find user by username")
	}

	return nil, errors.New("Invalid filter parameter")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
