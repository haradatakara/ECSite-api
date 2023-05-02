package controller

/*
controller パッケージは，入力に対するアダプターです．

ここでは，インプットポートとアウトプットポートを組み立てて，
インプットポートを実行します．
ユースケースレイヤからの戻り値を受け取って出力する必要はなく，
純粋にhttpを受け取り，ユースケースを実行します．
ここではインプットポートに入力されたデータを変換し渡す。
*/

import (
	"context"
	"database/sql"
	"ec_site_api/app/entities"
	"ec_site_api/app/usecases/port"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> presenter.NewUserOutputPort
	InputFactory func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	// -> interactor.NewUserInputPort
	RepoFactory func(c *sql.DB) port.UserRepository

	Conn *sql.DB
}

// GetUserByID は，httpを受け取り，portを組み立てて，inputPort.GetUserByIDを呼び出します．
func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	u.newInputPort(w).GetUserByID(r.Context(), userID)
}

func (u *User) Insert(w http.ResponseWriter, r *http.Request) {
	user := &entities.User{}

	if err := decodeInputUserInfo(r, user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(user.Gender)

	if err := validInputUserInfo(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("error: %s", err.Error())
		return
	}
	fmt.Println(user)

	u.newInputPort(w).Insert(r.Context(), user)
}

func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	user := &entities.User{}

	if err := decodeInputUserInfo(r, user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("error: %s", err.Error())
		return
	}

	if err := validInputUserInfo(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("error: %s", err.Error())
		return
	}

	u.newInputPort(w).Update(r.Context(), user)
}

func (u *User) GetAllUser(w http.ResponseWriter, ctx context.Context) {
	u.newInputPort(w).GetAllUser(ctx)
}

func (u *User) newInputPort(w http.ResponseWriter) port.UserInputPort {
	outputPort := u.OutputFactory(w)
	repository := u.RepoFactory(u.Conn)
	return u.InputFactory(outputPort, repository)
}

// 入力されたユーザー情報が正しいかどうかチェックする
func validInputUserInfo(user *entities.User) error {
	fmt.Println(user.Mail)

	return user.Validate()
}

// 入力されたユーザー情報をデコードする
func decodeInputUserInfo(r *http.Request, user *entities.User) error {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user pointer is nil")
	}

	return nil
}
