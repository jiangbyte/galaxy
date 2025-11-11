package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		baseHandler := BaseHandler{}
		baseHandler.Error(c, http.StatusUnauthorized, "Unauthorized")
	}
}
