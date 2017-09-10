package AuthMiddleware

import (
	TokenService "yellowroad_library/services/TokenService"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware gin.HandlerFunc

func isTokenValid(tokenService TokenService.TokenService, token string) (*TokenService.MyCustomClaims, error) {
	claims, err := tokenService.ValidateTokenString(token)

	if err != nil {
		return nil, err
	} else {
		return claims, nil
	}
}

func New(tokenService TokenService.TokenService) AuthMiddleware {
	return func(c *gin.Context) {
		var token string

		if authorizationHeader := c.GetHeader("Authorization"); len(authorizationHeader) > 0 {
			token = c.GetHeader("Authorization")
		}

		claims, tokenError := isTokenValid(tokenService, token)

		if len(token) == 0 || claims == nil || tokenError != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"woah": "there partner",
			})
		} else {
			c.Set(TokenService.TOKEN_CLAIMS_CONTEXT_KEY, claims)
			c.Next()
		}
	}
}
