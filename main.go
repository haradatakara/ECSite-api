package main

/*
プログラムの開始地点
*/
import (
	"ec_site_api/app/driver"
	"ec_site_api/conf"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ※Goではコードの記述順序は関係ないので、上に書いても下に書いても構いません。
func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("out!!")
	}
	//ログ設定の初期化
	conf.LoggingInit()
	fmt.Println(os.Getenv("API_PORT"))
	driver.Serve(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
}
