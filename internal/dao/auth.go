package dao

import "blog-service/internal/model"

//在dao层获取到数据库对象和模型
func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth:=model.Auth{AppKey:appKey,AppSecret:appSecret}
	return auth.Get(d.engine)
}
