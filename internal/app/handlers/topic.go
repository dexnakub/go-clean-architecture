package handlers

import (
	app_helpers "go-clean-achitech/internal/app/helpers"
	domain_models "go-clean-achitech/internal/domain/models"
	topic_usecase "go-clean-achitech/internal/domain/usecase/topic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTopic(ctx *gin.Context) {
	var topic domain_models.TopicCreateModel

	if err := ctx.ShouldBindJSON(&topic); err != nil {
		ctx.JSON(http.StatusBadRequest, app_helpers.BuildFailResponse(err.Error(), nil))
		return
	}

	result := topic_usecase.CreateTopic(topic)

	if result != nil {
		ctx.JSON(http.StatusInternalServerError, app_helpers.BuildErrorResponse(result.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, app_helpers.BuildSuccessResponse("craete topic success", nil))
}

func GetTopics(ctx *gin.Context) {
	limit := app_helpers.ParseQueryInt(ctx, "limit", 10)
	offset := app_helpers.ParseQueryInt(ctx, "offset", 0)

	topics, err := topic_usecase.GetTopics(limit, offset)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, app_helpers.BuildErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, app_helpers.BuildSuccessResponse("get topics success", topics))
}
