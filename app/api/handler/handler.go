package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type baseResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, base baseResponse, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": base.Code,
		"msg":  base.Message,
		"data": data,
	})
}

func BuildSuccessResponse(c *gin.Context, code int64, msg string, data interface{}) {
	Response(c, baseResponse{Code: code, Message: msg}, data)
}

func BuildFailResponse(c *gin.Context, code int64, msg string) {
	Response(c, baseResponse{Code: code, Message: msg}, nil)
}
