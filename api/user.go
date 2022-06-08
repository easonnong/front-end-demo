package api

import (
	"github.com/easonnong/front-end-demo/dao"
	"github.com/easonnong/front-end-demo/models"
	"github.com/easonnong/front-end-demo/response"
	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		response.Failed("参数错误", ctx)
		return
	}

	dao.Mgr.AddUser(&user)

	response.Success("添加用户成功", user, ctx)
}

func ListUser(ctx *gin.Context) {
	users := dao.Mgr.ListUser()
	response.Success("查询成功", users, ctx)
}
