package BookRoute

import (
	"net/http"
	"yellowroad_library/http/middleware/AuthMiddleware"

	"github.com/gin-gonic/gin"
)

func FetchSingleBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Book": "nah jk"})
		return
	}
}

func CreateBook(authMiddleware AuthMiddleware.AuthMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Book": "nah jk"})
	}
}
