package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8
type Fields map[string]interface{}

//定义不同级别的日志
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "Debug"
	case LevelInfo:
		return "Info"
	case LevelWarn:
		return "Warn"
	case LevelError:
		return "Error"
	case LevelFatal:
		return "Fatal"
	case LevelPanic:
		return "Panic"
	}
	return ""
}

//日志实例初始化
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

//取地址实例化
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}
//克隆一个日志对象 在原本上进行复制 返回一个指针对象
func (l *Logger) clone() *Logger {
	//地址就是指针类型 返回地址用指针类型接住  指针指向对象
	nl := *l
	return &nl
}
//设置日志等级
func (l *Logger) WithLevel(lvl Level) *Logger {
	ll:=l.clone()
	ll.level = lvl
	return ll
}
//设置fields
func (l *Logger) WithFields(fields Fields) *Logger {
	ll:= l.clone()
	//克隆一个对象
	if ll.fields == nil {
		//上面定义了 现在创建
		ll.fields = make(Fields)
	}
	//range返回键值对
	for k,v:=range fields{
		ll.fields[k] = v
	}
	return ll
}

func (l *Logger)WithContext(ctx context.Context)*Logger  {
	ll := l.clone()
	ll.ctx= ctx
	return ll
}
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc,file,line,ok:=runtime.Caller(skip)
	if ok {
		f:=runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s:%d%s",file,line,f.Name())}
	}
	return ll
}
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return l
}

func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprintf(format, v...))
}