package presenter

/*
presenter パッケージは，出力に対するアダプターです．

ここでは，アウトプットポートを実装します(interfaceを満たすようにmethodを追加するということ)
*/

import (
	"encoding/json"
	"log"
	"net/http"

	"ec_site_api/app/entities"
	"ec_site_api/app/usecases/port"
)

type User struct {
	w http.ResponseWriter
}

// NewUserOutputPort はUserOutputPortを取得します．
func NewUserOutputPort(w http.ResponseWriter) port.UserOutputPort {
	return &User{
		w: w,
	}
}

// usecase.UserOutputPortを実装している
// Render はNameを出力します．
func (u *User) Render(user *entities.User) {
	u.w.WriteHeader(http.StatusOK)
	json.NewEncoder(u.w).Encode(user)
}

// usecase.UserOutputPortを実装している
// Render はNameを出力します．
func (u *User) RenderAll(users []*entities.User) {
	println(&users)
	u.w.WriteHeader(http.StatusOK)
	json.NewEncoder(u.w).Encode(users)
}

// RenderError はErrorを出力します．
func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	log.Println(u.w, err)
}
