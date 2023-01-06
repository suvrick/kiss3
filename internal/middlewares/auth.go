package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suvrick/kiss/internal/jwthelper"
	"github.com/suvrick/kiss/internal/until"
)

//const notAuth [2]string = []string{"/api/user/new", "/api/user/login"}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method == "POST" {
			notAuth := []string{"/api/auth/login", "/api/auth/register"}
			requestPath := c.Request.URL.Path
			for _, value := range notAuth {
				if value == requestPath {
					c.Next()
					return
				}
			}
		}

		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			until.HTTPResponse(c, http.StatusUnauthorized, "Authentication fail.Invalid token.", nil, nil)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			until.HTTPResponse(c, http.StatusUnauthorized, "Authentication fail. Invalid token.", nil, nil)
			c.Abort()
			return
		}

		tk := splitted[1]
		userID, ok := jwthelper.Parse(tk)
		if !ok {
			until.HTTPResponse(c, http.StatusUnauthorized, "Authentication fail. Bad token.", nil, nil)
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
