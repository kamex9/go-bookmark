package restcontroller

import (
	"encoding/json"
	"go-bookmark/app/web/dto"
	"go-bookmark/core/usecase"
	"go-bookmark/core/utils"
	"io"
	"log/slog"
	"net/http"
)

type BookmarkRestController struct {
	l  *slog.Logger
	uc usecase.CrudBookmarkUseCase
}

func NewBookmarkRestController(l *slog.Logger, uc usecase.CrudBookmarkUseCase) *BookmarkRestController {
	return &BookmarkRestController{
		l:  l,
		uc: uc,
	}
}

func (c *BookmarkRestController) Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	c.l.Info("Params info", "body", string(reqBody))

	var req dto.BookmarkCreateRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		utils.WriteResponse(w, nil, dto.NewError("E0001", err))
		return
	}

	bm, _ := c.uc.Create(&req)
	utils.WriteResponse(w, bm, nil)
}

func (c *BookmarkRestController) FetchAll(w http.ResponseWriter, r *http.Request) {
	bms, err := c.uc.FindAll()
	if err != nil {
		utils.WriteResponse(w, nil, dto.NewError("E1001", err))
		return
	}
	res := make([]*dto.BookmarkResponse, 0, len(bms))
	for _, bm := range bms {
		res = append(res, dto.NewBookmarkResponse(bm))
	}
	utils.WriteResponse(w, res, nil)
}

func (c *BookmarkRestController) FetchById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.RetrievePathParamValue("id", r)
	if err != nil {
		utils.WriteResponse(w, nil, dto.NewError("E2001", err))
		return
	}
	bm, err := c.uc.FindById(id)
	if err != nil {
		utils.WriteResponse(w, nil, dto.NewError("E2002", err))
		return
	}
	utils.WriteResponse(w, dto.NewBookmarkResponse(bm), nil)
}

func (c *BookmarkRestController) DeleteAll(w http.ResponseWriter, r *http.Request) {
	err := c.uc.DeleteAll()
	if err != nil {
		utils.WriteResponse(w, nil, dto.NewError("E2002", err))
		return
	}
	utils.WriteResponse(w, nil, nil)
}

// func fetchBookmarkById(w http.ResponseWriter, r *http.Request) {
// 	id, err := utils.RetrievePathParamValue("id", r)
// 	if err != nil {
// 		writeResponse(w, nil, models.NewError("E2001", err))
// 		return
// 	}
// 	bm, err := repo.FindById(id)
// 	if err != nil {
// 		writeResponse(w, nil, models.NewError("E2002", err))
// 		return
// 	}
// 	writeResponse(w, models.NewBookmarkResponse(bm), nil)
// }
