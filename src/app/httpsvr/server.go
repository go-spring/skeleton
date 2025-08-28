package httpsvr

import (
	"net/http"

	"GS_PROJECT_MODULE/idl/http/proto"

	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(func() *http.ServeMux {
		mux := http.NewServeMux()
		proto.InitRouter(mux, &GS_PROJECT_NAMEController{})
		return mux
	})
}
