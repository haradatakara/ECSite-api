package entities

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

/*
entity パッケージは，ドメインモデルを実装します．．
*/

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Gender  int    `json:"gender,string"`
	Address string `json:"address"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(
			&u.Name,
			validation.Required.Error("名前は必須入力です"),
			validation.Length(1, 50).Error("名前は５０文字以内で入力してください")),
		validation.Field(
			&u.Mail,
			validation.Required.Error("メールアドレスは必須入力です"),
			validation.Match(regexp.MustCompile("^[A-Za-z0-9]{1}[A-Za-z0-9_.-]*@[A-Za-z0-9_.-]+.[A-Za-z0-9]+$")).Error("メールアドレスの形式に誤りがあります")),
		validation.Field(
			&u.Gender,
			validation.Required.Error("性別は選択必須です"),
			validation.In(1, 2, 3).Error("性別は不正な入力です")),
		validation.Field(
			&u.Address,
			validation.Length(0, 100).Error("住所は１００文字以内で入力してください")),
	)
}
