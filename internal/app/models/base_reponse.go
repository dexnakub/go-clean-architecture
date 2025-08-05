package app_models

import "github.com/gin-gonic/gin"

type BaseResponseModel struct {
	Success      bool    `json:"success"`
	Message      *string `json:"message,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Data         any     `json:"data,omitempty"`
}

func (data *BaseResponseModel) ToGinMap() gin.H {
	return gin.H{
		"success":      data.Success,
		"message":      data.Message,
		"errorMessage": data.ErrorMessage,
		"data":         data.Data,
	}
}
