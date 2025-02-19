// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"bookstore/api/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/add",
				Handler: AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/update",
				Handler: UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/check",
				Handler: CheckHandler(serverCtx),
			},
		},
	)
}
