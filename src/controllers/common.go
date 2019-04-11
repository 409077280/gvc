package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ResponseError(err error, ctx *gin.Context)  {
	if err != nil {
		errString := fmt.Sprintf("%s", err)
		ctx.JSON(400, gin.H{
			"code" : 1,
			"msg" : errString,
		})
	}
}