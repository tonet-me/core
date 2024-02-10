package main

import (
	"context"
	"fmt"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
)

func main() {
	cfg := mongodb.Config{URI: "mongodb://root:rootpassword@localhost:27017/"}
	db := mongodb.New(cfg)
	userDB := usermongo.New(usermongo.Config{
		DBName:   "test",
		CollName: "user",
	}, db)

	//isEist, user, err := userDB.IsUserExistByEmail(context.TODO(), "kswsssss@gmail.com")
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	s, err := userDB.DeActiveUser(context.TODO(), "65c728c64bb1081b4046d682")
	fmt.Println(s, err)
	//user, err := userDB.CreateNewUser(context.TODO(), entity.User{
	//	LastName:    "k",
	//	Email:       "kswssss@gmail.com",
	//	PhoneNumber: "091221702858",
	//	Status:      1,
	//})
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Printf("%s\n", user)
	//result, err := db.GetClient().Database("test").Collection("movie").InsertOne(
	//	context.TODO(),
	//	bson.D{
	//		{"item", "canvas"},
	//		{"qty", 100},
	//		{"tags", bson.A{"cotton"}},
	//		{"size", bson.D{
	//			{"h", 28},
	//			{"w", 35.5},
	//			{"uom", "cm"},
	//		}},
	//	})
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Printf("%s\n", result)
}
