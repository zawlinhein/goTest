package helper

import(
	"os"
	"asdf/structType"
	"encoding/json"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertFromFile(filename string,collection *mongo.Collection) error{
	file ,err:=os.Open(filename)
	if err!=nil{
		return err
	}
	defer file.Close()	

	var items []structType.Item
	decoder:=json.NewDecoder(file)
	if err:=decoder.Decode(&items); err!=nil{
		return err
	}

	var itemToAdd []structType.Item
	for i := 0; i < len(items); i++{	
		items[i].Id=i
		items[i].LastAccess=time.Now()
		itemToAdd=append(itemToAdd,items[i])
		
		if len(itemToAdd)==50{
			if err:=InsertItems(itemToAdd,collection);err!=nil{
				return err
			}
			itemToAdd=itemToAdd[:0]
		}
	}

	if len(itemToAdd)>0{
		if err:=InsertItems(itemToAdd,collection);err!=nil{
			return err
		}
		itemToAdd=nil
	}

	return nil
}

func InsertItems(items []structType.Item,collection *mongo.Collection) error{
	dataForThis:=make([]interface{},len(items))

	for i,item:=range items{
		dataForThis[i]=item
	}

	result,err:=collection.InsertMany(context.TODO(),dataForThis)
	if err!=nil{
		return err
	}
	fmt.Printf("Number of documents inserted: %d\n", len(result.InsertedIDs))
	return nil
}