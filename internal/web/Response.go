package web

import "github.com/gin-gonic/gin"

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewResponse(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data: data,
	})
}

func NewErrorResponse(ctx *gin.Context, status int, code string, message string) {
	ctx.JSON(status, errorResponse{
		Status:  status,
		Code:    code,
		Message: message,
	})
}
