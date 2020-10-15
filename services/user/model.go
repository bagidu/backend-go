package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Username   string             `json:"username,omitempty"`
	Email      string             `json:"email,omitempty"`
	Password   string             `json:"password,omitempty"`
	VerifiedAt *time.Time         `bson:"verified_at" json:"verified_at,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at,omitempty"`
}

// HashPassword hash password before save
func (u *User) HashPassword() error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	u.Password = string(pwd)
	return err
}

// CheckPassword func
func (u *User) CheckPassword(pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
}
