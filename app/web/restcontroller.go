package web

import (
	"encoding/json"
	"go-bookmark/core/models"
	"go-bookmark/core/repository"
	"go-bookmark/core/utils"
	"io"
	"net/http"
)

var repo = repository.NewRepository(repository.MEMORY)

func createBookmark(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	logger.Info("Params info", "body", string(reqBody))
	var req models.BookmarkPostRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		writeResponse(w, nil, models.NewError("E0001", err))
		return
	}

	bm := models.NewBookmark(req)
	repo.Save(bm)
	writeResponse(w, models.NewBookmarkResponse(bm), nil)
}

func fetchAllBookmark(w http.ResponseWriter, r *http.Request) {
	bms, err := repo.FindAll()
	if err != nil {
		writeResponse(w, nil, models.NewError("E1001", err))
		return
	}
	res_bms := make([]*models.BookmarkResponse, 0, len(bms))
	for _, bm := range bms {
		res_bms = append(res_bms, models.NewBookmarkResponse(bm))
	}
	writeResponse(w, res_bms, nil)
}

func fetchBookmarkById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.RetrievePathParamValue("id", r)
	if err != nil {
		writeResponse(w, nil, models.NewError("E2001", err))
		return
	}
	bm, err := repo.FindById(id)
	if err != nil {
		writeResponse(w, nil, models.NewError("E2002", err))
		return
	}
	writeResponse(w, models.NewBookmarkResponse(bm), nil)
}

func writeResponse(w http.ResponseWriter, data any, err *models.Error) {
	var response models.Response
	if err != nil {
		logger.Error("Error occurred in processing", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		// Dataフィールドにはany型初期値のnilが設定される
		response = models.Response{Error: err}
	} else {
		// Errorフィールドにはポインタ型初期値のnilが設定される
		// ※ポインタ型にしないと、ユーザー定義型の初期値である空オブジェクトが設定されてしまうためjson:omitemptyが効かなくなる
		response = models.Response{Data: data}
	}

	json.NewEncoder(w).Encode(response)
}
