package conf

import (
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if !strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ") {
			log.Fatal("Authorization ヘッダのフォーマットが適切でありません。")
		}

		idToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

		//firebaseSDKの初期化を行う
		client := InitFirebaseApp()

		//トークンIDの確認を行う
		token := client.VerifyIDToken(ctx, idToken)

		log.Printf("Vertifed ID token: %v\n", token)

		next.ServeHTTP(w, r)
	})
}
