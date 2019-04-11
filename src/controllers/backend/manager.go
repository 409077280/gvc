package backend

import (
	"controllers"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"models"
	"strconv"
)

func ManagerList(ctx *gin.Context)  {
	managerModel := models.ManagerModel{}
	managers, err := managerModel.Retrieve()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "Success.",
		"data": managers,
	})
}

func ManagerTableList(ctx *gin.Context)  {
	managerModel := models.ManagerModel{}
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
	menus, total, err := managerModel.RetrieveForTable(page, limit)
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

func ManagerInsert(ctx *gin.Context)  {
	managerModel := models.ManagerModel{}
	err := ctx.BindJSON(&managerModel)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	hash := sha256.Sum256([]byte(managerModel.Password))
	managerModel.Password = hex.EncodeToString(hash[:])
	result, err := managerModel.Create()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}

func ManagerReset(ctx *gin.Context)  {
	managerModel := models.ManagerModel{}
	err := ctx.BindJSON(&managerModel)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	hash := sha256.Sum256([]byte(managerModel.Password))
	managerModel.Password = hex.EncodeToString(hash[:])
	result, err := managerModel.Update()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}

func ManagerDelete(ctx *gin.Context)  {
	managerModel := models.ManagerModel{}
	id, err := strconv.ParseInt(ctx.Query("id"),10,64)
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	managerModel.Id = id
	result, err := managerModel.Delete()
	if err != nil {
		controllers.ResponseError(err, ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"code" : 0,
		"msg" : result,
	})
}
