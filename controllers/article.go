package controllers

import (
	"articles_pavlinov/database"
	"articles_pavlinov/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetArticles(c *gin.Context) {
	var articles []models.Article
	if err := database.DB.Find(&articles).Error; err != nil {
		log.Println("Error fetching articles:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
		return
	}

	log.Println("Fetched articles:", articles)
	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
    var article models.Article
    if err := database.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
        return
    }
    c.JSON(http.StatusOK, article)
}

func AddArticle(c *gin.Context) {
	log.Println("Alexey: Add Article")
	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)
	article := models.Article{Title: input.Title, Content: input.Content, UserID: userID}

	if err := database.DB.Create(&article).Error; err != nil {
		log.Println("Error creating article:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	log.Println("Created article:", article)
	c.JSON(http.StatusOK, article)
}

func RemoveArticle(c *gin.Context) {
	var article models.Article
	articleID := c.Param("id")

	if err := database.DB.Where("id = ?", articleID).First(&article).Error; err != nil {
		log.Println("Error finding article:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	if err := database.DB.Delete(&article).Error; err != nil {
		log.Println("Error deleting article:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}

	log.Println("Deleted article:", article)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
}
