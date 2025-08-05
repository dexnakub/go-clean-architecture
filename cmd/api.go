package cmd

import (
	app_middleware "go-clean-achitech/internal/app/middleware"
	"go-clean-achitech/internal/app/routes"
	"go-clean-achitech/internal/configs"
	data_migrations "go-clean-achitech/internal/data/migrations"

	"github.com/gin-gonic/gin"
)

func ApiMain() {
	data_migrations.StartMigration()

	gin.SetMode(configs.GetEnv().GinModeDev)
	app := gin.Default()
	app.Use(app_middleware.CORSMiddleware())
	routes.AddPublicRoutes(app)

	app.Run(":" + configs.GetEnv().PORT)

}
