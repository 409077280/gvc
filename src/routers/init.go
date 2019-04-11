package routers

import (
	"config"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"module/logger"
	"net/http"
)

func InitRouter() http.Handler {

	router := gin.Default()

	router.StaticFS("/files", http.Dir("./files"))

	if config.GetEnv().DEBUG {
		pprof.Register(router)   // 性能分析工具
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(handleErrors())  // 错误处理
	router.Use(cors.Default())	// 跨域访问

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : 1,
			"msg" : "The path can not find.",
		})
		return
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : 1,
			"msg" : "The method can not find.",
		})
		return
	})

	ApiRouter(router)

	return router
}

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(err)
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg": fmt.Sprintf("%s", err),
					})
			}
		}()
		c.Next()
	}
}