package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	secret "mn128.com/src/utils"
)

type singleton struct {
	Instance *mongo.Collection
}

var instance *singleton
var once sync.Once

var uri = secret.GetSecret()

func GetDbInstance() *singleton {
	var serverAPI = options.ServerAPI(options.ServerAPIVersion1)
	var opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	coll := client.Database("db").Collection("tea")

	once.Do(func() {
		instance = &singleton{Instance: coll}
	})
	return instance
}
