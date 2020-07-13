package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateBook
// @description   create a Book
// @param     book               model.Book
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateBook(book model.Book) (err error) {
	err = global.GVA_DB.Create(&book).Error
	return err
}

// @title    DeleteBook
// @description   delete a Book
// @auth                     （2020/04/05  20:22）
// @param     book               model.Book
// @return                    error

func DeleteBook(book model.Book) (err error) {
	err = global.GVA_DB.Delete(book).Error
	return err
}

// @title    DeleteBookByIds
// @description   delete Books
// @auth                     （2020/04/05  20:22）
// @param     book               model.Book
// @return                    error

func DeleteBookByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Book{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateBook
// @description   update a Book
// @param     book          *model.Book
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateBook(book *model.Book) (err error) {
	err = global.GVA_DB.Save(book).Error
	return err
}

// @title    GetBook
// @description   get the info of a Book
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Book        Book

func GetBook(id uint) (err error, book model.Book) {
	err = global.GVA_DB.Where("id = ?", id).First(&book).Error
	return
}

// @title    GetBookInfoList
// @description   get Book list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetBookInfoList(info request.BookSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Book{})
    var books []model.Book
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.SrcWeb != "" {
        db = db.Where("src_web LIKE ?","%"+ info.SrcWeb+"%")
    }
    if info.SrcHost != "" {
        db = db.Where("src_host LIKE ?","%"+ info.SrcHost+"%")
    }
    if info.SrcUrl != "" {
        db = db.Where("src_url LIKE ?","%"+ info.SrcUrl+"%")
    }
    if info.BookTitle != "" {
        db = db.Where("book_title LIKE ?","%"+ info.BookTitle+"%")
    }
    if info.BookId != 0 {
        db = db.Where("book_id = ?",info.BookId)
    }
    if info.BookWriter != "" {
        db = db.Where("book_writer LIKE ?","%"+ info.BookWriter+"%")
    }
    if info.BookType != "" {
        db = db.Where("book_type LIKE ?","%"+ info.BookType+"%")
    }
    if info.SectionNum != 0 {
        db = db.Where("section_num = ?",info.SectionNum)
    }
    if info.SectionTitle != "" {
        db = db.Where("section_title LIKE ?","%"+ info.SectionTitle+"%")
    }
    if info.SectionContent != "" {
        db = db.Where("section_content LIKE ?","%"+ info.SectionContent+"%")
    }
    if info.Remarks != "" {
        db = db.Where("remarks LIKE ?","%"+ info.Remarks+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&books).Error
	return err, books, total
}