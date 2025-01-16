package bun

import (
	"database/sql"
	"errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"warhoop/app/log"
)

type DB struct {
	*bun.DB
}

func Dial(dbDSN string) (*DB, error) {
	logger := log.Get()

	if dbDSN == "" {
		err := errors.New("missing DSN: database connection string is empty")
		logger.Error("store.Dial",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	sqlDB, err := sql.Open("mysql", dbDSN)
	if err != nil {
		logger.Error("store.Dial",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		logger.Error("store.Dial",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	bunDB := bun.NewDB(sqlDB, mysqldialect.New(), bun.WithDiscardUnknownColumns())

	logger.Debug("store.Dial",
		log.String("DSN", dbDSN),
	)

	return &DB{bunDB}, nil
}
