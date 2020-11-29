package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}
//创建一个新的限流器   对于某个方法的限流器
func NewMethodLimiter() LimiterIface {
	return MethodLimiter{
		&Limiter{
			limiterBuckets:
			make(map[string]*ratelimit.Bucket)},
	}
}

//在Key方法中根据RequestURI切割出核心路由作为键值对名称  获取到键值对名称
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?") //index 索引  返回substr之前第一个实例的索引
	if index == -1 {
		return uri
	}
	return uri[:index]
}
//根据键值对获取令牌桶
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}
//传入令牌桶的内容获得多个令牌桶
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _,rule:=range rules{
		if _, ok := l.limiterBuckets[rule.Key];!ok {
			l.limiterBuckets[rule.Key]=
				ratelimit.NewBucketWithQuantum(rule.FillInterval,rule.Capacity,rule.Quantum)
		}
	}
	return l
}
