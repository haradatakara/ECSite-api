package port

/*
port パッケージは，出力や入力などのポート(interface)を提供します．
*/

import (
	"context"
	"ec_site_api/app/entities"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
	GetAllUser(ctx context.Context)
	Insert(ctx context.Context, user *entities.User)
	Update(ctx context.Context, user *entities.User)
}

type UserOutputPort interface {
	Render(*entities.User)
	RenderAll([]*entities.User)
	RenderError(error)
}

// userのCRUDに対するDB用のポート
type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entities.User, error)
	GetAllUser(ctx context.Context) ([]*entities.User, error)
	Insert(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
}
