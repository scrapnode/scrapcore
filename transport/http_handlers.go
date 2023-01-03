package transport

import (
	"context"
	"github.com/scrapnode/scrapcore/xconfig"
	"net/http"
	"time"
)

func NewHttpPing(ctx context.Context, cfg *xconfig.Configs) *HttpHandler {
	return &HttpHandler{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(writer http.ResponseWriter, request *http.Request) {
			data := map[string]interface{}{
				"version":      cfg.Version,
				"environement": cfg.Env,
				"ts":           time.Now().UnixMilli(),
			}
			WriteJSON(writer, data)
		},
	}
}
