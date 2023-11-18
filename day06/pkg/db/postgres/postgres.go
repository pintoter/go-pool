package postgres

import (
	"database/sql"
	"day06/internal/entity"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	driverName = "postgres"
)

type Config interface {
	GetDSN() string
}

func ConnectDB(cfg Config) (*gorm.DB, error) {
	postgreDB, err := sql.Open(driverName, cfg.GetDSN())
	if err != nil {
		return nil, errors.Wrap(err, "error openning DB")
	}

	if err := postgreDB.Ping(); err != nil {
		return nil, errors.Wrap(err, "error ping DB")
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: postgreDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "error openning gorm")
	}

	gormDB.AutoMigrate(&entity.Article{})

	return gormDB, nil
}
