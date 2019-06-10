package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var client *mongo.Client


func ConnecToDB() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil{
		fmt.Println(err)
	}

	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	err =client.Connect(ctx)
	fmt.Println("mongodb 连接成功")
}


func GetCtx(count time.Duration) context.Context{
	ctx,_ := context.WithTimeout(context.Background(),count * time.Second)
	return ctx
}

func Col(name string) *mongo.Collection{
	dbname := viper.GetString("database.dbname")
	return client.Database(dbname).Collection(name)
}
