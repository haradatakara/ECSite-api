package conf

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type firebaseAppInterface interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

type firebaseApp struct {
	*firebase.App
}

func InitFirebaseApp() *firebaseApp {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal("Firebaseの初期化に失敗しました。")
	}
	return &firebaseApp{app}
}

func (app *firebaseApp) VerifyIDToken(ctx context.Context, idToken string) *auth.Token {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatal("認証に失敗しました。")
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Println("存在しないトークンです。")
	}
	return token
}
