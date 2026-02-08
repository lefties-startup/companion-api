package http

import (
	"encoding/json"
	"fmt"
	"github.com/lefties-startup/companion-api/internal/service"
	userServ "github.com/lefties-startup/companion-api/internal/service/user"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service userServ.ServiceUser
	logger  *zap.Logger
}

func NewUserHandler(svc userServ.ServiceUser, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: svc,
		logger:  logger.With(zap.String("handler", "user")),
	}
}

func (h *UserHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	tgIDStr := r.URL.Query().Get("tg_user_id")
	fmt.Println("test")
	if tgIDStr == "" {
		http.Error(w, `{"error": "missing tg_user_id"}`, http.StatusBadRequest)
		return
	}

	tgID, err := strconv.ParseInt(tgIDStr, 10, 64)
	if err != nil {
		http.Error(w, `{"error": "invalid tg_user_id"}`, http.StatusBadRequest)
		return
	}

	user, err := h.service.GetInfo(r.Context(), tgID)
	if err != nil {
		if errors.Is(err, service.NotFoundErr) {
			http.Error(w, `{"error": "user not found"}`, http.StatusNotFound)
			return
		}
		h.logger.Error("failed to get user", zap.Error(err), zap.Int64("tg_user_id", tgID))
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		h.logger.Error("failed to encode response", zap.Error(err))
	}
}
