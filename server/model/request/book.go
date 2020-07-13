package request

import "gin-vue-admin/model"

type BookSearch struct{
    model.Book
    PageInfo
}