package book_route

import (
	"net/http"
	"yellowroad_library/http/middleware/auth_middleware"

	"github.com/gin-gonic/gin"
)

func FetchSingleBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Book": "nah jk"})
		return
	}
}

func CreateBook(authMiddleware auth_middleware.AuthMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Book": "nah jk"})
	}
}
