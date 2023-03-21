package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Article struct {
	Model
	// 外键, gorm会通过类名+ID的方式去找到这两个类之间的关联关系
	// gorm:index用于声明这个字段为索引, 如果你使用了自动迁移功能则会有所影响, 在不使用则无影响
	TagID int `json:"tag_id" gorm:"index"`
	// Tag字段, 实际是一个嵌套的struct, 它利用TagID与Tag模型相互关联, 在执行查询的时候, 能够达到Article、Tag关联查询的功能
	// 可以通过Related进行关联查询
	Tag Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	// time.Now().Unix()返回当前的时间戳
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	return article.ID > 0
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// 预加载器, 它会执行两条SQL, 分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);在查询出结构后, gorm内部处理对应的映射逻辑, 将其填充到Article的Tag中, 会特别方便, 并且避免了循环查询
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		// (int)是强制类型转换, 将接收到的interface{}类型数据转换为int类型, 以便在创建Article结构体时正确传入参数
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	// Gorm在操作时会自动将传入的参数转换为指针类型, 因此我们可以不传递指针类型的参数
	// 传&Article{}可能更好, 因为这样可以提供更多的信息给Delete函数, 例如字段的默认值等等
	db.Where("id = ?", id).Delete(Article{})

	return true
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
