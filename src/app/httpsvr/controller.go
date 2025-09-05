package httpsvr

import (
	"context"

	"GS_PROJECT_MODULE/idl/http/proto"
	order "GS_PROJECT_MODULE/src/internal/order/controller"
	user "GS_PROJECT_MODULE/src/internal/user/controller"

	"github.com/go-spring/spring-core/gs"
	redigo "github.com/gomodule/redigo/redis"
	goRedis "github.com/redis/go-redis/v9"
)

func init() {
	gs.Object(&GS_PROJECT_NAMEController{})
}

type GS_PROJECT_NAMEController struct {
	order.OrderController
	user.UserController

	RedigoPool    *redigo.Pool    `autowire:""`
	GoRedisClient *goRedis.Client `autowire:""`
}

func (c *GS_PROJECT_NAMEController) Ping(ctx context.Context, req *proto.PingReq) *proto.PingResp {
	req.Name = "Go-Spring"
	c.GoRedisClient.Set(ctx, "ping", req.Name, 0)
	p := c.RedigoPool.Get()
	reply, err := p.Do("GET", "ping")
	if err != nil {
		panic(err)
	}
	data, err := redigo.String(reply, err)
	if err != nil {
		panic(err)
	}
	p.Close()
	return &proto.PingResp{
		Errno:  0,
		Errmsg: "",
		Data:   data,
	}
}
