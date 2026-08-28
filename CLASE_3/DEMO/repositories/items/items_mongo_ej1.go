package items

import (
	"context"
	"main/models/items"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ==========================================
// EJERCICIO 1: Medir el Problema
// ==========================================

type ItemsMongoDB struct {
	Client *mongo.Client
}

func (repo ItemsMongoDB) GetItemByID(id string) (items.ItemModel, error) {

	// Simulamos que ir a la BD en disco es muy lento
	time.Sleep(5000 * time.Millisecond)

	var item items.ItemModel
	err := repo.Client.Database("items-api").Collection("items").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item)
	return item, err
}

func (repo ItemsMongoDB) CreateItem(item items.ItemModel) error {
	_, err := repo.Client.Database("items-api").Collection("items").InsertOne(context.TODO(), item)
	return err
}
