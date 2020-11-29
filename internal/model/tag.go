package model

import (
	"blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (a Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
//实际的增删查改操作
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" { //如果该条记录的名字不为空就查询改记录
		db = db.Where("name= ?", t.Name) //根据条件返回记录
	}
	db = db.Where("state=?", t.State)
	//count获取记录 查询is_del = o 的记录 在t这个结构  上面先查名字和状态 再查is_del = 0的记录
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		//查询name为t.name
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State) //获取记录
	//查询is_del = 0 的所有记录返回到tags中
	if err = db.Where("is_del=?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error //根据结构体t创建表 并返回错误 t相当于一条实例
}
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	if err:=db.Model(t).Updates(values).Where("id= ?AND is_del=?",t.ID).Error;err!=nil{
		return err
	}
	return nil
}
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?",t.ID,0).Delete(&t).Error//删除本身id=t.id 和is_del=0的记录
}

//对模型操作的封装 对结构体t建表
