package database

import (
	"database/sql"

	_ "github.com/mattn/go-oci8"

    "app/setting"
)

func DbOpen() (*sql.DB, error) {
    // 設定取得
    conf := setting.GetConfig()

    connectStr := conf.OracleUser + "/" + conf.OraclePass + "@" + conf.OracleIP +
                    ":" + conf.OraclePort + "/" + conf.OracleService

    db, err := sql.Open("oci8", connectStr)

    return db, err
}