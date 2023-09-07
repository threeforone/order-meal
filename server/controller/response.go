package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// data
type Data struct {
	Data  interface{} `json:"data"`  // 返回数据
	Page  int         `json:"page"`  // 当前页
	Total int64       `json:"total"` // 数据总数
}

// boolReturn
type BoolReturn struct {
	Result bool `json:"result"` //操作结果
}

func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithCodeMsg(ctx *gin.Context, code MyCode, errMsg string) {
	rd := &ResponseData{
		Code:    code,
		Message: errMsg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}
func ResponseErrorWithMsg(ctx *gin.Context, errMsg string) {
	rd := &ResponseData{
		Code:    CodeFailure,
		Message: errMsg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithData(ctx *gin.Context, code MyCode, errMsg string, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: errMsg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
