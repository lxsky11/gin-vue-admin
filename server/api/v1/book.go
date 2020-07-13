package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Book
// @Summary 创建Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Book true "创建Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /book/createBook [post]
func CreateBook(c *gin.Context) {
	var book model.Book
	_ = c.ShouldBindJSON(&book)
	err := service.CreateBook(book)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Book
// @Summary 删除Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Book true "删除Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /book/deleteBook [delete]
func DeleteBook(c *gin.Context) {
	var book model.Book
	_ = c.ShouldBindJSON(&book)
	err := service.DeleteBook(book)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Book
// @Summary 批量删除Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /book/deleteBookByIds [delete]
func DeleteBookByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteBookByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Book
// @Summary 更新Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Book true "更新Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /book/updateBook [put]
func UpdateBook(c *gin.Context) {
	var book model.Book
	_ = c.ShouldBindJSON(&book)
	err := service.UpdateBook(&book)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Book
// @Summary 用id查询Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Book true "用id查询Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /book/findBook [get]
func FindBook(c *gin.Context) {
	var book model.Book
	_ = c.ShouldBindQuery(&book)
	err, rebook := service.GetBook(book.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rebook": rebook}, c)
	}
}

// @Tags Book
// @Summary 分页获取Book列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BookSearch true "分页获取Book列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /book/getBookList [get]
func GetBookList(c *gin.Context) {
	var pageInfo request.BookSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetBookInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
