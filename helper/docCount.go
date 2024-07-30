package helper

import(
	"context"
	"fmt"
	"asdf/structType"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DocCount(collection *mongo.Collection) error{
	cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return err
    }
	
	var results []structType.Item
    if err := cursor.All(context.TODO(), &results); err != nil {
        return err
    }

	fmt.Printf("Found %d documents\n", len(results))
 	 for _, item := range results {
        fmt.Printf("Found document: %+v\n", item)
    } 
	return nil
}