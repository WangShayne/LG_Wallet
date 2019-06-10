package common

import (
	"LG_wallet/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateResponseBody(ctx *gin.Context,req *model.ResponseDefaultBody,resBody interface{},errCode int) {

	var res model.CreateResp

	if errCode == 0 {
		res = &model.ResponseOKBody{
			req,
			resBody,
		}
	} else {
		errInfo := CreateErrorInfo(errCode)
		fmt.Println(errInfo)
		res = &model.ResponseErrorBody{
			req,
			errInfo,
		}
	}
	ctx.JSON(200,res)
}