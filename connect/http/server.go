package http

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type OpsServer struct {
	server *http.Server
	logger *zap.Logger
}

// NewOpsServer constructs a new OpsServer instance.
func NewOpsServer(addr string, logger *zap.Logger) *OpsServer {
	mux := http.NewServeMux()

	return &OpsServer{
		server: &http.Server{
			Addr:              addr,
			Handler:           mux,
			ReadHeaderTimeout: 10 * time.Second,
		},
		logger: logger.With(zap.String("component", "ops-server")),
	}
}

func (o *OpsServer) Register(pattern string, handler http.HandlerFunc) {
	o.server.Handler.(*http.ServeMux).HandleFunc(pattern, handler)
}

// Run поднимает http сервер.
func (o *OpsServer) Run(ctx context.Context) error {
	o.logger.Info("operation server is running", zap.String("address", o.server.Addr))
	errCh := make(chan error, 1)
	go func() {
		err := o.server.ListenAndServe()
		if err != nil {
			errCh <- err
		}
		close(errCh)
	}()
	select {
	case <-ctx.Done():
		o.logger.Info("operation server context cancelled, shutting down")
		return o.shutdown()
	case err := <-errCh:
		return err
	}
}

// Shutdown останавливает сервер.
func (o *OpsServer) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := o.server.Shutdown(ctx); err != nil {
		o.logger.Error("operation server forced to shutdown", zap.Error(err))
		return o.server.Close()
	}

	o.logger.Info("operation server stopped")
	return nil
}
