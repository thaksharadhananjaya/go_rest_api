package response

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Response struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func JSON(ctx *gin.Context, status int, data interface{}) {
    ctx.JSON(status, Response{
        Status:  status,
        Message: http.StatusText(status),
        Data:    data,
    })
}
