package routers

import (
	"controllers/backend"
	"github.com/gin-gonic/gin"
	"middleware/auth"
)

/*
api 参数通过Context的Param方法来获取

router.GET("/string/:name", func(c *gin.Context) {
    	name := c.Param("name")
    	fmt.Println("Hello %s", name)
    })
 */
func ApiRouter(router *gin.Engine) {
	apirouter := router.Group("/api")

		apirouter.GET("/menulist", backend.MenuList)
		apirouter.GET("/menutablelist", backend.MenuTableList)
		apirouter.POST("/menuinsert", backend.MenuInsert)
		apirouter.PUT("/menureset", backend.MenuReset)
		apirouter.DELETE("/menudelete", backend.MenuDelete)

		apirouter.GET("/managerlist", backend.ManagerList)
		apirouter.GET("/managertablelist", backend.ManagerTableList)
		apirouter.POST("/managerinsert", backend.ManagerInsert)
		apirouter.PUT("/managerreset", backend.ManagerReset)
		apirouter.DELETE("/managerdelete", backend.ManagerDelete)

		apirouter.GET("/apiofuser", backend.ApiOfUser)
		apirouter.GET("/apitablelist", backend.ApiTableList)
		apirouter.POST("/apiinsert", backend.ApiInsert)
		apirouter.PUT("/apireset", backend.ApiReset)
		apirouter.DELETE("/apidelete", backend.ApiDelete)

		apirouter.GET("/userinfo", backend.UserInfo)
		apirouter.GET("/userreferrer", backend.UserReferrer)
		apirouter.GET("/usertablelist", backend.UserTableList)
		apirouter.POST("/userinsert", backend.UserInsert)
		apirouter.PUT("/userreset", backend.UserReset)
		apirouter.DELETE("/userdelete", backend.UserDelete)

	apirouter.GET("/login", backend.Login)
	// cookie auth middleware
	apirouter.Use(auth.Middleware("cookie"))
	{
	}

	/*
	jwtApi := router.Group("/api")
	jwtApi.GET("/jwt/set/:userid", controllers.JwtSetExample)

	// jwt auth middleware
	jwtApi.Use(auth.Middleware("jwt"))
	{
		jwtApi.GET("/jwt/get", controllers.JwtGetExample)
	}
	*/
}
