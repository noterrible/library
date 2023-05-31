package router

import (
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appv1/logic"
)

func userRouter(r *gin.Engine) {
	base := r.Group("/user")
	base.Use(userCheck())
	user := base.Group("/users")
	{
		user.GET("", logic.GetUser)
		user.PUT("/:id", logic.UpdateUser)
		//user.DELETE(":id", logic.DeleteUser)
		user.GET("/:id/records", logic.GetUserRecords)
		user.GET("/:id/records/:status", logic.GetUserStatusRecords)
		//用户自助借书还书
		user.POST("/records/:bookId", logic.BorrowBook)
		user.PUT("/records/:id", logic.ReturnBook)
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
func userCheck() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
