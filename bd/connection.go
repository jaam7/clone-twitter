package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection, err = connectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://alejandro:j4v13r1997@tuiter.mccatz9.mongodb.net/test")

func connectBD() (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	log.Println("succesfull connection")
	return client, err
}

func CheckConnection() bool {
	err = MongoConnection.Ping(context.TODO(), nil)
	return err == nil
}
