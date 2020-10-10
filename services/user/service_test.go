package user

import (
	"context"
	"testing"

	"github.com/benweissmann/memongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDB(t *testing.T) (*mongo.Database, *memongo.Server) {
	memDB, err := memongo.Start("4.0.5")
	if err != nil {
		t.Error(err)
	}

	dbName := memongo.RandomDatabase()
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(memDB.URI()))
	if err != nil {
		t.Error(err)
	}

	return client.Database(dbName), memDB
}

func TestFindOne(t *testing.T) {
	db, mem := setupDB(t)
	defer mem.Stop()

	// Fake data
	u := User{
		ID:    primitive.NewObjectID(),
		Name:  "Sucipto",
		Email: "email@local.dev",
	}
	create, err := db.Collection("users").InsertOne(context.Background(), u)
	if err != nil {
		t.Error(err)
	}

	t.Logf("User id: %s, inserted id: %v", u.ID.Hex(), create.InsertedID)

	var s Service = NewUserService(db)
	res, err := s.FindOne(u.ID.Hex())
	if err != nil {
		t.Error(err)
	}

	if res.Name != u.Name {
		t.Error("Name should same")
	}

}
