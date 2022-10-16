package bd

import (
	"context"
	"time"

	"github.com/jalamar/clone-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuiter")
	collection := db.Collection("user")

	user.Password, _ = EncrytpPassword(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
