// 自动生成模板Book
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type Book struct {
      gorm.Model
      SrcWeb  string `json:"srcWeb" form:"srcWeb" gorm:"column:src_web;comment:'原网站';type:char(20)"`
      SrcHost  string `json:"srcHost" form:"srcHost" gorm:"column:src_host;comment:'原域名';type:char(20)"`
      SrcUrl  string `json:"srcUrl" form:"srcUrl" gorm:"column:src_url;comment:'原网址';type:char(20)"`
      BookTitle  string `json:"bookTitle" form:"bookTitle" gorm:"column:book_title;comment:'小说标题';type:char(20)"`
      BookId  int `json:"bookId" form:"bookId" gorm:"column:book_id;comment:'小说编号';type:int(10)"`
      BookWriter  string `json:"bookWriter" form:"bookWriter" gorm:"column:book_writer;comment:'作者';type:char(20)"`
      BookType  string `json:"bookType" form:"bookType" gorm:"column:book_type;comment:'小说类型';type:char(20)"`
      SectionNum  int `json:"sectionNum" form:"sectionNum" gorm:"column:section_num;comment:'章节编号';type:int(10)"`
      SectionTitle  string `json:"sectionTitle" form:"sectionTitle" gorm:"column:section_title;comment:'章节标题';type:char(20)"`
      SectionContent  string `json:"sectionContent" form:"sectionContent" gorm:"column:section_content;comment:'章节内容';type:varchar(255)"`
      Remarks  string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:'备注';type:varchar(50)"` 
}


func (Book) TableName() string {
  return "book"
}
