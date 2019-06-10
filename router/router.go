package router

import (
	ctrl "LG_wallet/controller"
	"LG_wallet/model"
	"github.com/gin-gonic/gin"
)



func Routers(ctx *gin.Context){
	var res model.RequestBody
	var req *model.ResponseDefaultBody
	if err := ctx.Bind(&res);err!=nil{
		ctx.JSON(200,gin.H{"err":err.Error()})
		return
	}
	req = &model.ResponseDefaultBody{
		res.ID,
		res.Jsonrpc,
	}
	switch res.Method {

	case "LG_transaction":
			ctrl.GetTransaction(res.Params,ctx,req)
		}

}