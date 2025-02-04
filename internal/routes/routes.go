package routes

import (
	"example/apps/config"
	"example/apps/internal/handler"
	"example/apps/internal/repository"
	"example/apps/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	providerRepo := repository.NewProviderRepository(config.DB)
	providerService := service.NewProviderService(providerRepo)
	providerHandler := handler.NewProviderHandler(providerService)

	r.GET("/v1/providers", providerHandler.GetAll)
	r.GET("/v1/providers/:uuid", providerHandler.GetByUuid)
	r.POST("/v1/providers", providerHandler.Store)
	r.PUT("/v1/providers/:uuid", providerHandler.Update)
}
