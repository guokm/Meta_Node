package middleware

import (
	"fmt"
	"github/metanode/go_blog/config"
	"github/metanode/go_blog/model/vo"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

// 生成jwt --TODO
func GenerateToken(customClaims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// 验证jwt
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		logx.Info("tokenString is " + tokenString)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, vo.Error(vo.CodeParamError, "missing token"))
			return
		}

		var claims CustomClaims
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("非预期的签名算法: %v", token.Header["alg"])
			}
			return []byte(config.GlobalConfig.Jwt.Secret), nil
		})
		// 处理解析错误
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, vo.Error(vo.CodeParamError, "令牌无效或已过期"))
			return
		}

		// 验证签发者（可选）
		if claims.Issuer != config.GlobalConfig.Jwt.Issuer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, vo.Error(vo.CodeParamError, "签发机构错误"))
			return
		}

		c.Set("ID", claims.ID)
		c.Set("Username", claims.Username)
		c.Next()
	}
}

type CustomClaims struct {
	ID       uint   `form:"id"`
	Username string `form:"username"`
	jwt.RegisteredClaims
}
