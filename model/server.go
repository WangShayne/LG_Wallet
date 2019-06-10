package model

type RequestBody struct {
	ID int `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params interface{} `json:"params"`
}

//	默认返回格式
type ResponseDefaultBody struct {
	ID int `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

//	成功的返回格式
type ResponseOKBody struct {
	*ResponseDefaultBody
	Result interface{} `json:"result"`
}

//	错误的返回格式
type ResponseErrorBody struct{
	*ResponseDefaultBody
	Error *ResponseErrorInfo `json:"error"`
}

//	错误消息
type ResponseErrorInfo struct {
	Code int `json:"code"`
	Message string `json:"message"`
}


type CreateResp interface {
	CreateResponse()
}

func (o *ResponseOKBody) CreateResponse(){

}


func (e *ResponseErrorBody) CreateResponse(){

}

