package logic

import (
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appV2/model"
	"libraryManagementSystem/appV2/tools"
	"net/http"
	"strconv"
)

// GetMessage godoc
//
// @Summary		用户收件箱
// @Description	获取用户收件箱信息(这个接口获取消息没有什么意义，只是提示有书没还)
// @Tags		user/users
// @Produce		json
// @CookieParam id string true "用户id"
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} tools.Response
// @Router			/user/users/messages [GET]
func GetMessage(context *gin.Context) {
	userIdString, _ := context.Cookie("id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	messages := model.GetMessage(userId)
	if len(messages) > 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "您有书没有还",
			Data:    messages,
		})
	}
}
