package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"articles_pavlinov/database"
	"articles_pavlinov/models"
	"articles_pavlinov/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Setup
	database.SetupDatabase()
	database.DB.Exec("PRAGMA foreign_keys = ON")
	database.DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Preference{})
	// Run tests
	m.Run()
	// Teardown
	database.DB.Exec("DROP TABLE users")
	database.DB.Exec("DROP TABLE articles")
	database.DB.Exec("DROP TABLE preferences")
}

func performRequest(r http.Handler, method, path string, body interface{}, token string) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRegisterAndLogin(t *testing.T) {
	r := routes.SetupRouter()

	// Register
	registerPayload := gin.H{
		"username": "user1",
		"password": "password",
	}
	w := performRequest(r, "POST", "/auth/register", registerPayload, "")
	assert.Equal(t, http.StatusOK, w.Code)

	// Login
	loginPayload := gin.H{
		"username": "user1",
		"password": "password",
	}
	w = performRequest(r, "POST", "/auth/login", loginPayload, "")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response["token"])
}

func TestAddArticle(t *testing.T) {
	r := routes.SetupRouter()

	// Register and Login
	registerPayload := gin.H{
		"username": "user1",
		"password": "password",
	}
	performRequest(r, "POST", "/auth/register", registerPayload, "")
	loginPayload := gin.H{
		"username": "user1",
		"password": "password",
	}
	w := performRequest(r, "POST", "/auth/login", loginPayload, "")

	var loginResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &loginResponse)
	token := loginResponse["token"]

	// Add Article
	articlePayload := gin.H{
		"title":   "Article 1",
		"content": "Content of Article 1",
	}
	w = performRequest(r, "POST", "/articles", articlePayload, token)
	assert.Equal(t, http.StatusOK, w.Code)

	var articleResponse models.Article
	err := json.Unmarshal(w.Body.Bytes(), &articleResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Article 1", articleResponse.Title)
	assert.Equal(t, "Content of Article 1", articleResponse.Content)
}

func TestGetArticle(t *testing.T) {
    r := routes.SetupRouter()

    // Register and Login
    registerPayload := gin.H{
        "username": "user1",
        "password": "password",
    }
    performRequest(r, "POST", "/auth/register", registerPayload, "")
    loginPayload := gin.H{
        "username": "user1",
        "password": "password",
    }
    w := performRequest(r, "POST", "/auth/login", loginPayload, "")

    var loginResponse map[string]string
    json.Unmarshal(w.Body.Bytes(), &loginResponse)
    token := loginResponse["token"]

    // Add Article
    articlePayload := gin.H{
        "title":   "Article 1",
        "content": "Content of Article 1",
    }
    w = performRequest(r, "POST", "/articles", articlePayload, token)

    var articleResponse models.Article
    json.Unmarshal(w.Body.Bytes(), &articleResponse)

    // Get Article
    w = performRequest(r, "GET", "/articles/"+string(articleResponse.ID), nil, token)
    assert.Equal(t, http.StatusOK, w.Code)

    var getArticleResponse models.Article
    err := json.Unmarshal(w.Body.Bytes(), &getArticleResponse)
    assert.NoError(t, err)
    assert.Equal(t, articleResponse.Title, getArticleResponse.Title)
    assert.Equal(t, articleResponse.Content, getArticleResponse.Content)
}


func TestSetPreference(t *testing.T) {
	r := routes.SetupRouter()

	// Register and Login
	registerPayload := gin.H{
		"username": "user1",
		"password": "password",
	}
	performRequest(r, "POST", "/auth/register", registerPayload, "")
	loginPayload := gin.H{
		"username": "user1",
		"password": "password",
	}
	w := performRequest(r, "POST", "/auth/login", loginPayload, "")

	var loginResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &loginResponse)
	token := loginResponse["token"]

	// Add Article
	articlePayload := gin.H{
		"title":   "Article 1",
		"content": "Content of Article 1",
	}
	w = performRequest(r, "POST", "/articles", articlePayload, token)

	var articleResponse models.Article
	json.Unmarshal(w.Body.Bytes(), &articleResponse)

	// Set Preference
	preferencePayload := gin.H{
		"article_id": articleResponse.ID,
		"liked":      true,
	}
	w = performRequest(r, "POST", "/preferences", preferencePayload, token)
	assert.Equal(t, http.StatusOK, w.Code)

	var preferenceResponse models.Preference
	err := json.Unmarshal(w.Body.Bytes(), &preferenceResponse)
	assert.NoError(t, err)
	assert.Equal(t, articleResponse.ID, preferenceResponse.ArticleID)
	assert.Equal(t, true, preferenceResponse.Liked)
}
