package BookRoute

import (
	"net/http"
	"yellowroad_library/http/middleware/authMiddleware"

	"github.com/gin-gonic/gin"
)

func FetchSingleBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Book": "nah jk"})
		return
	}
}

func CreateBook(authMiddleware authMiddleware.AuthMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Book": "nah jk"})
	}
}
