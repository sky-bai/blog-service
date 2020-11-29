package middleware

import (
	"blog-service/global"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

//中间件将主体框架与第三方框架相结合
func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {



		var ctx context.Context
		//1.找到span所要包含的范围
		span := opentracing.SpanFromContext(c.Request.Context()) //返回的是一个与传入参数ctx相关联的span    根据传入的context获取与它相关联的span
		if span != nil {     //如果传入的ctx有span
			span, ctx = //返回的是带有operation name的span 第三个参数就是name 这里如果有span我们就创建一个子节点
				opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
					global.Tracer, c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {    //如果之前context没有与之相连的span
			span, ctx = // 这里将根据传入的请求context创建一个rootspan operation name是请求的路径
				opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
					global.Tracer, c.Request.URL.Path)
		}
		defer span.Finish()

		var spanContext= span.Context()
		var traceID string
		var SpanID string

		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceID =
				spanContext.(jaeger.SpanContext).TraceID().String()
			SpanID =
				spanContext.(jaeger.SpanContext).SpanID().String()
		}
		//将获取到的TraceID和SpanID存在context中
		c.Set("X-Trace-ID",traceID)
		c.Set("X-Trace-ID",SpanID)

		c.Request = c.Request.WithContext(ctx)//将请求的context更新为传入的ctx   这里是浅拷贝 复制的是指针 深拷贝是创建一个新对象
		c.Next()
	}
}
