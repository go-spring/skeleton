package httpsvr

import (
	"fmt"
	"net/http"

	"GS_PROJECT_MODULE/idl/http/proto"

	"github.com/go-spring/spring-core/gs"
)

// Middleware 定义中间件的原型
type Middleware func(next http.Handler) http.Handler

// Middlewares 定义中间件列表
type Middlewares []Middleware

// Chain 创建一个中间件链
func Chain(h http.Handler, mws Middlewares) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func init() {
	gs.Object(&ServerConfig{})
	gs.Provide(func(config *ServerConfig, server *GS_PROJECT_NAMEController) http.Handler {
		mux := http.NewServeMux()
		proto.InitRouter(mux, server)
		mws := Middlewares{
			Recovery(config.RecoveryConfig),
			Trace(config.TraceConfig),
			Metric(config.MetricConfig),
		}
		return Chain(mux, mws)
	})
}

// ServerConfig 服务配置
type ServerConfig struct {
	RecoveryConfig RecoveryConfig `value:"${recovery}"`
	TraceConfig    TraceConfig    `value:"${trace}"`
	MetricConfig   MetricConfig   `value:"${metric}"`
}

// RecoveryConfig 崩溃恢复配置
type RecoveryConfig struct {
	Msg string `value:"${msg:=recovery}"`
}

// TraceConfig 链路追踪配置
type TraceConfig struct {
	Msg string `value:"${msg:=trace}"`
}

// MetricConfig 监控指标配置
type MetricConfig struct {
	Msg string `value:"${msg:=metric}"`
}

// Recovery 崩溃恢复中间件
func Recovery(config RecoveryConfig) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(config.Msg)
			defer func() {
				if err := recover(); err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// Trace 链路追踪中间件
func Trace(config TraceConfig) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(config.Msg)
			next.ServeHTTP(w, r)
		})
	}
}

// Metric 监控指标中间件
func Metric(config MetricConfig) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(config.Msg)
			next.ServeHTTP(w, r)
		})
	}
}
