package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"patreon/internal/app"
	"patreon/internal/app/handlers/base_handler"
	"patreon/internal/app/handlers/handler_errors"
	"patreon/internal/app/sessions/middleware"
	"patreon/internal/models"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type CreatorCreateHandler struct {
	dataStorage app.DataStorage
	base_handler.BaseHandler
}

func NewCreatorCreateHandler(log *logrus.Logger, storage app.DataStorage) *CreatorCreateHandler {
	h := &CreatorCreateHandler{
		BaseHandler: *base_handler.NewBaseHandler(log),
		dataStorage: storage,
	}
	h.AddMethod(http.MethodGet, h.GET)
	h.AddMethod(http.MethodPost, h.POST)

	h.AddMiddleware(middleware.NewSessionMiddleware(h.dataStorage.SessionManager(), h.Log()).Check)
	return h
}

// Get Creator
// @Summary get creator
// @Description get creator with id from path
// @Produce json
// @Param id path int true "Get creator with id"
// @Success 200 {object} models.Creator "Get profile successfully"
// @Failure 503 {object} models.BaseResponse "Internal error"
// @Router /creators/{:id} [GET]
// Create Creator
// @Summary create creator
// @Description create creator with id from path, and respond created creator
// @Produce json
// @Param creator body models.RequestCreator true "Request body for create"
// @Success 200 {object} models.Creator "Create profile successfully"
// @Failure 400 {object} models.BaseResponse "Invalid request query"
// @Failure 404 {object} models.BaseResponse "User with id not found"
// @Failure 409 {object} models.BaseResponse "Creator already exist"
// @Failure 422 {object} models.BaseResponse "Invalid request body"
// @Failure 503 {object} models.BaseResponse "Internal error"
// @Router /creators/{:id} [POST]
func (h *CreatorCreateHandler) POST(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.Log().Error(err)
		}
	}(r.Body)

	vars := mux.Vars(r)
	id, ok := vars["id"]
	h.Log().Info("in /creators/id")
	idInt, err := strconv.Atoi(id)
	if len(vars) > 1 || !ok || err != nil {
		h.Log().Info(vars)
		h.Error(w, r, http.StatusBadRequest, handler_errors.InvalidParameters)
		return
	}
	req := &models.RequestCreator{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		h.Log().Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}
	u, err := h.dataStorage.Store().User().FindByID(int64(idInt))
	if err != nil {
		h.Log().Errorf("get: %s err:%s can not get user from db", u, err)
		h.Error(w, r, http.StatusNotFound, handler_errors.UserNotFound)
		return
	}
	if _, err := h.dataStorage.Store().Creator().GetCreator(int64(idInt)); err == nil {
		h.Log().Errorf("get: %s err:%s", u, handler_errors.ProfileAlreadyExist)
		h.Error(w, r, http.StatusConflict, handler_errors.ProfileAlreadyExist)
		return
	}
	cr := &models.Creator{
		ID:          u.ID,
		Nickname:    u.Nickname,
		Category:    req.Category,
		Description: req.Description,
	}
	if err := cr.Validate(); err != nil {
		toLog, _ := json.Marshal(err)
		h.Log().Errorf("get: %v err:%v ", cr, string(toLog))
		h.Error(w, r, http.StatusBadRequest, handler_errors.InvalidBody)
		return
	}
	if err := h.dataStorage.Store().Creator().Create(cr); err != nil {
		h.Log().Errorf("get: %v err:%v can not create profile", cr, err)
		h.Error(w, r, http.StatusServiceUnavailable, handler_errors.BDError)
		return
	}
	h.Respond(w, r, http.StatusOK, cr)
}

func (h *CreatorCreateHandler) GET(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.Log().Error(err)
		}
	}(r.Body)

	vars := mux.Vars(r)
	id, ok := vars["id"]
	h.Log().Info("in /creators/id")
	idInt, err := strconv.Atoi(id)
	if len(vars) > 1 || !ok || err != nil {
		h.Log().Info(vars)
		h.Error(w, r, http.StatusBadRequest, handler_errors.InvalidParameters)
		return
	}
	creator, err := h.dataStorage.Store().Creator().GetCreator(int64(idInt))
	if err != nil {
		h.Log().Errorf("get: %v err:%v can not get user from db", creator, err)
		h.Error(w, r, http.StatusServiceUnavailable, handler_errors.GetProfileFail)
		return
	}

	h.Log().Debugf("get creator %v with id %v", creator, id)
	h.Respond(w, r, http.StatusOK, creator)

}
