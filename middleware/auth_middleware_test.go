package middleware

import (
	"articles_pavlinov/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock ParseToken function
var parseTokenMock func(tokenString string) (uint, error)

func TestAuthMiddleware(t *testing.T) {
	// Initialize Gin in test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	r := gin.New()
	r.Use(AuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	tests := []struct {
		name           string
		token          string
		expectedStatus int
		setupMock      func()
	}{
		{
			name:           "No Authorization header",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
			setupMock:      func() {},
		},
		{
			name:           "Invalid Authorization header format",
			token:          "InvalidToken",
			expectedStatus: http.StatusUnauthorized,
			setupMock:      func() {},
		},
		{
			name:           "Invalid token",
			token:          "Bearer invalid.token.here",
			expectedStatus: http.StatusUnauthorized,
			setupMock: func() {
				parseTokenMock = func(tokenString string) (uint, error) {
					return 0, utils.ErrInvalidToken
				}
			},
		},
		{
			name:           "Valid token",
			token:          "Bearer valid.token.here",
			expectedStatus: http.StatusOK,
			setupMock: func() {
				parseTokenMock = func(tokenString string) (uint, error) {
					return 1, nil
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func init() {
	// Override utils.ParseToken with the mock function
	utils.ParseToken = func(tokenString string) (uint, error) {
		return parseTokenMock(tokenString)
	}
}
