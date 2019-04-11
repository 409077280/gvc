package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"middleware/auth/drivers"
)
func Login(ctx *gin.Context)  {
	userInfo := map[string]interface{}{
		"location": "backend",
		"username": "admin",
	}
	jwt := drivers.NewJwtAuthDriver()
	token := jwt.Login(ctx.Request, ctx.Writer, userInfo)
	fmt.Println(token)
}

func Logout(ctx *gin.Context) {
	ctx.Abort()
}
