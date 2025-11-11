// pkg/middleware/auth.go
package middleware

import (
	"galaxy/pkg/service"
	"galaxy/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"data":    nil,
				"message": "未提供认证令牌",
				"success": false,
			})
			c.Abort()
			return
		}

		// 去掉 "Bearer " 前缀
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		// 检查token是否在黑名单中
		banned, err := service.Blacklist.IsInBlacklist(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"data":    nil,
				"message": "服务器内部错误",
				"success": false,
			})
			c.Abort()
			return
		}

		if banned {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"data":    nil,
				"message": "认证令牌已被封禁",
				"success": false,
			})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"data":    nil,
				"message": "无效的认证令牌",
				"success": false,
			})
			c.Abort()
			return
		}

		// 检查用户token是否被强制失效
		invalidationTime, err := service.Blacklist.GetUserInvalidationTime(claims.AccountId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"data":    nil,
				"message": "服务器内部错误",
				"success": false,
			})
			c.Abort()
			return
		}

		if invalidationTime > 0 && claims.IssuedAt != nil {
			// 如果token的签发时间早于失效时间，则token无效
			if claims.IssuedAt.Time.Unix() < invalidationTime {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"data":    nil,
					"message": "认证令牌已失效",
					"success": false,
				})
				c.Abort()
				return
			}
		}

		c.Set("account_id", claims.AccountId)
		c.Set("token", token) // 将token存入context，方便后续使用
		c.Next()
	}
}

// OptionalAuthMiddleware 可选认证中间件（不强制要求认证）
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Next()
			return
		}

		// 去掉 "Bearer " 前缀
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		//if claims, err := utils.ParseToken(token); err == nil {
		//	c.Set("user_id", claims.AccountId)
		//}

		// 检查token是否在黑名单中
		banned, err := service.Blacklist.IsInBlacklist(token)
		if err == nil && banned {
			c.Next() // 即使被封禁，也继续处理（但不设置user_id）
			return
		}

		if claims, err := utils.ParseToken(token); err == nil {
			// 检查用户token是否被强制失效
			invalidationTime, err := service.Blacklist.GetUserInvalidationTime(claims.AccountId)
			if err == nil {
				if invalidationTime > 0 && claims.IssuedAt != nil {
					if claims.IssuedAt.Time.Unix() >= invalidationTime {
						c.Set("user_id", claims.AccountId)
						c.Set("token", token)
					}
				} else {
					c.Set("user_id", claims.AccountId)
					c.Set("token", token)
				}
			}
		}

		c.Next()
	}
}
