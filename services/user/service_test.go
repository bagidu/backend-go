package user

import (
	"context"
	"testing"

	"github.com/benweissmann/memongo"
	"go.mongodb.org/mongo-driver/bson"
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

	if u.ID.Hex() != create.InsertedID.(primitive.ObjectID).Hex() {
		t.Errorf("Inserted id not same, user id: %s vs inserted id: %v", u.ID.Hex(), create.InsertedID)
	}

	// Create Service
	var s Service = NewUserService(db)
	t.Run("Should found existing id", func(t *testing.T) {

		res, err := s.FindOne(u.ID.Hex())
		if err != nil {
			t.Error(err)
		}

		if res.Name != u.Name {
			t.Error("Name should same")
		}
	})

	t.Run("Invalid id should not found", func(t *testing.T) {
		_, err := s.FindOne(primitive.NewObjectID().Hex())
		if err == nil {
			t.Error("Should error / not found")
		}
	})

	t.Run("Invalid object id hex", func(t *testing.T) {
		_, err := s.FindOne("xxxxx")
		if err == nil {
			t.Error("Should error / not found")
		}
	})

}

func TestAll(t *testing.T) {

	db, mem := setupDB(t)
	defer mem.Stop()

	// Fake data
	users := []interface{}{
		User{
			ID:    primitive.NewObjectID(),
			Name:  "Sucipto",
			Email: "email@local.dev",
		},
		User{
			ID:    primitive.NewObjectID(),
			Name:  "Lubna",
			Email: "lubna@local.dev",
		},
	}

	created, err := db.Collection("users").InsertMany(context.Background(), users)
	if err != nil {
		t.Error(err)
	}

	if len(created.InsertedIDs) != 2 {
		t.Errorf("Inserted id expected 2, got %d", len(created.InsertedIDs))
	}

	// Create Service
	var s Service = NewUserService(db)

	t.Run("Return all result", func(t *testing.T) {
		u, e := s.All(50, 0)
		if e != nil {
			t.Errorf("Should not error: %s", e.Error())
		}

		if len(u) != 2 {
			t.Errorf("Expected result 2, got %d", len(u))
		}
	})

	t.Run("Should return 1", func(t *testing.T) {
		u, e := s.All(1, 0)
		if e != nil {
			t.Errorf("Should not error: %s", e.Error())
		}

		if len(u) != 1 {
			t.Errorf("Expected result 1, got %d", len(u))
		}
	})

	t.Run("Should return 0", func(t *testing.T) {
		u, _ := s.All(1, 3)
		// if e == nil {
		// 	t.Errorf("Should  error: not found")
		// }

		if len(u) != 0 {
			t.Errorf("Expected result 0, got %d", len(u))
		}
	})

}

func TestCreate(t *testing.T) {

	db, mem := setupDB(t)
	defer mem.Stop()

	// Create Service
	var s Service = NewUserService(db)

	t.Run("Success on create user", func(t *testing.T) {
		u := User{
			Name:     "Sucipto",
			Email:    "email@local.dev",
			Username: "sucipto",
		}
		if err := s.Create(&u); err != nil {
			t.Errorf("Should not error: %s", err.Error())
		}

		if u.ID.IsZero() {
			t.Error("User id should generated")
		}

		var uv User
		if err := db.Collection("users").FindOne(context.TODO(), bson.M{"_id": u.ID}).Decode(&uv); err != nil {
			t.Error(err)
		}

		if uv.Name != u.Name {
			t.Errorf("Expected result Name is equal, got: %s", uv.Name)
		}
	})

	t.Run("Should not able create user with same email", func(t *testing.T) {
		u := User{
			Name:     "Lubna",
			Email:    "lubna@local.dev",
			Username: "lubna",
		}

		// First creation
		if err := s.Create(&u); err != nil {
			t.Errorf("Should not error: %s", err.Error())
		}

		// Create again with same payload
		if err := s.Create(&u); err == nil {
			t.Errorf("Should error on create with same email")
		}
	})
}
