package app

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//做自定义的信息保存
type Claims struct {

	//自定义的验证信息  这对应的是payload字段
	AppKey             string `json:"app_key"`
	AppSecret          string `json:"app_secret"`
	jwt.StandardClaims        //对应的payload字段 中间字段 作为需要传输的字段
}

//获取密钥
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

//生成Token  这是给客户端的Token
func GenerateToken(appKey, appSecret string) (string, error) {
	// 1.Token包含过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	// 2.设置有效字段  包括自定义和官方建议设置的
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey), //这里对应的是条用md5算法对其进行加密
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //对自定义的有效字段结构体进行加密
	token, err := tokenClaims.SignedString(GetJWTSecret())           //加入密钥然后生成签名 生成标准的token
	return token, err
}

//解析Token  这是客户端接受然后解析 并验证
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return GetJWTSecret(), nil
	})
	//如果返回不为空
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims,nil
		}
	}
	//如果返回为空 就返回空
	return nil,err
}
