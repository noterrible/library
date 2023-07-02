package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
// @Accept		json
// @Produce		json
// @Param currentPage  query string true "当前页"
// @Param pageSize  query string true "页大小"
// @Param ISBN  query string true "书籍编号"
// @Param bookName  query string true "书籍名称"
// @Success 200 {object} tools.Response{data=model.ListResponse[model.BookInfo]{}}
// @Router			/books [GET]
func SearchBook(context *gin.Context) {
	ISBN := context.Query("ISBN")
	bookName := context.Query("bookName")
	currentPageString := context.Query("currentPage")
	pageSizeString := context.Query("pageSize")
	currentPage, _ := strconv.Atoi(currentPageString)
	pageSize, _ := strconv.Atoi(pageSizeString)
	offset := pageSize * (currentPage - 1)
	books, count := model.SearchBook(ISBN, bookName, currentPage, offset)
	page := model.ListResponse[model.BookInfo]{
		Count: count,
		List:  books,
	}
	if books == nil {
		context.JSON(http.StatusInternalServerError, tools.Response{
			Code:    tools.OK,
			Message: "查询错误",
		})
		return
	}
	if page.Count == 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "搜索的书籍不存在",
		})
		return
	}
	if len(page.List) == 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "没有此页书籍",
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

// PageBook godoc
//
// @Summary		分页图书
// @Description	获取所有图书并进行分页,可以向前或者向后进行翻页
// @Tags		public
// @Accept		json
// @Produce		json
// @Param turnPageInfo  query model.TurnPageInfo true "分页信息"
// @Success 200 {object} tools.Response{data=model.TurnPageInfo[model.BookInfo]}
// @Router			/books/page [GET]
func PageBook(context *gin.Context) {
	var pageInfo model.TurnPageInfo
	if err := context.ShouldBind(&pageInfo); err != nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "参数错误",
		})
		return
	}
	//查缓存
	key := fmt.Sprintf("books_%v_%v_%v", pageInfo.Page, pageInfo.Limit, pageInfo.Sort)
	pageRedisGzip, err := model.InfoCacheRedisConn.Get(context, key).Result()
	var pageRedis model.ListResponse[model.BookInfo]
	if err != redis.Nil {
		//解压
		if errGT := tools.GzipTo(&pageRedis, pageRedisGzip); errGT != nil {
			context.JSON(http.StatusInternalServerError, tools.Response{
				Code:    tools.OK,
				Message: "解压错误" + err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "查询到缓存",
			Data:    pageRedis,
		})
		return
	}
	//缓存没有，去查数据库
	//游标分页获取数据
	page, err := model.TurnPages(model.BookInfo{}, pageInfo)
	if err != nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "查询错误" + err.Error(),
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
	/*查的数据存入redis*/
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	//失效时间 将查出的数据存到redis
	randInt := rand.Intn(43) + 20
	loseTime := time.Duration(randInt) * time.Second
	//压缩
	var pageGzipData string
	var errTG error
	if pageGzipData, errTG = tools.ToGzip(page); errTG != nil {
		context.JSON(http.StatusInternalServerError, tools.Response{
			Code:    tools.OK,
			Message: "压缩错误" + err.Error(),
		})
		return
	}
	model.InfoCacheRedisConn.Set(context, key, pageGzipData, loseTime)
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
// @Param book body model.BookInfo true "书籍"
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
	if err := model.AddBook(book); err != nil {
		context.JSON(http.StatusInternalServerError, tools.Response{
			Code:    tools.UserInfoError,
			Message: "添加书籍失败:" + err.Error(),
			Data:    nil,
		})
		return
	}
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

//限流
//登陆方式-采用存id限制用户
//不登陆-IP+UA
//WAF-宝塔、阿里云防火墙
//前端弹出验证码、短信验证码
//如果qps100->1000->100000怎么优化
//把限流单独抽象出一个中间件

//gzip压缩算法如何实现的

//redis大key怎么解决

//缓存问题：缓存一致性、多级缓存、击穿穿透雪崩
//多级缓存：Api缓存 logic、
//localcache本地缓存：go语言goche

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
