package logic

import "github.com/gin-gonic/gin"

// GetCategory godoc
//
// @Summary		管理员获取某个分类信息
// @Description	管理员获取某个分类信息
// @Tags		category
// @Produce		json
// @Param id path int64 true "分类id"
// @Success 200 {object} tools.Response{data=model.Category}
// @Router			/admin/categories/{id} [GET]
func GetCategory(context *gin.Context) {

}
func GetCategoryBooks(context *gin.Context) {

}

// SearchCategory godoc
//
// @Summary		搜索分类
// @Description	搜索获取分类信息
// @Tags		category
// @Produce		json
// @Param q  query string true "查询条件"
// @Success 200 {object} tools.Response{data=[]model.Category{}}
// @Router			/categories [GET]
func SearchCategory(context *gin.Context) {

}

// AddCategory godoc
//
// @Summary		添加分类
// @Description	添加分类信息
// @Tags		category
// @Produce		json
// @Param name  formData string true "分类名称"
// @Success 200 {object} tools.Response
// @Router			/categories [POST]
func AddCategory(context *gin.Context) {

}

// UpdateCategory godoc
//
// @Summary		更新分类
// @Description	更新分类信息
// @Tags		category
// @Produce		json
// @Param id path int64 true "分类id"
// @Param name  formData string true "分类名称"
// @Success 200 {object} tools.Response
// @Router			/categories/{id} [PUT]
func UpdateCategory(context *gin.Context) {

}

// DeleteCategory godoc
//
// @Summary		删除分类
// @Description	删除分类信息
// @Tags		category
// @Produce		json
// @Param id path int64 true "分类id"
// @Success 200 {object} tools.Response
// @Router			/categories/{id} [DELETE]
func DeleteCategory(context *gin.Context) {

}
