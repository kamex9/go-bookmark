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
		logger.Info("Access Info", "method", r.Method, "path", r.URL.Path, "elapsed", elapsed)
	})
}

// ハンドラーのラッパー関数
// func withLogging(handler func(w http.ResponseWriter, r *http.Request) (any, error)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// リクエスト情報のログ
// 		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
// 		start := time.Now()

// 		// レスポンスヘッダーの設定
// 		w.Header().Set("Content-Type", "application/json")

// 		// 実際のハンドラー処理を実行
// 		data, err := handler(w, r)

// 		// Errorフィールドにはポインタ型初期値のnilが設定される
// 		// ※ポインタ型にしないと、ユーザー定義型の初期値である空オブジェクトが設定されてしまうためjson:omitemptyが効かなくなる
// 		response := Response{Data: data}
// 		if err != nil {
// 			log.Printf("Error processing request: %v", err)
// 			// http.Error(w, err.Error(), http.StatusBadRequest)
// 			w.WriteHeader(http.StatusBadRequest)
// 			pError := &Error{
// 				Code:    "E0001",
// 				Message: err.Error(),
// 			}
// 			// Dataフィールドにはany型初期値のnilが設定される
// 			response = Response{Error: pError}
// 		}

// 		// レスポンスの書き込み
// 		json.NewEncoder(w).Encode(response)

// 		// 処理時間のログ
// 		elapsed := time.Since(start)
// 		log.Printf("Request processed in %v", elapsed)
// 	}
// }

// func storeHandler(w http.ResponseWriter, r *http.Request) (any, error) {
// 	vars := mux.Vars(r)
// 	idStr, exists := vars["id"]
// 	if !exists {
// 		return nil, fmt.Errorf("store ID not found")
// 	}

// 	id, err := strconv.ParseUint(idStr, 10, 64)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid store ID")
// 	}

// 	bookmark, err := core.NewBookmark(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s.Add(*bookmark)
// 	return s.GetAll(), nil
// }

// func addHandler(w http.ResponseWriter, r *http.Request) (any, error) {
// 	vars := mux.Vars(r)
// 	idStr, exists := vars["id"]
// 	if !exists {
// 		return nil, fmt.Errorf("store ID not found")
// 	}

// 	id, err := utils.RetrieveGetParamUintValue(r, "id")
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid store ID")
// 	}

// 	var b B
// 	switch id {
// 	case 1:
// 		b =
// 	}
// }

// func resetHandler(w http.ResponseWriter, r *http.Request) (any, error) {
// 	s.Reset()
// 	return s.GetAll()
// }

func StartServer() {
	// ログの初期設定（タイムスタンプ、ファイル名、行番号を表示）
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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
