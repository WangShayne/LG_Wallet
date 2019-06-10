package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Trs_params struct {
	//TransactionId string `json:"transactionId"`
	Address  string `json:"address"`
	Contract string `json:"contract"`
	Balance  int64  `json:"balance"`
	PageNo   int64  `json:"pageNo"`
}

type Token struct {
	Id     string               `json:"id" bson:"_id"`
	Tokens []primitive.ObjectID `json:"tokens" bson:"tokens"`
}

type LastBlock struct {
	Id                      primitive.ObjectID `json:"_id" bson:"_id"`
	LastBlock               int32              `bson:"lastBlock"`
	LastBackwardBlock       int32              `bson:"lastBackwardBlock"`
	LastPusherBlock         int32              `bson:"lastPusherBlock"`
	LastTokensBlock         int32              `bson:"lastTokensBlock"`
	LastTokensBackwardBlock int32              `bson:"lastTokensBackwardBlock"`
}

// 交易记录
type TransAction struct {
	ID             string               `json:"_id" bson:"_id"`
	Addresses      []string             `json:"addresses" bson:"addresses"`
	Operations     []primitive.ObjectID `json:"operations" bson:"operations"`
	Contract       interface{}          `json:"contract" bson:"contract"`
	BlockNumber    int32                `json:"blockNumber" bson:"blockNumber"`
	TimeStamp      string               `json:"timeStamp" bson:"timeStamp"`
	Nonce          int32                `json:"nonce" bson:"nonce"`
	From           string               `json:"from" bson:"from"`
	To             string               `json:"to" bson:"to"`
	Value          string               `json:"value" bson:"value"`
	Gas            string               `json:"gas" bson:"gas"`
	GasPrice       string               `json:"gasPrice" bson:"gasPrice"`
	GasUsed        string               `json:"gasUsed" bson:"gasUsed"`
	Input          string               `json:"input" bson:"input"`
	Error          string               `json:"error" bson:"error"`
	OperationsList []Operation          `json:"operationsList" bson:"operation"`
}

// 交易记录
type Operation struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	TransactionID string             `json:"transactionId" bson:"transactionId"`
	Contract      primitive.ObjectID `json:"contract" bson:"contract"`
	From          string             `json:"from" bson:"from"`
	To            string             `json:"to" bson:"to"`
	Type          string             `json:"type" bson:"type"`
	Value         string             `json:"value" bson:"value"`
	// ContractInfo  *Contract          `json:"contractInfo"`
}

// 返回的交易记录
type TransActions struct {
	TransAction
	Operation []Operation `json:"operation"`
	Balance   int32       `json:"balance"`
}

type Contract struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Adderss     string             `json:"address" bson:"address"`
	Decimals    int32              `json:"decimals" bson:"decimals"`
	Enabled     bool               `json:"enabled" bson:"enabled"`
	Name        string             `json:"name" bson:"name"`
	Symbol      string             `json:"symbol" bson:"symbol"`
	TotalSupply string             `json:"totalSupply" bson:"totalSupply"`
	Verified    bool               `json:"verified" bson:"verified"`
}
