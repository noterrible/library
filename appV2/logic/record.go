package logic

import (
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appV2/model"
	"libraryManagementSystem/appV2/tools"
	"net/http"
	"strconv"
)

// GetUserRecordsByAdmin godoc
//
// @Summary		获取用户信息
// @Description	获取用户信息
// @Tags		admin/users
// @Produce		json
// @Param id path int64 true "用户Id"
// @Success 200 {object} tools.Response{data=[]model.Record{}}
// @Router			/admin/users/{id}/records [GET]
func GetUserRecordsByAdmin(context *gin.Context) {
	userIdString := context.Param("id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	records := model.GetAllRecordsByUserId(userId)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "查询成功",
		Data:    records,
	})
}

// GetUserRecords godoc
//
// @Summary		用户获取所有记录或者借/还记录
// @Description	获取用户所有记录或者借/还记录
// @Tags		user/users
// @Produce		json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param status query string false "是否归还"
// @Success 200 {object} tools.Response{data=[]model.Record{}}
// @Router			/user/users/records [GET]
func GetUserRecords(context *gin.Context) {
	userIdString, _ := context.Cookie("id")
	statusString := context.Query("status")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	//传空值，查所有记录
	if statusString == "" {
		records := model.GetAllRecordsByUserId(userId)
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "查询成功",
			Data:    records,
		})
		return
	}
	//非空查借/换记录
	status, _ := strconv.Atoi(statusString)
	records := model.GetStatusRecordsByUserId(userId, status)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "查询成功",
		Data:    records,
	})
}

// GetUserStatusRecordsByAdmin godoc
//
// @Summary		获取用户借/还记录
// @Description	获取某个用户的借/还记录
// @Tags		admin/users
// @Produce		json
// @Param id path int64 true "用户Id"
// @Param status path int true "标记是否归还字段"
// @Success 200 {object} tools.Response{data=[]model.Record{}}
// @Router			/admin/users/{id}/records/{status} [GET]
func GetUserStatusRecordsByAdmin(context *gin.Context) {
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
// @Tags		admin/records
// @Produce		json
// @Param status path int true "标记是否归还字段"
// @Success 200 {object} tools.Response{data=[]model.Record{}}
// @Router			/admin/records/{status} [GET]
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
