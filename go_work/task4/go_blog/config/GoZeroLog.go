package config

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

func InitLog() {
	logconfig := logx.LogConf{
		Path:       GlobalConfig.Log.LogPath,
		Level:      GlobalConfig.Log.Level,
		Mode:       GlobalConfig.Log.Model,
		MaxSize:    GlobalConfig.Log.MaxSize,
		MaxBackups: GlobalConfig.Log.MaxBackups,
	}

	logx.MustSetup(logconfig)
}

func ReqLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		p := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 处理请求
		c.Next()

		duration := time.Since(start)

		logx.Infof("请求IP:[%s],请求方法:[%s],请求路径:[%s],请求参数:[%s],花费时间:[%s]", c.ClientIP(), c.Request.Method, p, query, duration)
	}
}
