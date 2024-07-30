package helper

import(
	"time"
	"strconv"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTime(id_str string,collection *mongo.Collection) error{
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err
	}

	updateData:=bson.M{
		"$set":bson.M{
			"lastaccess":time.Now(),
		},
	}
	updateStatus,updateErr:=collection.UpdateOne(context.TODO(),bson.M{"id":id},updateData)
	if updateErr != nil{
		return updateErr
	}
	fmt.Println(updateStatus)
	return nil
}