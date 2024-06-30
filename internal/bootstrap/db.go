package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"main/internal/config"
	_ "modernc.org/sqlite"
)

func InitSqlxDB(cfg config.Config) (*sqlx.DB, error) {
	return sqlx.Connect(cfg.DBDriver, cfg.StoragePath+cfg.DBName)
}
