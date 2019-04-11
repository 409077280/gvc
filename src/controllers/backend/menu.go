package backend

import (
	"controllers"
	"github.com/gin-gonic/gin"
	"models"
	"strconv"
)

func MenuList(ctx *gin.Context)  {
	menuModel := models.MenuModel{}
	menus, err := menuModel.Retrieve()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "Success.",
		"data": menus,
	})
}

func MenuTableList(ctx *gin.Context)  {
	menuModel := models.MenuModel{}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	parentId, err := strconv.Atoi(ctx.Query("parent_id"))
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	menus, total, err := menuModel.RetrieveLevelForTable(page, limit, parentId)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : "Success.",
		"data" : struct {
			Total int `json:"total"`
			Data []interface{} `json:"data"`
		}{total,
		menus},
	})
}

func MenuInsert(ctx *gin.Context)  {
	menuModel := models.MenuModel{}
	err := ctx.BindJSON(&menuModel)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	result, err := menuModel.Create()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}

func MenuReset(ctx *gin.Context)  {
	menuModel := models.MenuModel{}
	/*
	body,_ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println(string(body))
	return
	*/
	err := ctx.BindJSON(&menuModel)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	result, err := menuModel.Update()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}

func MenuDelete(ctx *gin.Context)  {
	menuModel := models.MenuModel{}
	id, err := strconv.ParseInt(ctx.Query("id"),10,64)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	menuModel.Id = id
	result, err := menuModel.Delete()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}
