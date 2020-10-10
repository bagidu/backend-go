package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service user
type Service interface {
	FindOne(id string) (*User, error)
	All(limit int, offset int) ([]*User, error)
	Create(*User) error
}

// NewUserService ...
func NewUserService(db *mongo.Database) Service {
	return &mongoService{db}
}

type mongoService struct {
	db *mongo.Database
}

func (s *mongoService) collection() *mongo.Collection {
	return s.db.Collection("users")
}

// FindOne ...
func (s *mongoService) FindOne(id string) (*User, error) {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user User
	if err := s.collection().FindOne(context.TODO(), bson.M{"_id": uid}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// All ...
func (s *mongoService) All(limit int, offset int) ([]*User, error) {
	res, err := s.collection().Aggregate(context.TODO(), []bson.M{
		{
			"$limit": limit,
		},
		{
			"$skip": offset,
		},
	})

	if err != nil {
		return nil, err
	}

	var users []*User
	for res.Next(context.TODO()) {
		var user User

		if err := res.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// Create ...
func (s *mongoService) Create(u *User) error {
	u.ID = primitive.NewObjectID()
	_, err := s.collection().InsertOne(context.TODO(), u)
	return err
}
