package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tracker/controller/repositories"
	"tracker/infrastructure/dto/mapper"
	"tracker/infrastructure/dto/model"
)

type SettingsHandler struct {
	repo *repositories.SettingsRepository
}

func NewSettingsHandler(repo *repositories.SettingsRepository) *SettingsHandler {
	return &SettingsHandler{repo: repo}
}

func (h *SettingsHandler) GetSettings(c *gin.Context) {
	settings, err := h.repo.GetSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if settings == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Settings not found"})
		return
	}
	c.JSON(http.StatusOK, settings)
}

func (h *SettingsHandler) UpdateSettings(c *gin.Context) {
	var receivedSettings model.SettingsDTO
	if err := c.ShouldBindJSON(&receivedSettings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSettings := mapper.CalculateSettingsData(receivedSettings)

	if err := h.repo.UpdateSettings(&updatedSettings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully"})
}
