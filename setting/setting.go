package setting

import (
    "io/ioutil"
    "fmt"
	"encoding/json"
)

const (
    // DefaultPort - デフォルトポート
    DefaultPort = "80"
    // DefaultOracleUser - デフォルトOracleユーザ
    DefaultOracleUser = "USER"
    // DefaultOraclePass - デフォルトOracleパスワード
    DefaultOraclePass = "PASS"
    // DefaultOracleIP - デフォルトOracleIPアドレス
    DefaultOracleIP = "127.0.0.1"
    // DefaultOraclePort - デフォルトOracleポート番号
    DefaultOraclePort = "1521"
    // DefaultOracleService - デフォルトOracleサービスネーム
    DefaultOracleService = "ORCL"
)

// Config - ユーザ設定ファイル用
type Config struct {
    Port string `json:"port"`
    OracleUser string `json:"oracle_user"`
    OraclePass string `json:"oracle_pass"`
    OracleIP string `json:"oracle_ip"`
    OraclePort string `json:"oracle_port"`
    OracleService string `json:"oracle_service"`
}

// GetConfig - ユーザ設定読み込み
func GetConfig() Config {
    conf, err := configParse("./config.json")
    if err != nil {
        fmt.Printf("error: %v ", err)
        // デフォルト設定
        conf.Port = DefaultPort
        conf.OracleUser = DefaultOracleUser
        conf.OraclePass = DefaultOraclePass
        conf.OracleIP = DefaultOracleIP
        conf.OraclePort = DefaultOraclePort
        conf.OracleService = DefaultOracleService
    }
    return conf
}

func configParse(filename string) (Config, error) {
    var c Config
    jsonStr, err := ioutil.ReadFile(filename)

    if err != nil {
        fmt.Printf("error: %v ", err)
        return c, err
    }

    err = json.Unmarshal(jsonStr, &c)
    if err != nil {
        fmt.Printf("error: %v ", err)
        return c, err
    }
    return c, nil
}