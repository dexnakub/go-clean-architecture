package app_helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseQueryInt(ctx *gin.Context, key string, defaultValue int) int {
	valueStr := ctx.DefaultQuery(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil || value < 0 {
		return defaultValue
	}
	return value
}
