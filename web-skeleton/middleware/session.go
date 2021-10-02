package middleware

import (
	"context"

	"github.com/Summer2021-Project/web-skeleton/di"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := di.Session()
		store, err := session.Start(context.Background(), c.Writer, c.Request)
		if err != nil {
			panic(err)
		}
		if _, ok := store.Get("userinfo"); !ok {
			c.Status(401)
			c.Abort()
			return
		}

		c.Next()
	}
}
