package csrf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func GenerateToken() string {
	t := time.Now().UnixNano()
	i := rand.Int()
	return fmt.Sprintf("%d_%d", t, i)
}

func Protect() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			token := GenerateToken()
			c.SetCookie("csrf_token", token, 10*60, "/", "", false, false)
			c.Set("csrf_token", token)
			c.Next()
		}
	}
}

func Validate(c *gin.Context) bool {
	cookieToken, err := c.Cookie("csrf_token")

	if err != nil {
		return false
	}

	formToken := c.PostForm("csrf_token")

	if formToken == "" || cookieToken == "" {
		return false
	}

	if formToken != cookieToken {
		return false
	}

	return true
}

func Token(c *gin.Context) string {
	value, exists := c.Get("csrf_token")

	if !exists {
		return ""
	}

	return value.(string)
}
