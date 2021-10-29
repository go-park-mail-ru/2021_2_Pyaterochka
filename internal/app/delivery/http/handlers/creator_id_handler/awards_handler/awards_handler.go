package awards_handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"image/color"
	"io"
	"net/http"
	"patreon/internal/app"
	bh "patreon/internal/app/delivery/http/handlers/base_handler"
	"patreon/internal/app/delivery/http/handlers/handler_errors"
	"patreon/internal/app/delivery/http/models"
	"patreon/internal/app/middleware"
	db_models "patreon/internal/app/models"
	"patreon/internal/app/sessions"
	sessionMid "patreon/internal/app/sessions/middleware"

	useAwards "patreon/internal/app/usecase/awards"
)

type AwardsHandler struct {
	awardsUsecase useAwards.Usecase
	bh.BaseHandler
}

func NewAwardsHandler(log *logrus.Logger, router *mux.Router, cors *app.CorsConfig,
	ucAwards useAwards.Usecase, manager sessions.SessionsManager) *AwardsHandler {
	h := &AwardsHandler{
		BaseHandler:   *bh.NewBaseHandler(log, router, cors),
		awardsUsecase: ucAwards,
	}
	h.AddMethod(http.MethodGet, h.GET)
	h.AddMethod(http.MethodPost, h.POST, sessionMid.NewSessionMiddleware(manager, log).CheckFunc,
		middleware.NewCreatorsMiddleware(log).CheckAllowUserFunc)
	return h
}

// GET Awards
// @Summary get list of awards of some creator
// @Description get list of awards which belongs the creator
// @Produce json
// @Success 201 {array} models.ResponseCreator
// @Failure 500 {object} models.ErrResponse "can not do bd operation"
// @Failure 400 {object} models.ErrResponse "invalid parameters"
// @Router /creators/{:creator_id}/awards [GET]
func (h *AwardsHandler) GET(w http.ResponseWriter, r *http.Request) {
	idInt, ok := h.GetInt64FromParam(w, r, "creator_id")
	if !ok {
		return
	}

	if len(mux.Vars(r)) > 1 {
		h.Log(r).Warnf("Too many parametres %v", mux.Vars(r))
		h.Error(w, r, http.StatusBadRequest, handler_errors.InvalidParameters)
		return
	}

	awards, err := h.awardsUsecase.GetAwards(idInt)
	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	respondAwards := make([]models.ResponseAwards, len(awards))
	for i, aw := range awards {
		respondAwards[i] = models.ToResponseAwards(aw)
	}

	h.Log(r).Debugf("get creators %v", respondAwards)
	h.Respond(w, r, http.StatusOK, respondAwards)
}

// POST Create Awards
// @Summary create awards
// @Description create awards to creator with id from path
// @Param user body models.RequestAwards true "Request body for awards"
// @Produce json
// @Success 201 {object} models.AwardsResponse "id awards"
// @Failure 422 {object} models.ErrResponse "invalid body in request"
// @Failure 400 {object} models.ErrResponse "invalid parameters"
// @Failure 400 {object} models.ErrResponse "empty name in request"
// @Failure 400 {object} models.ErrResponse "incorrect value of price"
// @Failure 500 {object} models.ErrResponse "can not do bd operation"
// @Failure 500 {object} models.ErrResponse "server error"
// @Failure 500 {object} models.ErrResponse "can not get info from context"
// @Failure 403 {object} models.ErrResponse "for this user forbidden change creator"
// @Failure 404 "User are not authorized"
// @Router /creators/{:creator_id}/awards [POST]
func (h *AwardsHandler) POST(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.Log(r).Error(err)
		}
	}(r.Body)

	idInt, ok := h.GetInt64FromParam(w, r, "creator_id")
	if !ok {
		return
	}
	if len(mux.Vars(r)) > 1 {
		h.Log(r).Warnf("Too many parametres %v", mux.Vars(r))
		h.Error(w, r, http.StatusBadRequest, handler_errors.InvalidParameters)
		return
	}

	req := &models.RequestAwards{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		h.Log(r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	aw := &db_models.Awards{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Color:       color.RGBA{R: req.Color.R, G: req.Color.G, B: req.Color.B, A: req.Color.A},
		CreatorId:   idInt,
	}

	awardsId, err := h.awardsUsecase.Create(aw)
	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	h.Respond(w, r, http.StatusCreated, &models.AwardsResponse{ID: awardsId})
}