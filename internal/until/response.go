package until

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

/*
HTTPResponse инициализирует структуру Response ответа

context - Контекст

code - ответа

message - сообщения

err - ошибка

payload - данные
*/
func HTTPResponse(c *gin.Context, code int, message string, err error, payload interface{}) {

	c.Request.Header.Add("Content-Type", "application/json")

	response := &Response{
		Code:    code,
		Message: message,
		Data:    payload,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(code, response)
}
