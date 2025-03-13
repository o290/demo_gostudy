// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"demo_gostudy/zero/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: ApiHandler(serverCtx),
			},
		},
	)
}
