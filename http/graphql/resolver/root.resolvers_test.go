package resolver

import (
	"errors"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bagiduid/backend/http/graphql/generated"
	"github.com/bagiduid/backend/services/user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setup(t *testing.T, r *Resolver) *client.Client {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r})))
	return c
}

// func TestRegisterMutation(t *testing.T) {
// 	userService := &user.MockService{}
// 	r := Resolver{UserService: userService}
// 	c := setup(t, &r)

// 	t.Run("Success register", func(t *testing.T) {
// 		q := `
// 		mutation {
// 			register(input:{
// 				name:"Sucipto"
// 			})
// 		}
// 		`
// 	})
// }

func TestGetUsers(t *testing.T) {
	userService := &user.MockService{}
	r := Resolver{UserService: userService}
	c := setup(t, &r)

	t.Run("Test with 1 result", func(t *testing.T) {

		q := `
		query {
			users(limit: 1, offset: 2) {
				id
				name
			}
		}
		`
		var resp struct {
			Users []struct {
				ID   string
				Name string
			}
		}
		fake := []*user.User{
			{
				ID:   primitive.NewObjectID(),
				Name: "Sucipto",
			},
		}
		userService.On("All", 1, 2).Return(fake, nil).Once()
		c.MustPost(q, &resp)

		userService.AssertExpectations(t)

		assert.Equal(t, fake[0].ID.Hex(), resp.Users[0].ID, "Result id shold same")
		assert.Equal(t, fake[0].Name, resp.Users[0].Name, "Result name shold same")

	})

	t.Run("Test error handling", func(t *testing.T) {
		q := `
		query {
			users(limit: 1, offset: 3) {
				id
				name
			}
		}
		`
		var resp struct {
			Users []struct {
				ID   string
				Name string
			}
		}
		fake := []*user.User{}
		userService.On("All", 1, 3).Return(fake, errors.New("not found")).Once()
		assert.Error(t, c.Post(q, &resp))

		userService.AssertExpectations(t)
	})
}
