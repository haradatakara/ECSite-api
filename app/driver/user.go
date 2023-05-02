package driver

/**
MySQLとの接続を設定する
*/

import (
	"database/sql"
	"ec_site_api/app/adapter/controller"
	"ec_site_api/app/adapter/gateway"
	"ec_site_api/app/adapter/presenter"
	"ec_site_api/app/usecases/interactor"
	"ec_site_api/conf"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Serve(addr string) {
	// 設定値から接続文字列を生成
	conStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DATABASE"),
		os.Getenv("CHARSET"))

	// データベース接続
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		fmt.Println(err.Error())
		// deferで処理終了前に必ず接続をクローズする
		defer db.Close()
	}

	// 接続確認
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	}

	fmt.Println("データベース接続成功！")

	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          db,
	}

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})

	//ルータの設定
	router := mux.NewRouter()

	//ユーザーの新規登録に関するルータ
	userRegisterRouter := router.PathPrefix("/auth").Subrouter()
	userRegisterRouter.HandleFunc("/", user.Insert).Methods("POST") // ユーザー情報を登録

	//ユーザー情報に関するルータ
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", user.Update).Methods("PUT")           //ユーザー情報を更新
	userRouter.HandleFunc("/{uid}", user.GetUserByID).Methods("GET") //ユーザー情報を取得
	userRouter.Use(conf.AuthMiddleware)

	//アイテムに関するルータ
	// itemRouter := router.PathPrefix("/items").Subrouter()

	//サーバー起動
	err = http.ListenAndServe(addr, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}

}
