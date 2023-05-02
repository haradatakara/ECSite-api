package utils

import (
	"net/http"
)

type OptionHandle func(w http.ResponseWriter, r *http.Request)

func OptionHandler(handle OptionHandle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ヘッダの追加
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		//プリフライトリクエストへの応答
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		//Handler関数の実行
		handle(w, r)
	}
}
