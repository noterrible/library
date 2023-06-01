package router

import (
	"github.com/gin-gonic/gin"
	"libraryManagementSystem/appv1/logic"
	"libraryManagementSystem/appv1/model"
)

func adminRouter(r *gin.Engine) {
	//librarian := r.Group("/librarians").Use(librarianCheck())
	base := r.Group("/admin")
	base.Use(librarianCheck())
	user := base.Group("/users")
	{
		user.GET("/:id", logic.GetUserById)
		user.GET("", logic.SearchUser)
		user.PUT("/:id", logic.UpdateUserByAdmin)
		user.DELETE("/:id", logic.DeleteUser)
		//获取用户已归还或者未归还的所有记录
		user.GET("/:id/records/:status", logic.GetUserStatusRecords)
		//user.POST("/:id/records/:bookId", logic.BorrowBook)
		//user.PUT("/:id/records/:bookId", logic.ReturnBook)
	}
	//书的所有资源
	book := base.Group("/books")
	{
		//book.GET("/:id", logic.GetBook)
		//book.GET("", logic.SearchBook)
		book.POST("", logic.AddBook)
		book.PUT("/:id", logic.UpdateBook)
		book.DELETE("/:id", logic.DeleteBook)
	}
	category := base.Group("/categories")
	{
		category.GET("/:id", logic.GetCategory)
		//category.GET("", logic.SearchCategory)
		category.POST("", logic.AddCategory)
		category.PUT("/:id", logic.UpdateCategory)
		category.DELETE("/:id", logic.DeleteCategory)
	}
	//记录表的资源
	record := base.Group("/records")
	{
		//所有借书还书记录
		//record.GET("", logic.GetRecords)
		//所有归还或者未归还的记录
		record.GET("/:status", logic.GetStatusRecords)
	}
}
func librarianCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := model.GetSession(context)
		idInter := session["id"]
		nameInter := session["name"]
		id := idInter.(int64)
		name := nameInter.(string)
		if model.AdminCheck(id, name) {
			context.Next()
			return
		}
		context.Abort()
		return
	}
}
