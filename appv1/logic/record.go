package logic

import (
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appv1/model"
	"libraryManagementSystem/appv1/tools"
	"net/http"
	"strconv"
)

// GetRecords godoc
//
// @Summary		获取用户信息
// @Description	获取用户信息
// @Tags		user
// @Produce		json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path int64 true "用户Id"
// @Success 200 {object} tools.Response{data=[]model.Record{}}
// @Router			/user/users/{id}/records [GET]
func GetRecords(context *gin.Context) {
}

// GetUserRecords godoc
//
// @Summary		获取某个用户的所有记录
// @Description	获取用户所有的借还记录
// @Tags		user
// @Produce		json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path int64 true "用户Id"
// @Success 200 {object} tools.Response{data=[]model.Record{}}
// @Router			/user/users/{id}/records [GET]
func GetUserRecords(context *gin.Context) {
	userIdString := context.Param("id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	records := model.GetAllRecordsByUserId(userId)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "查询成功",
		Data:    records,
	})
}

// GetUserStatusRecords godoc
//
// @Summary		获取用户借/还记录
// @Description	获取某个用户的借/还记录
// @Tags		user
// @Produce		json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path int64 true "用户Id"
// @Param status path int true "标记是否归还字段"
// @Success 200 {object} tools.Response{data=[]model.Record{}}
// @Router			/user/users/{id}/records/{status} [GET]
// @Router			/admin/users/{id}/records/{status} [GET]
func GetUserStatusRecords(context *gin.Context) {
	userIdString := context.Param("id")
	statusString := context.Param("status")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	status, _ := strconv.Atoi(statusString)
	records := model.GetStatusRecordsByUserId(userId, status)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "查询成功",
		Data:    records,
	})
}

// GetStatusRecords godoc
//
// @Summary		获取所有借/还记录
// @Description	获取图书馆所有的借/还记录
// @Tags		Record
// @Produce		json
// @Param status path int true "标记是否归还字段"
// @Success 200 {object} tools.Response{data=[]model.Record{}}
// @Router			/records/{status} [GET]
func GetStatusRecords(context *gin.Context) {
	statusString := context.Param("status")
	status, _ := strconv.Atoi(statusString)
	records := model.GetStatusRecords(status)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "查询成功",
		Data:    records,
	})
}
