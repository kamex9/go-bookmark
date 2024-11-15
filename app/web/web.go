package web

import (
	"encoding/json"
	"fmt"
	"go-bookmark/core"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var store = core.NewStore()

// 共通のレスポンス構造体
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data  any    `json:"data"`
	Error *Error `json:"error,omitempty"`
}

// ハンドラーのラッパー関数
func withLogging(handler func(w http.ResponseWriter, r *http.Request) (any, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエスト情報のログ
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		start := time.Now()

		// レスポンスヘッダーの設定
		w.Header().Set("Content-Type", "application/json")

		// 実際のハンドラー処理を実行
		data, err := handler(w, r)

		// Errorフィールドにはポインタ型初期値のnilが設定される
		// ※ポインタ型にしないと、ユーザー定義型の初期値である空オブジェクトが設定されてしまうためjson:omitemptyが効かなくなる
		response := Response{Data: data}
		if err != nil {
			log.Printf("Error processing request: %v", err)
			// http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			pError := &Error{
				Code:    "E0001",
				Message: err.Error(),
			}
			// Dataフィールドにはany型初期値のnilが設定される
			response = Response{Error: pError}
		}

		// レスポンスの書き込み
		json.NewEncoder(w).Encode(response)

		// 処理時間のログ
		elapsed := time.Since(start)
		log.Printf("Request processed in %v", elapsed)
	}
}

// 各ハンドラーの実装をシンプルに
func pingHandler(w http.ResponseWriter, r *http.Request) (any, error) {
	return map[string]string{"message": "pong"}, nil
}

func storeHandler(w http.ResponseWriter, r *http.Request) (any, error) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		return nil, fmt.Errorf("store ID not found")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid store ID")
	}

	bookmark, err := core.NewBookmark(id)
	if err != nil {
		return nil, err
	}

	store.Add(*bookmark)
	return store.GetAll(), nil
}

func resetHandler(w http.ResponseWriter, r *http.Request) (any, error) {
	store.Reset()
	return store.GetAll(), nil
}

func StartServer() {
	// ログの初期設定（タイムスタンプ、ファイル名、行番号を表示）
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	router := mux.NewRouter().StrictSlash(true)

	// ハンドラーの登録
	router.HandleFunc("/ping", withLogging(pingHandler)).Methods("GET")
	router.HandleFunc("/store/{id}", withLogging(storeHandler)).Methods("GET")
	router.HandleFunc("/reset", withLogging(resetHandler)).Methods("GET")

	// サーバー起動のログ
	log.Println("Starting HTTP server on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
