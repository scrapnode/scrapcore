package transport

import (
	"context"
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/scrapnode/scrapcore/xlogger"
	"net/http"
)

func NewHttpPing(ctx context.Context, cfg *xconfig.Configs) *Handler {
	logger := xlogger.FromContext(ctx)
	return &Handler{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(writer http.ResponseWriter, request *http.Request) {
			data := map[string]interface{}{
				"version":      cfg.Version,
				"environement": cfg.Env,
			}
			if err := WriteJSON(writer, data); err != nil {
				logger.Errorw("could not send json data to client", "error", err.Error())
			}
		},
	}
}
