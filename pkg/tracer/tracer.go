package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

//第三方包实现具体功能
//根据服务名称和用户主机端口设置跟踪器的配置信息
//1.创建Tracer
func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	//设置配置
	cfg := &config.Configuration{
		ServiceName: serviceName, //ServiceName specifies the service name to use on the tracer.
		Sampler: &config.SamplerConfig{
			Type:  "const", //指定类型
			Param: 1,       //const中1是true
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	//根据上面的配置信息创建一个新的追踪器
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	//没有错的话就注册到全局跟踪器中
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, err
}
