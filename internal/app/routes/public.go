package routes

import (
	"go-clean-achitech/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func AddPublicRoutes(app *gin.Engine) {
	topic := app.Group("/topic")
	{
		topic.POST("/create-topic", handlers.CreateTopic)
		topic.GET("/get-topics", handlers.GetTopics)
	}

	topicItem := app.Group("/topic-item")
	{
		topicItem.POST("/create-item", handlers.CreateItem)
		topicItem.GET("/get-items/:topicID", handlers.GetItems)
		topicItem.PUT("/update-item", handlers.UpdateItem)
		topicItem.PUT("/update-delete-status", handlers.UpdateDeleteStatus)
		topicItem.PUT("/update-sequence", handlers.UpdateSequence)
		topicItem.DELETE("/delete-item/:topicID/:itemID", handlers.Deleteitem)
	}
}
