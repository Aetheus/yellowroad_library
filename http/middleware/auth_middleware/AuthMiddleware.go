package auth_middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yellowroad_library/services/auth_domain"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type AuthMiddleware gin.HandlerFunc

const USER_KEY = "LOGGED_IN_USER"

func GetUser (c *gin.Context) (user entities.User, err app_error.AppError){
	potentialUser, exists := c.Get(USER_KEY)

	if !exists {
		err = app_error.New(http.StatusUnauthorized,
				"",
			"You need to be logged in to access this!")
		return
	}

	user = potentialUser.(entities.User)
	return
}

func New(getLoggedInUser auth_domain.GetLoggedInUser) AuthMiddleware {

	return func(c *gin.Context) {
		var token string

		if authorizationHeader := c.GetHeader("Authorization"); len(authorizationHeader) > 0 {
			token = c.GetHeader("Authorization")
		}

		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"message" : "No valid login token provided!",
			})
			return
		}

		user, err := getLoggedInUser.Execute(token)
		if err != nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"message" : err.EndpointMessage(),
			})
			return
		}

		c.Set(USER_KEY, user)
		c.Next()
	}
}
