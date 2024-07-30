package helper

import(
	"context"
	"asdf/structType"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	/* "go.mongodb.org/mongo-driver/mongo/options" */
)

func DuplicateCheck(collection *mongo.Collection) error{
	checkArr:= make(map[int]int)
	/* findOptions := options.Find().SetLimit(20) */
	cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return err
    }
	
	var results []structType.Item
    if err := cursor.All(context.TODO(), &results); err != nil {
        return err
    }

	flag:=false
	for _,point:=range(results){
		if _,exist:=checkArr[point.Id]; exist{
			if _, err = collection.DeleteOne(context.TODO(), bson.M{"id": point.Id}); err!=nil{
				return err
			}
			fmt.Printf("duplicate of id %d found and deleted\n",point.Id)
			checkArr[point.Id]++
			flag=true
		} else {checkArr[point.Id]=0}		
	}

	
	if !flag{
		fmt.Println("There's no duplicate data!")
	}else{
		for id:=range checkArr{
			if checkArr[id]==1{
				fmt.Printf("%d item of id-%d is deleted\n",checkArr[id],id)	
			}else if checkArr[id]>1{
				fmt.Printf("%d items of id-%d is deleted\n",checkArr[id],id)	
			}
		}
	}
	return nil
}