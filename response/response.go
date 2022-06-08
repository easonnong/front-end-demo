package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{200, message, data})
}

func Failed(message string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{400, message, nil})
}
