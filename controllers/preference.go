package controllers

import (
	"articles_pavlinov/database"
	"articles_pavlinov/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetPreference(c *gin.Context) {
	var input models.Preference
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.MustGet("user_id").(uint)
	preference := models.Preference{UserID: userID, ArticleID: input.ArticleID, Liked: input.Liked}
	database.DB.Where(models.Preference{UserID: userID, ArticleID: input.ArticleID}).FirstOrCreate(&preference)
	database.DB.Model(&preference).Updates(input)
	c.JSON(http.StatusOK, preference)
}

func GetPreferences(c *gin.Context) {
	var preferences []models.Preference
	userID := c.MustGet("user_id").(uint)
	database.DB.Where("user_id = ?", userID).Find(&preferences)
	c.JSON(http.StatusOK, preferences)
}
