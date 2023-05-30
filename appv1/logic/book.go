package logic

import (
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appv1/model"
	"libraryManagementSystem/appv1/tools"
	"net/http"
	"strconv"
	"time"
)

// GetBook godoc
//
// @Summary		获取图书信息
// @Description	获取一个图书的信息
// @Tags		book
// @Produce		json
// @Param id path int64 true "书籍id"
// @Success 200 {object} tools.Response{data=model.Book}
// @Router			/books/{id} [GET]
func GetBook(context *gin.Context) {

}

// SearchBook godoc
//
// @Summary		搜索图书
// @Description	获取所有图书或者搜索图书
// @Tags		book
// @Produce		json
// @Param q  query string true "查询条件"
// @Success 200 {object} tools.Response{data=[]model.Book{}}
// @Router			/books [GET]
func SearchBook(context *gin.Context) {
	//if
}

// AddBook godoc
//
// @Summary		新增图书
// @Description	管理员添加图书
// @Tags		book
// @Accept		multipart/form-data
// @Produce		json
// @Param bn formData string true "图书编号"
// @Param name formData string true "图书名称"
// @Param description formData string true "图书描述"
// @Param count formData string true "图书数量"
// @Param category_id formData string true "图书种类id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books [POST]
func AddBook(context *gin.Context) {

}

// UpdateBook godoc
//
// @Summary		修改图书
// @Description	管理员修改图书
// @Tags		book
// @Accept		multipart/form-data
// @Produce		json
// @Param id path int64 true "图书id"
// @Param bn formData string true "图书编号"
// @Param name formData string true "图书名称"
// @Param description formData string true "图书描述"
// @Param count formData string true "图书数量"
// @Param category_id formData string true "图书种类id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books/{id} [PUT]
func UpdateBook(context *gin.Context) {

}

// BorrowBook godoc
//
// @Summary		用户借书
// @Description	用户借书
// @Tags		book
// @Produce		json
// @Param Authorization header string false "Bearer 用户令牌"
// @CookieParam id  string true "用户id"
// @Param bookId path string true "书籍id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/user/users/records/:bookId [POST]
func BorrowBook(context *gin.Context) {
	userIdString, _ := context.Cookie("id")
	bookIdString := context.Param("bookId")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	bookId, _ := strconv.ParseInt(bookIdString, 10, 64)
	record := model.Record{
		UserId:    userId,
		BookId:    bookId,
		StartTime: time.Now(),
		OverTime:  time.Now().Add(tools.T),
	}
	model.CreateRecord(record)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "借书成功",
	})
}

// ReturnBook godoc
//
// @Summary		用户还书
// @Description	用户还书
// @Tags		book
// @Accept		multipart/form-data
// @Produce		json
// @Param bookId path string true "书籍id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/user/users/records/:bookId [PUT]
func ReturnBook(context *gin.Context) {

}
func GetStatusBook(context *gin.Context) {

}

// DeleteBook godoc
//
// @Summary		管理员删除图书
// @Description	管理员删除图书
// @Tags		book
// @Produce		json
// @Param id path string true "书籍id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books/{id} [DELETE]
func DeleteBook(context *gin.Context) {

}
