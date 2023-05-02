package gateway

/*
gateway パッケージは，DB操作に対するアダプターです．
*/

import (
	"context"
	"database/sql"
	"ec_site_api/app/entities"
	"ec_site_api/app/usecases/port"
	"errors"
	"fmt"
	"log"
	"time"
)

type UserRepository struct {
	conn *sql.DB
}

// NewUserRepository はUserRepositoryを返します．
func NewUserRepository(conn *sql.DB) port.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// GetUserByID はDBからデータを取得します．
func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*entities.User, error) {
	conn := u.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT id, name, mail, gender_id, COALESCE(address, '') FROM `user` WHERE id = ?", userID)
	user := entities.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Mail, &user.Gender, &user.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found. UserID = %s", userID)
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapter/gateway/GetUserByID")
	}
	return &user, nil
}

// GetUserByID はDBからデータを取得します．
func (u *UserRepository) GetAllUser(ctx context.Context) ([]*entities.User, error) {
	conn := u.GetDBConn()
	rows, errors := conn.QueryContext(ctx, "SELECT * FROM `user`")

	if errors != nil {
		log.Fatal("Not Found!!")
	}

	defer rows.Close()

	users := []*entities.User{}
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalln(err)
		}

		user := entities.User{
			ID:   id,
			Name: name,
		}
		// 配列に追加
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}
	return users, nil
}

func (u *UserRepository) Insert(ctx context.Context, user *entities.User) (*entities.User, error) {
	if user == nil {
		return nil, errors.New("ユーザー情報が存在しません。")
	}
	conn := u.GetDBConn()
	_, err := conn.Exec(
		"INSERT INTO user(id, name, mail, gender_id, created_at) VALUES (?, ?, ?, ?, ?)",
		user.ID,
		user.Name,
		user.Mail,
		user.Gender,
		time.Now())
	if err != nil {
		return nil, err
	}

	log.Println("ユーザー情報の登録に成功しました。")

	return user, nil
}

func (u *UserRepository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	if user == nil {
		return nil, errors.New("ユーザー情報が存在しません。")
	}
	conn := u.GetDBConn()
	_, err := conn.Exec(
		"UPDATE user SET name = ? , mail = ?, gender_id = ?, address = ?, updated_at = ? WHERE id = ?",
		user.Name,
		user.Mail,
		user.Gender,
		user.Address,
		time.Now(),
		user.ID)
	if err != nil {
		return nil, err
	}

	log.Println("ユーザー情報の更新に成功しました。")
	return user, nil
}

// GetDBConn はconnectionを取得します．
func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}
