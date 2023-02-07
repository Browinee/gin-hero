package mysql

import (
	"fmt"
	"master-gin/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.HOST, cfg.Port, cfg.DbName)

	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		zap.L().Error("connect DB failed, err %v/n", zap.Error(err))
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
