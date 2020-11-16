package dao

//dao层	与数据库访问的操作  internal 内部模块
import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

//处理标签模块的dao操作
//对确定的标签页进行操作
//封装对标签模块的dao操作 数据库操作
func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}
func (d *Dao) GetTag(id uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}
//获取标签列表  接受者是谁 就是谁做事
func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}
func (d *Dao) CreateTag(naem string, state uint8, creareBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			CreatedBy: "",
		},
		Name:  naem,
		State: state,
	}
	return tag.Create(d.engine)
}
func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}