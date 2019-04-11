package backend

import (
	"controllers"
	"github.com/gin-gonic/gin"
	"models"
	"strconv"
)

func ApiOfUser(ctx *gin.Context)  {
	ApiListModel := models.ApisModel{}
	var list []int64
	apiLists, err := ApiListModel.RetrieveWhereIn(list)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "Success.",
		"data": apiLists,
	})
}

func ApiTableList(ctx *gin.Context)  {
	ApiListModel := models.ApisModel{}
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
	menus, total, err := ApiListModel.RetrieveForTable(page, limit)
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

func ApiInsert(ctx *gin.Context)  {
	ApiListModel := models.ApisModel{}
	err := ctx.BindJSON(&ApiListModel)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	result, err := ApiListModel.Create()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}

func ApiReset(ctx *gin.Context)  {
	ApiListModel := models.ApisModel{}
	err := ctx.BindJSON(&ApiListModel)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	result, err := ApiListModel.Update()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}

func ApiDelete(ctx *gin.Context)  {
	ApiListModel := models.ApisModel{}
	id, err := strconv.ParseInt(ctx.Query("id"),10,64)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ApiListModel.Id = id
	result, err := ApiListModel.Delete()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}
