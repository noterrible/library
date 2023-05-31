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
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	book := model.GetBook(id)
	if book.Id > 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "响应成功",
		})
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.InternalServerError,
		Message: "获取失败",
	})
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
	query := context.Query("q")
	books := model.SearchBook(query)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "查询书籍成功",
		Data:    books,
	})
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
// @Param count formData int true "图书数量"
// @Param category_id formData int64 true "图书种类id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books [POST]
func AddBook(context *gin.Context) {
	var book model.Book
	if err := context.ShouldBind(&book); err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.UserInfoError,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
	}
	model.AddBook(book)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "添加书籍成功",
		Data:    nil,
	})
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
// @Param count formData int true "图书数量"
// @Param category_id formData int64 true "图书种类id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books/{id} [PUT]
func UpdateBook(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	var updateBook model.Book
	if err := context.ShouldBind(&updateBook); err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.NotAcceptable,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
	}
	updateBook.Id = id
	model.UpdateBook(updateBook)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "修改书籍成功",
		Data:    nil,
	})
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
// @Param id path int64 true "借书记录的id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/user/users/records/:id [PUT]
func ReturnBook(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	model.UpdateRecordAndBook(id)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "还书成功",
	})
}

// DeleteBook godoc
//
// @Summary		管理员删除图书
// @Description	管理员删除图书
// @Tags		book
// @Produce		json
// @Param id path int64 true "书籍id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books/{id} [DELETE]
func DeleteBook(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	model.DeleteBook(id)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "删除书籍成功",
	})
}
