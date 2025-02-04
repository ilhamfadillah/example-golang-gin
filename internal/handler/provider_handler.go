package handler

import (
	"example/apps/internal/model"
	"example/apps/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProviderHandler struct {
	Service service.ProviderService
}

func NewProviderHandler(service service.ProviderService) *ProviderHandler {
	return &ProviderHandler{Service: service}
}

func (handler *ProviderHandler) GetAll(c *gin.Context) {
	providers, err := handler.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, providers)
}

func (handler *ProviderHandler) GetByUuid(c *gin.Context) {
	uuid := c.Param("uuid")

	provider, err := handler.Service.GetByUuid(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, provider)
}

func (handler *ProviderHandler) Store(c *gin.Context) {
	var provider *model.Provider

	if err := c.BindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.Service.Store(provider)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (handler *ProviderHandler) Update(c *gin.Context) {
	uuid := c.Param("uuid")

	var provider *model.Provider

	if err := c.BindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.Service.Update(uuid, provider)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
