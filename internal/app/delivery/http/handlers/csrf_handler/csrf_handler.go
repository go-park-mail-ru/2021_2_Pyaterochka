package csrf_handler

import (
	"io"
	"net/http"
	"patreon/internal/app"
	usecase_csrf "patreon/internal/app/csrf/usecase"
	bh "patreon/internal/app/delivery/http/handlers/base_handler"
	"patreon/internal/app/delivery/http/handlers/handler_errors"
	models_respond "patreon/internal/app/delivery/http/models"
	"patreon/internal/app/sessions"
	"patreon/internal/app/sessions/middleware"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type CsrfHandler struct {
	csrfUsecase    usecase_csrf.Usecase
	sessionManager sessions.SessionsManager
	bh.BaseHandler
}

func NewCsrfHandler(log *logrus.Logger, router *mux.Router, cors *app.CorsConfig, sManager sessions.SessionsManager,
	uc usecase_csrf.Usecase) *CsrfHandler {
	h := &CsrfHandler{
		BaseHandler:    *bh.NewBaseHandler(log, router, cors),
		sessionManager: sManager,
		csrfUsecase:    uc,
	}
	h.AddMethod(http.MethodGet, h.GET)
	h.AddMiddleware(middleware.NewSessionMiddleware(sManager, log).Check)
	return h
}

// GET CSRF Token
// @Summary get CSRF Token
// @Description generate usecase token and return to client
// @Produce json
// @Success 200 {object} models_respond.TokenResponse
// @Failure 500 {object} models.ErrResponse "server error"
// @Router /token [GET]
func (h *CsrfHandler) GET(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.Log(r).Error(err)
		}
	}(r.Body)
	sessionId, ok := r.Context().Value("session_id").(string)
	if !ok {
		h.Log(r).Error("invalid conversation session_id from context to string")
		h.Error(w, r, http.StatusInternalServerError, handler_errors.InternalError)
		return
	}
	userId, ok := r.Context().Value("user_id").(int64)
	if !ok {
		h.Log(r).Error("invalid conversation userId from context to int64")
		h.Error(w, r, http.StatusInternalServerError, handler_errors.InternalError)
		return
	}
	token, err := h.csrfUsecase.Create(sessionId, userId)
	if err != nil {
		h.Log(r).Error("can not get user_id from context")
		h.UsecaseError(w, r, err, codeByErrors)
		return
	}
	h.Log(r).Debugf("get token %v", token)
	h.Respond(w, r, http.StatusOK, models_respond.TokenResponse{Token: token})
}