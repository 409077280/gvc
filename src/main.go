package main

import (
	"config"
	"github.com/gin-gonic/gin"
	"runtime"
	"service"
)

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	service.StartService()
}
