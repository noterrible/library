package logic

import (
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appv1/model"
	"libraryManagementSystem/appv1/tools"
	"net/http"
)

//type Session struct {
//	SessionId string `json:"sessionId"`
//}

// LibrarianLogin godoc
//
// @Summary		图书管理员登录
// @Description	会执行图书管理员登录操作
// @Tags		Librarian
// @Accept		multipart/form-data
// @Produce		json
// @Param userName formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/adminLogin [POST]
func LibrarianLogin(context *gin.Context) {
	var admin model.Librarian
	if err := context.ShouldBind(admin); err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.UserInfoError,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
	}
	if admin = model.GetAdmin(admin.UserName, admin.Password); admin.Id > 0 {
		err := model.SetSession(context, admin.Id, admin.UserName)
		if err != nil {
			context.JSON(http.StatusInternalServerError, tools.Response{
				Code:    tools.InternalServerError,
				Message: "登陆失败" + err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "登陆成功",
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "登陆失败",
	})
	return
}
