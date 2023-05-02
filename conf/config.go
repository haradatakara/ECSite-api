package conf

type ConfigList struct {
	Host    string `json:"host"`     // ホスト名
	Port    int    `json:"port"`     // ポート番号
	DbName  string `json:"db-name"`  // 接続先DB名
	Charset string `json:"charset"`  // 文字コード
	LogFile string `json:"log_file"` // ログファイル
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	Config = ConfigList{}
}
