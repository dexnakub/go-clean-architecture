package handlers

import (
	app_helpers "go-clean-achitech/internal/app/helpers"
	domain_models "go-clean-achitech/internal/domain/models"
	topic_item_usecase "go-clean-achitech/internal/domain/usecase/topic_item"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(ctx *gin.Context) {
	var topic domain_models.TopicItemCreateModel

	if err := ctx.ShouldBindJSON(&topic); err != nil {
		ctx.JSON(http.StatusBadRequest, app_helpers.BuildFailResponse(err.Error(), nil))
		return
	}

	result := topic_item_usecase.CreateItem(topic)

	if result != nil {
		ctx.JSON(http.StatusBadRequest, app_helpers.BuildErrorResponse(result.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, app_helpers.BuildSuccessResponse("craete topic success", nil))
}

func GetItems(ctx *gin.Context) {
	topicID := ctx.Param("topicID")

	limit := app_helpers.ParseQueryInt(ctx, "limit", 10)
	offset := app_helpers.ParseQueryInt(ctx, "offset", 0)

	topics, err := topic_item_usecase.GetItems(topicID, limit, offset)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, app_helpers.BuildErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, app_helpers.BuildSuccessResponse("get topics success", topics))
}

func UpdateItem(ctx *gin.Context) {
	var topicItem domain_models.TopicItemUpdateModel
	if err := ctx.ShouldBindJSON(&topicItem); err != nil {
		ctx.JSON(http.StatusBadRequest, app_helpers.BuildFailResponse("invalid request body", nil))
		return
	}

	if err := topic_item_usecase.Updateitem(topicItem); err != nil {
		ctx.JSON(http.StatusInternalServerError, app_helpers.BuildErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, app_helpers.BuildSuccessResponse("update item success", nil))
}

func UpdateDeleteStatus(ctx *gin.Context) {
	var topicItem domain_models.TopicItemUpdateDeleteStatusModel
	if err := ctx.ShouldBindJSON(&topicItem); err != nil {
		ctx.JSON(http.StatusBadRequest, app_helpers.BuildFailResponse("invalid request body", nil))
		return
	}

	if err := topic_item_usecase.UpdateDeleteStatus(topicItem); err != nil {
		ctx.JSON(http.StatusInternalServerError, app_helpers.BuildErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, app_helpers.BuildSuccessResponse("update delete status success", nil))
}

func UpdateSequence(ctx *gin.Context) {
	var topicItem []domain_models.TopicItemUpdateSequenceModel
	if err := ctx.ShouldBindJSON(&topicItem); err != nil {
		ctx.JSON(http.StatusBadRequest, app_helpers.BuildFailResponse("invalid request body", nil))
		return
	}

	if err := topic_item_usecase.UpdateSequence(topicItem); err != nil {
		ctx.JSON(http.StatusInternalServerError, app_helpers.BuildErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, app_helpers.BuildSuccessResponse("update sequence success", nil))
}

func Deleteitem(ctx *gin.Context) {
	topicID := ctx.Param("topicID")
	itemID := ctx.Param("itemID")

	if itemID == "" || topicID == "" {
		ctx.JSON(http.StatusBadRequest, app_helpers.BuildFailResponse("itemID and topicID are required", nil))
		return
	}

	if err := topic_item_usecase.DeleteItem(topicID, itemID); err != nil {
		ctx.JSON(http.StatusInternalServerError, app_helpers.BuildErrorResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "topic item deleted successfully",
	})
}
