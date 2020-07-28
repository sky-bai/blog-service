package setting

import "github.com/spf13/viper"
//对读取文件配置的行为进行封装
type Setting struct {
	vp *viper.Viper
}

//获取配置文件
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}
