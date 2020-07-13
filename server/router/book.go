package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBookRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("book").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		BookRouter.POST("createBook", v1.CreateBook)   // 新建Book
		BookRouter.DELETE("deleteBook", v1.DeleteBook) // 删除Book
		BookRouter.DELETE("deleteBookByIds", v1.DeleteBookByIds) // 批量删除Book
		BookRouter.PUT("updateBook", v1.UpdateBook)    // 更新Book
		BookRouter.GET("findBook", v1.FindBook)        // 根据ID获取Book
		BookRouter.GET("getBookList", v1.GetBookList)  // 获取Book列表
	}
}
