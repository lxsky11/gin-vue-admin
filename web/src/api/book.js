import service from '@/utils/request'

// @Tags Book
// @Summary 创建Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Book true "创建Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /book/createBook [post]
export const createBook = (data) => {
     return service({
         url: "/book/createBook",
         method: 'post',
         data
     })
 }


// @Tags Book
// @Summary 删除Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Book true "删除Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /book/deleteBook [delete]
 export const deleteBook = (data) => {
     return service({
         url: "/book/deleteBook",
         method: 'delete',
         data
     })
 }

// @Tags Book
// @Summary 删除Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /book/deleteBook [delete]
 export const deleteBookByIds = (data) => {
     return service({
         url: "/book/deleteBookByIds",
         method: 'delete',
         data
     })
 }

// @Tags Book
// @Summary 更新Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Book true "更新Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /book/updateBook [put]
 export const updateBook = (data) => {
     return service({
         url: "/book/updateBook",
         method: 'put',
         data
     })
 }


// @Tags Book
// @Summary 用id查询Book
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Book true "用id查询Book"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /book/findBook [get]
 export const findBook = (params) => {
     return service({
         url: "/book/findBook",
         method: 'get',
         params
     })
 }


// @Tags Book
// @Summary 分页获取Book列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Book列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /book/getBookList [get]
 export const getBookList = (params) => {
     return service({
         url: "/book/getBookList",
         method: 'get',
         params
     })
 }