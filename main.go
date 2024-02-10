package main

import (
	"context"
	"fmt"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	cfg := mongodb.Config{URI: "mongodb://root:rootpassword@localhost:27017/"}
	ctx := context.Background()
	db := mongodb.New(ctx, cfg)
	result, err := db.GetClient().Database("test").Collection("movie").InsertOne(
		context.TODO(),
		bson.D{
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
		})
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%s\n", result)
}
