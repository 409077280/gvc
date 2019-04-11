package backend

import (
	"controllers"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"models"
	"strconv"
)

func UserInfo(ctx *gin.Context)  {
	usersModel := models.UsersModel{}
	id, err := strconv.ParseInt(ctx.Query("id"),10,64)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	usersModel.Id = id
	userInfo, err := usersModel.RetrieveIndividual()
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(200, gin.H{
				"code": 0,
				"msg":  "Success.",
				"data": []interface{}{},
			})
			return
		} else{
			controllers.ResponseError(err, ctx)
			return
		}
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "Success.",
		"data": userInfo,
	})
}

func UserTableList(ctx *gin.Context)  {
	usersModel := models.UsersModel{}
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
	users, total, err := usersModel.RetrieveForTable(page, limit)
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
			users},
	})
}

func UserReferrer(ctx *gin.Context)  {
	usersModel := models.UsersModel{}
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
	refferer, err := strconv.Atoi(ctx.Query("refferer"))
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	users, total, err := usersModel.RetrieveLevelForTable(page, limit, refferer)
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
			users},
	})
}

func UserInsert(ctx *gin.Context)  {
	usersModel := models.UsersModel{}
	err := ctx.BindJSON(&usersModel)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	result, err := usersModel.Create()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}

func UserReset(ctx *gin.Context)  {
	usersModel := models.UsersModel{}
	err := ctx.BindJSON(&usersModel)
	if err != nil {
		fmt.Println(err)
		controllers.ResponseError(err, ctx)
		return
	}
	result, err := usersModel.Update()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}

func UserDelete(ctx *gin.Context)  {
	usersModel := models.UsersModel{}
	id, err := strconv.ParseInt(ctx.Query("id"),10,64)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	usersModel.Id = id
	result, err := usersModel.Delete()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}
