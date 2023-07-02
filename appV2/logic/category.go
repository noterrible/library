package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appV2/model"
	"libraryManagementSystem/appV2/tools"
	"net/http"
	"strconv"
)

// GetCategory godoc
//
// @Summary		管理员获取某个分类信息
// @Description	管理员获取某个分类信息
// @Tags		admin/categories
// @Produce		json
// @Param id path int64 true "分类id"
// @Success 200 {object} tools.Response{data=model.Category}
// @Router			/admin/categories/{id} [GET]
func GetCategory(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	category := model.GetCategory(id)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "获取成功",
		Data:    category,
	})
}

// PageCategory godoc
//
// @Summary		分页分类
// @Description	获取分类信息
// @Tags		public
// @Accept		json
// @Produce		json
// @Param pageInfo  query model.PageInfo true "分页信息"
// @Success 200 {object} tools.Response{data=model.ListResponse[model.Category]{}}
// @Router			/categories [GET]
func PageCategory(context *gin.Context) {
	//普通分页
	var pageInfo model.PageInfo
	if err := context.ShouldBind(&pageInfo); err != nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "错误" + err.Error(),
			Data:    nil,
		})
		return
	}
	var page model.ListResponse[model.Category]
	err := model.Pages(&page, pageInfo)
	if err != nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "错误" + err.Error(),
			Data:    nil,
		})
		return
	}
	if len(page.List) == 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "没有此页数据",
			Data:    nil,
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "查询分类成功",
		Data:    page,
	})
}

// AddCategory godoc
//
// @Summary		添加分类
// @Description	添加分类信息
// @Tags		admin/categories
// @Accept		multipart/form-data
// @Produce		json
// @Param name formData string true "分类名称"
// @Success 200 {object} tools.Response
// @Router			/admin/categories [POST]
func AddCategory(context *gin.Context) {
	category := model.Category{}
	category.Name = context.PostForm("name")
	//这里为什么绑定不成功？
	//if err := context.ShouldBind(&category); err != nil {
	//	context.JSON(http.StatusNotAcceptable, tools.Response{
	//		Code:    tools.UserInfoError,
	//		Message: "绑定失败" + err.Error(),
	//	})
	//	return
	//}
	model.AddCategory(category)
	context.JSON(http.StatusCreated, tools.Response{
		Code:    tools.OK,
		Message: "添加分类成功",
	})
}

// UpdateCategory godoc
//
// @Summary		更新分类
// @Description	更新分类信息
// @Tags		admin/categories
// @Accept		multipart/form-data
// @Produce		json
// @Param id path int64 true "分类id"
// @Param name formData string true "分类名称"
// @Success 200 {object} tools.Response
// @Router			/admin/categories/{id} [PUT]
func UpdateCategory(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	var updateCategory model.Category
	updateCategory.Name = context.PostForm("name")
	//绑定失败同上
	//if err := context.ShouldBind(&updateCategory); err != nil {
	//	context.JSON(http.StatusNotAcceptable, tools.Response{
	//		Code:    tools.UserInfoError,
	//		Message: "绑定失败" + err.Error(),
	//		Data:    nil,
	//	})
	//	return
	//}
	fmt.Println("id:", id)
	updateCategory.Id = id
	err := model.UpdateCategory(updateCategory)
	if err != nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "更新失败" + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "更新成功",
	})
}

// DeleteCategory godoc
//
// @Summary		删除分类
// @Description	删除分类信息
// @Tags		admin/categories
// @Produce		json
// @Param id path int64 true "分类id"
// @Success 200 {object} tools.Response
// @Router			/admin/categories/{id} [DELETE]
func DeleteCategory(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	err := model.DeleteCategory(id)
	if err != nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "删除失败失败，" + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "删除成功",
	})
}
