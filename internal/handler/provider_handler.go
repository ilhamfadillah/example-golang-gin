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

/*
func (h *ProviderHandler) Store(c *gin.Context) {
	var provider *model.Provider

	if err := c.BindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid := uuid.New().String()
	now := time.Now()

	query := "INSERT INTO providers (uuid, code, name, address, phone, city, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := config.DB.Exec(query, uuid, provider.Code, provider.Name, provider.Address, provider.Phone, provider.City, now, now)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "provider inserted successfully"})
}

func (h *ProviderHandler) Update(c *gin.Context) {
	uuid := c.Param("uuid")
	now := time.Now()

	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM providers WHERE uuid = ?)"
	err := config.DB.QueryRow(checkQuery, uuid).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !exists {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": sql.ErrNoRows.Error()})
		return
	}

	var provider *model.Provider

	if err := c.BindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE providers SET code = ?, name = ?, address = ?, phone = ?, city = ?, updated_at = ? WHERE uuid = ?"
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(provider.Code, provider.Name, provider.Address, provider.Phone, provider.City, now, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "provider update successfully"})
}
*/
