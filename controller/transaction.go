package controller

import (
	"LG_wallet/common"
	"LG_wallet/db"
	"LG_wallet/model"
	"fmt"

	"github.com/gin-gonic/gin"
	. "github.com/mitchellh/mapstructure"
)

func GetTransaction(params interface{}, ctx *gin.Context, req *model.ResponseDefaultBody) {

	var p model.Trs_params

	err := Decode(params, &p)
	if err != nil {
		fmt.Println(err)
	}

	if p.Address == "" {
		common.CreateResponseBody(ctx, req, nil, 10004)
		return
	}

	t, err := db.QueryTransactions(p)

	if err != nil {
		common.CreateResponseBody(ctx, req, nil, 10005)
		return
	}

	if t == nil {
		common.CreateResponseBody(ctx, req, nil, 10005)
		return
	}

	common.CreateResponseBody(ctx, req, t, 0)
	return
}
