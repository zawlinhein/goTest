package main

import (
    "context"
    "fmt"
    "log"
    "asdf/helper"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


func main() {

    clientOptions := options.Client().ApplyURI()

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    collection := client.Database("dummy").Collection("locData")

/*     insertData := bson.M{"id":6,"name":"asdf","description":"iii","price":77.77,"stock":9}
    inserted,err:= collection.InsertOne(ctx,insertData) 
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie with id",inserted.InsertedID)  */

    if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "bulk-insert":
            if err:=helper.InsertFromFile(os.Args[2],collection); err!=nil{
                log.Fatal(err)
            }
        case "delete-all":
            if err:=helper.DeleteAll(collection); err!=nil{
                log.Fatal(err)
            }
        case "doc-count":
            if err:=helper.DocCount(collection); err!=nil{
                log.Fatal(err)
            }
        case "duplicate-check":
            if err:=helper.DuplicateCheck(collection); err!=nil{
                log.Fatal(err)
            }
        case "check-status":
            if err:=helper.CheckStatus(collection); err!=nil{
                log.Fatal(err)
            }
        case "update-time":
            if err:=helper.UpdateTime(os.Args[2],collection); err!=nil{
                log.Fatal(err)
            }
		default:
			showHelper()
		}
	} else {
		showHelper()
	}

}


func showHelper() {
	fmt.Println("Usage ./geocoder command [OPTION] inputfile")
}