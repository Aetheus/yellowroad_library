package auth_middleware

import (
	TokenService "yellowroad_library/services/token_serv"

	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware gin.HandlerFunc

func isTokenValid(tokenService TokenService.TokenService, token string) (TokenService.LoginClaim, error) {
	claims, err := tokenService.ValidateTokenString(token)

	if err != nil {
		return TokenService.LoginClaim{}, err
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

		if len(token) == 0 || tokenError != nil {
			//TODO : return an actual JSON with more info (e.g: a message)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"message" : "No valid login token provided!",
			})
		} else {
			c.Set(TokenService.TOKEN_CLAIMS_CONTEXT_KEY, claims)
			c.Next()
		}
	}
}
