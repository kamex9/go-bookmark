package web

import (
	"encoding/json"
	"go-bookmark/core/models"
	"go-bookmark/core/repository"
	"go-bookmark/core/utils"
	"io"
	"log"
	"net/http"
)

var repo = repository.NewRepository(repository.MEMORY)

func createBookmark(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var req models.BookmarkPostRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		log.Fatal(err)
	}

	bm := models.NewBookmark(req)
	repo.Save(bm)
	bms, _ := repo.FindAll()
	json.NewEncoder(w).Encode(bms)
}

func fetchAllBookmark(w http.ResponseWriter, r *http.Request) {
	bms, _ := repo.FindAll()
	json.NewEncoder(w).Encode(bms)
}

func fetchBookmarkById(w http.ResponseWriter, r *http.Request) {
	id, _ := utils.RetrievePathParamValue("id", r)
	bm, _ := repo.FindById(id)
	json.NewEncoder(w).Encode(bm)
}
