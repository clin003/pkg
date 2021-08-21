package handler

import (
	"net/http"

	"gitee.com/lyhuilin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Json
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	//always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// Redirect
func SendRedirect(c *gin.Context, data string) {
	c.Redirect(http.StatusMovedPermanently, data)
}

// String
func SendString(c *gin.Context, data string) {
	//always return http.StatusOK
	c.String(http.StatusOK, data)
}
