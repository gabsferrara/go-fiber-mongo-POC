package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Find(collection string, documents any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbname).Collection(collection)

	cursor, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	return cursor.All(context.Background(), documents)
}

func FindByID(collection, id string, document any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbname).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}

	return c.FindOne(context.Background(), filter).Decode(document)

}
