package items

import (
	"context"
	"main/models/items"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemsMongoDB struct {
	Client *mongo.Client
}

func (repo ItemsMongoDB) GetItemByID(id string) (items.ItemModel, error) {
	var item items.ItemModel
	// Usamos el cliente para buscar en la colección "items"
	err := repo.Client.Database("items-api").Collection("items").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item)
	return item, err
}

func (repo ItemsMongoDB) CreateItem(item items.ItemModel) error {
	_, err := repo.Client.Database("items-api").Collection("items").InsertOne(context.TODO(), item)
	return err
}
