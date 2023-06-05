package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"libraryManagementSystem/appV2/logic"
	_ "libraryManagementSystem/docs"
)

func New() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./appV2/resource")
	userRouter(r)
	adminRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/getCode", logic.GetCode)
	r.POST("/userLogin", logic.UserLogin)
	r.POST("/users", logic.AddUser)
	r.POST("/adminLogin", logic.LibrarianLogin)
	//游客可以浏览书籍和分类
	book := r.Group("/books")
	{
		book.GET("", logic.SearchBook)
		book.GET("/:id", logic.GetBook)
	}
	r.GET("/categories", logic.SearchCategory)
	return r
}
