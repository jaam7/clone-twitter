package bd

import (
	"context"
	"time"

	"github.com/jalamar/clone-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckAvailabilityEmail(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("tuiter")
	collection := db.Collection("user")

	condition := bson.M{"email": email}

	var user models.Usuario
	err := collection.FindOne(ctx, condition).Decode(&user)
	ID := user.ID.Hex()

	if err != nil {
		return user, false, ID
	}
	
	return user, true, ID
}
