package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appV2/logic"
	"libraryManagementSystem/appV2/tools"
	"net/http"
)

func userRouter(r *gin.Engine) {
	base := r.Group("/user")
	base.Use(userCheck())
	user := base.Group("/users")
	{
		user.GET("", logic.GetUser)
		user.PUT("", logic.UpdateUser)
		//user.DELETE("/:id", logic.DeleteUser)
		user.GET("/records", logic.GetUserRecords)
		//user.GET("/records/:status", logic.GetUserStatusRecords)
		//用户自助借书还书
		user.POST("/records/:bookId", logic.BorrowBook)
		user.PUT("/records/:id", logic.ReturnBook)
		user.GET("/messages", logic.GetMessage)

	}
	//book := base.Group("/books")
	//{
	//	//book.GET("/:id", logic.GetBook)
	//	//book.POST("/:id", logic.AddBook)
	//	//book.DELETE("/:id", logic.DeleteBook)
	//}
	//category := base.Group("/categories")
	//{
	//	//category.GET("/:id", logic.GetCategory)
	//	category.GET("/:id/books", logic.GetCategoryBooks)
	//}
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIxMjM0LWxpYnJhcnkiLCJleHAiOjE2ODU3ODU5OTEsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiMSJ9.XLEHaAWVRRsKb4MlsVTUTcW_tqtMD2kRvGKAUDRw-IU
func userCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		//拦截成功Chec
		auth := context.GetHeader("Authorization")
		fmt.Println("auth:", auth)
		data, err := tools.Token.VerifyToken(auth)
		if err != nil {
			fmt.Println("验签失败," + err.Error())
			context.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
				Code:    tools.NoLogin,
				Message: "验签失败" + err.Error(),
				Data:    nil,
			})
			//去掉这个return没有token的会报500
			return
		}
		fmt.Printf("data:%+v\n", data)
		if data.ID <= 0 || data.Name == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
				Code:    tools.NoLogin,
				Message: "用户信息错误",
				Data:    nil,
			})
			return
		}
		context.Set("id", data.ID)
	}
}
