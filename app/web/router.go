package web

import (
	"go-bookmark/app/web/restcontroller"
	"go-bookmark/core/constants"
	"go-bookmark/core/usecase"
	"go-bookmark/core/utils"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var l = utils.GetLogger()

var brc = restcontroller.NewBookmarkRestController(
	l,
	usecase.NewCrudBookmarkService(constants.MEMORY),
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		l.Info("Access info", "method", r.Method, "uri", r.URL.RequestURI(), "elapsed", elapsed)
	})
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	router.HandleFunc("/bookmark", brc.Create).Methods("POST")
	router.HandleFunc("/bookmark/{id}", brc.FetchById).Methods("GET")
	router.HandleFunc("/bookmarks", brc.FetchAll).Methods("GET")
	router.HandleFunc("/bookmarks", brc.DeleteAll).Methods("DELETE")

	l.Info("Starting HTTP server on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		l.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
