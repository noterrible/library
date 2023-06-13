package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appV2/model"
	"libraryManagementSystem/appV2/tools"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetBook godoc
//
// @Summary		获取图书信息
// @Description	获取一个图书的信息
// @Tags		public
// @Produce		json
// @Param id path int64 true "书籍id"
// @Success 200 {object} tools.Response{data=model.BookInfo}
// @Router			/books/{id} [GET]
func GetBook(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	book := model.GetBook(id)
	if book.Id > 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "响应成功",
			Data:    book,
		})
		return
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
// @Tags		public
// @Produce		json
// @Param currentPage  query string true "当前页"
// @Param pageSize  query string true "页大小"
// @Param ISBN  query string false "书籍编号"
// @Param bookName  query string false "图书名称"
// @Success 200 {object} tools.Response{data=tools.Page[model.BookInfo]{}}
// @Router			/books [GET]
func SearchBook(context *gin.Context) {
	ISBN := context.Query("ISBN")
	bookName := context.Query("bookName")
	currentPageString := context.Query("currentPage")
	pageSizeString := context.Query("pageSize")
	//查缓存
	key := fmt.Sprintf("books_%v_%v_%v_%v", currentPageString, pageSizeString, ISBN, bookName)
	pageRedisJson, err := model.InfoCacheRedisConn.Get(context, key).Result()
	// 反序列化数据为结构体
	var pageRedis tools.Page[model.BookInfo]
	err = json.Unmarshal([]byte(pageRedisJson), &pageRedis)
	if err == nil { //查到有数据，返回缓存数据
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "查询到缓存",
			Data:    pageRedis,
		})
		return
	}
	//缓存没有，去查数据库
	books := model.SearchBook(ISBN, bookName)
	//查出所有数据后分页函数进行分页
	page := tools.Pages(books, currentPageString, pageSizeString)
	//将查出的数据存到redis
	pageJson, _ := json.Marshal(page)
	//失效时间
	loseTime := time.Duration(rand.Intn(100)) * time.Second
	model.InfoCacheRedisConn.Set(context, key, pageJson, loseTime)
	if page.Total == 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "没有此页数据",
			Data:    nil,
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "查询书籍成功",
		Data:    page,
	})
}

// AddBook godoc
//
// @Summary		新增图书
// @Description	管理员添加图书
// @Tags		admin/books
// @Accept		json
// @Produce		json
// @Param bookName body string true "书名"
// @Param author body string true "作者"
// @Param publishingHouse body string true "出版社"
// @Param translator body string true "译者"
// @Param publishDate body string true "发行时间"
// @Param pages body string true "页数"
// @Param ISBN body string true "ISBN号码"
// @Param price body string true "价格"
// @Param briefIntroduction body string true "内容简介"
// @Param authorIntroduction body string true "作者简介"
// @Param imgUrl body string true "封面地址"
// @Param delFlg body string true "删除标识"
// @Param count body int true "图书数量"
// @Param categoryId body int64 true "图书种类id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books [POST]
func AddBook(context *gin.Context) {
	var book model.BookInfo
	if err := context.ShouldBind(&book); err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.UserInfoError,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
		return

	}
	model.AddBook(book)
	context.JSON(http.StatusCreated, tools.Response{
		Code:    tools.OK,
		Message: "添加书籍成功",
		Data:    nil,
	})
}

// UpdateBook godoc
//
// @Summary		修改图书
// @Description	管理员修改图书
// @Tags		admin/books
// @Accept		json
// @Produce		json
// @Param id path int64 true "图书id"
// @Param bookName body string true "书名"
// @Param author body string true "作者"
// @Param publishingHouse body string true "出版社"
// @Param translator body string true "译者"
// @Param publishDate body string true "发行时间"
// @Param pages body string true "页数"
// @Param ISBN body string true "ISBN号码"
// @Param price body string true "价格"
// @Param briefIntroduction body string true "内容简介"
// @Param authorIntroduction body string true "作者简介"
// @Param imgUrl body string true "封面地址"
// @Param delFlg body string true "删除标识"
// @Param count body int true "图书数量"
// @Param categoryId body int64 true "图书种类id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books/{id} [PUT]
func UpdateBook(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	var updateBook model.BookInfo
	if err := context.ShouldBind(&updateBook); err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.NotAcceptable,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
		return
	}
	updateBook.Id = id
	if err := model.UpdateBook(updateBook); err != nil {
		context.JSON(http.StatusNotFound, tools.Response{
			Code:    tools.SourceNotFound,
			Message: "没有此书，修改失败",
		})
		return
	}
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
// @Tags		user/books
// @Produce		json
// @Param Authorization header string true "Bearer 用户令牌"
// @CookieParam id  string true "用户id"
// @Param bookId path int64 true "书籍id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/user/users/records/{bookId} [POST]
func BorrowBook(context *gin.Context) {
	userIdString, _ := context.Cookie("id")
	//查询用户3秒内是否请求过
	url := context.Request.URL.Path
	pathStr := fmt.Sprintf("%v%v", url, userIdString)
	requestQuery := model.StopRestartRequestConn
	//如果redis存过，则提醒休息重试
	redisPathStr, _ := requestQuery.Get(context, pathStr).Result()
	if pathStr == redisPathStr {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "请求太快，请休息一下重试~",
			Data:    nil,
		})
		return
	}
	//每次请求时，将请求url存到redis

	requestQuery.Set(context, pathStr, pathStr, time.Second*3)
	bookIdString := context.Param("bookId")
	fmt.Println("bookIdString:", bookIdString)
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	bookId, err := strconv.ParseInt(bookIdString, 10, 64)
	if err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.NotAcceptable,
			Message: "失败" + err.Error(),
		})
		return
	}
	record := model.Record{
		UserId:    userId,
		BookId:    bookId,
		Status:    0,
		StartTime: time.Now(),
		OverTime:  time.Now().Add(tools.T),
	}
	record.Id, err = model.CreateRecord(record)
	if record.Id > 0 && err == nil {
		context.JSON(http.StatusCreated, tools.Response{
			Code:    tools.OK,
			Message: "借书成功",
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "借书失败" + err.Error(),
	})
}

// ReturnBook godoc
//
// @Summary		用户还书
// @Description	用户还书
// @Tags		user/books
// @Accept		multipart/form-data
// @Produce		json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int64 true "借书记录的id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/user/users/records/{id} [PUT]
func ReturnBook(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	id, err := model.UpdateRecordAndBook(id)
	if id > 0 && err == nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "还书成功",
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "还书失败" + err.Error(),
	})
}

// DeleteBook godoc
//
// @Summary		管理员删除图书
// @Description	管理员删除图书
// @Tags		admin/books
// @Produce		json
// @Param id path int64 true "书籍id"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/books/{id} [DELETE]
func DeleteBook(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	err := model.DeleteBook(id)
	if err != nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "没有此书",
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "删除书籍成功",
	})
}
