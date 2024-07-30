package helper

import(
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteAll(collection *mongo.Collection) error{
	result,err:=collection.DeleteMany(context.TODO(),bson.M{})
	if err!=nil{
		return err
	}
	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)
	return nil
}