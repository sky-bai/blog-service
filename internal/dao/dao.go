package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB //
}

//为实体属性赋值
func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

