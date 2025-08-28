package httpsvr

import (
	"context"

	"GS_PROJECT_MODULE/idl/http/proto"
	order "GS_PROJECT_MODULE/src/internal/order/controller"
	user "GS_PROJECT_MODULE/src/internal/user/controller"
)

type GS_PROJECT_NAMEController struct {
	order.OrderController
	user.UserController
}

func (c *GS_PROJECT_NAMEController) Ping(ctx context.Context, req *proto.PingReq) *proto.PingResp {
	return proto.NewPingResp()
}
