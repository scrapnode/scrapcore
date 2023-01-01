package transport

import (
	"context"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/scrapnode/scrapcore/xlogger"
	"go.uber.org/zap"
	"net/http"
)

func NewHttp(ctx context.Context, cfg *Configs, handlers []*Handler) (Transport, error) {
	router := httprouter.New()
	for _, handler := range handlers {
		router.HandlerFunc(handler.Method, handler.Path, handler.Handler)
	}

	transport := &Http{
		logger: xlogger.FromContext(ctx).With("pkg", "transport.http"),
		server: &http.Server{
			Addr:    cfg.ListenAddress,
			Handler: router,
		},
	}

	return transport, nil
}

type Handler struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Http struct {
	logger *zap.SugaredLogger
	server *http.Server
}

func (transport *Http) Start(ctx context.Context) error {
	transport.logger.Debug("starting")
	return nil
}

func (transport *Http) Stop(ctx context.Context) error {
	transport.logger.Debug("stopping")

	if transport.server == nil {
		return nil
	}

	return transport.server.Shutdown(ctx)
}

func (transport *Http) Run(ctx context.Context) error {
	if transport.server == nil {
		return errors.New("transport.http: no server was configured")
	}

	transport.logger.Debugw("running", "listen_address", transport.server.Addr)

	if err := transport.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
