package helper

import(
	"context"
	"asdf/structType"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckStatus(collection *mongo.Collection) error{
	//add all data to results
	cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return err
    }
	
	var results []structType.Item
    if err := cursor.All(context.TODO(), &results); err != nil {
        return err
    }
	///

	t2 := time.Now()
	for _,point:=range(results){
		t1 := point.LastAccess
		duration := t2.Sub(t1)
		totalMinutes := int(duration.Minutes())
		if totalMinutes <= 5{
			update := bson.M{
				"$set": bson.M{
					"status": "up",
				},
			}
			_,err:=collection.UpdateOne(context.TODO(), bson.M{"id":point.Id}, update)
			if err != nil {
				return err
			}
		} else {
			update := bson.M{
				"$set": bson.M{
					"status": "down",
				},
			}
			_,err:=collection.UpdateOne(context.TODO(), bson.M{"id":point.Id}, update)
			if err != nil {
				return err
			}
		}
		
	}
	fmt.Println("success")
	return nil
} 