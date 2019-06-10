package common

import (
	"LG_wallet/model"
)

const (
	// normal

	// 系统级错误
	SERVER_ERR = 10001
	DB_ERR     = 10002
	METHOD_ERR = 10003
	PARAMS_ERR = 10004

	// - 接口级错误 -

	// 汇率查询错误 rate
	RATE_ERR = 20001

	// 代币查询错误 erc20
	ERC20_ERR = 20002

	// 代币价格查询错误 price
	ERC20_PRICE_ERR = 30001

	// 域名注册错误 ens register
	ENS_REGISTER_ERR       = 40001
	ENS_REGISTER_EXIST_ERR = 40002
	// 域名查询错误
	ENS_SEARCH_ERR      = 50001
	ENS_SEARCHBYKEY_ERR = 50002
	// 文章查询错误
	ARTICLE_SEARCH_ERR = 60001
	// 交易记录查询错误
	TRANSACTION_SEARCH_ERR = 70001
)

func CreateErrorInfo(code int) (errInfo *model.ResponseErrorInfo){

	switch code {

	case SERVER_ERR:
		errInfo = &model.ResponseErrorInfo{
			10001,
			"系统错误,请稍后再试",
		}
	case DB_ERR:
		errInfo = &model.ResponseErrorInfo{
			10002,
			"查询失败,请稍后再试",
		}
	case METHOD_ERR:
		errInfo = &model.ResponseErrorInfo{
			10003,
			"请检查接口地址",
		}
	case PARAMS_ERR:
		errInfo = &model.ResponseErrorInfo{
			10004,
			"请检查提交参数",
		}
	case ENS_REGISTER_ERR:
		errInfo = &model.ResponseErrorInfo{
			40001,
			"注册失败,请稍后再试",
		}
	case ENS_REGISTER_EXIST_ERR:
		errInfo = &model.ResponseErrorInfo{
			40002,
			"注册失败,域名已存在",
		}
	case ENS_SEARCH_ERR:
		errInfo = &model.ResponseErrorInfo{
			50001,
			"未查询到相应公钥",
		}
	case ENS_SEARCHBYKEY_ERR:
		errInfo = &model.ResponseErrorInfo{
			50002,
			"未查询到相应域名",
		}
	case ARTICLE_SEARCH_ERR:
		errInfo = &model.ResponseErrorInfo{
			60001,
			"未查询到相应文档",
		}
	case TRANSACTION_SEARCH_ERR:
		errInfo = &model.ResponseErrorInfo{
			70001,
			"当前没有交易记录",
		}
	default:
		errInfo = &model.ResponseErrorInfo{
			10001,
			"系统错误,请稍后再试",
		}
	}

	return  errInfo
}