package db

import (
	"LG_wallet/model"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func QueryTransactions(p model.Trs_params) (t []model.TransAction, Err error) {
	var colName = viper.GetString("database.Col_transactions")
	var colOptions = viper.GetString("database.Col_transactionoperations")
	c := Col(colName)
	ctx := GetCtx(30)

	var limit int64 = 10

	pipes := mongo.Pipeline{
		{{"$match",bson.D{{"addresses",p.Address}}}},
		{{"$skip",(p.PageNo - 1) * limit}},
		{{"$limit",limit}},
		{{"$lookup",bson.D{{"from",colOptions},{"localField","operations"},{"foreignField","_id"},{"as","operation"}}}},
		//{{"$sort",bson.D{{"timeStamp",-1}}}},
	}
	pipes = append(pipes,bson.D{{"$sort",bson.D{{"timeStamp",-1}}}})
	fmt.Println(pipes)

	option := options.Aggregate()

	cur, err := c.Aggregate(ctx, pipes, option)
	fmt.Println(cur)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var a model.TransAction
		_ = cur.Decode(&a)
		t = append(t, a)
	}
	return t, nil
}


func QueryOperation(optionsId primitive.ObjectID) (o model.Operation) {

	colName := viper.GetString("database.Col_transactionoperations")
	c := Col(colName)
	ctx := GetCtx(5)

	filter := bson.D{
		{"_id", optionsId},
	}

	_ = c.FindOne(ctx, filter).Decode(&o)

	return o
}

func QueryContract(contractId primitive.ObjectID) (cont model.Contract) {

	colName := viper.GetString("database.Col_erc20contracts")
	c := Col(colName)
	ctx := GetCtx(5)

	filter := bson.D{
		{"_id", contractId},
	}

	_ = c.FindOne(ctx, filter).Decode(&cont)

	return cont
}
