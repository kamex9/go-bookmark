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
	uc usecase.CreateBookmarkUseCase
}

func NewBookmarkRestController(l *slog.Logger, uc usecase.CreateBookmarkUseCase) *BookmarkRestController {
	return &BookmarkRestController{
		l:  l,
		uc: uc,
	}
}

func (c *BookmarkRestController) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	c.l.Info("Params info", "body", string(reqBody))

	var req dto.BookmarkCreateRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		utils.WriteResponse(w, nil, dto.NewError("E0001", err))
		return
	}

	bm, _ := c.uc.CreateBookmark(&req)
	utils.WriteResponse(w, bm, nil)
}

// func fetchAllBookmark(w http.ResponseWriter, r *http.Request) {
// 	bms, err := repo.FindAll()
// 	if err != nil {
// 		writeResponse(w, nil, models.NewError("E1001", err))
// 		return
// 	}
// 	res_bms := make([]*models.BookmarkResponse, 0, len(bms))
// 	for _, bm := range bms {
// 		res_bms = append(res_bms, models.NewBookmarkResponse(bm))
// 	}
// 	writeResponse(w, res_bms, nil)
// }

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
