package http

import (
	userServ "github.com/lefties-startup/companion-api/internal/service/user"
	"go.uber.org/zap"
	"net/http"
)

func RegisterUserRoutes(
	userSvc userServ.ServiceUser,
	mux *http.ServeMux,
	logger *zap.Logger,
) *http.ServeMux {
	handler := NewUserHandler(userSvc, logger)

	// Роуты
	mux.HandleFunc("GET /api/v1/users/info", handler.GetInfo)
	///////////////////////////////////////
	return mux
}
