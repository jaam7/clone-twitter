package bd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jalamar/clone-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuiter")
	collection := db.Collection("user")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, condicion).Decode(&profile)
	if err != nil {
		fmt.Println("Register not found" + err.Error())
		return profile, errors.New("Register not found" + err.Error())
	}

	profile.Password = ""

	return profile, nil
}
