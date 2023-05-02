package interactor

/*
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．

interactorはアウトプットポートに依存し(importするということ)，
インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．
*/

import (
	"context"
	"ec_site_api/app/entities"
	"ec_site_api/app/usecases/port"
	"fmt"
)

type Gender int

const (
	Man   Gender = 1
	Woman Gender = 2
	Other Gender = 3
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

// usecase.UserInputPortを実装している
// GetUserByID は，UserRepo.GetUserByIDを呼び出し，その結果をOutputPort.Render or OutputPort.RenderErrorに渡します．
func (u *User) GetUserByID(ctx context.Context, userID string) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.Render(user)
}

// すべてのユーザーの情報を返却する
// GetUserByID は，UserRepo.GetUserByIDを呼び出し，その結果をOutputPort.Render or OutputPort.RenderErrorに渡します．
func (u *User) GetAllUser(ctx context.Context) {
	// users := []entity.User{}
	users, err := u.UserRepo.GetAllUser(ctx)
	fmt.Printf("%T\n", users)
	fmt.Println(users)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.RenderAll(users)
}

// ユーザーの情報を登録する
func (u *User) Insert(ctx context.Context, user *entities.User) {
	_, err := u.UserRepo.Insert(ctx, user)
	if err != nil {
		u.OutputPort.RenderError(err)
	}
	u.OutputPort.Render(user)
}

// ユーザーの情報を更新する
func (u *User) Update(ctx context.Context, user *entities.User) {
	_, err := u.UserRepo.Update(ctx, user)
	if err != nil {
		u.OutputPort.RenderError(err)
	}
	u.OutputPort.Render(user)
}
