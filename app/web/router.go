package web

import (
	"go-bookmark/core/logging"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var logger = logging.GetLogger()

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		logger.Info("Access info", "method", r.Method, "uri", r.URL.RequestURI(), "elapsed", elapsed)
	})
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	router.HandleFunc("/bookmark", createBookmark).Methods("POST")
	router.HandleFunc("/bookmark/{id}", fetchBookmarkById).Methods("GET")
	router.HandleFunc("/bookmarks", fetchAllBookmark).Methods("GET")

	logger.Info("Starting HTTP server on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
