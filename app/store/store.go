package store

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"
	"warhoop/app/config"
	"warhoop/app/log"
	"warhoop/app/store/bun"
)

// Store ...
type Store struct {
	AuthRepo *bun.AuthRepo
	SaitRepo *bun.SaitRepo
	CharRepo *bun.CharRepo
}

var cfg = config.Get()

func NewBun(logger *log.Logger) (*Store, error) {
	if cfg == nil {
		panic("configuration is nil")
	}

	auth, err := connect(cfg.DB.Auth, "auth")
	if err != nil {
		logger.Error("store.NewBun.auth",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	char, err := connect(cfg.DB.Characters, "characters")
	if err != nil {
		logger.Error("store.NewBun.characters",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	sait, err := connect(cfg.DB.Sait, "sait")
	if err != nil {
		logger.Error("store.NewBun.sait",
			log.String("error", err.Error()),
		)
		return nil, err
	}

	saitRepo := bun.NewSaitRepo(sait, logger)
	if saitRepo == nil {
		logger.Error("store.NewBun.SaitRepo",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	authRepo := bun.NewAuthRepo(auth, logger, saitRepo)
	if authRepo == nil {
		if err != nil {
			logger.Error("tore.NewBun.AuthRepo",
				log.String("err", err.Error()),
			)
			return nil, err
		}
	}

	charRepo := bun.NewCharRepo(char, logger)
	if charRepo == nil {
		if err != nil {
			logger.Error("tore.NewBun.charRepo",
				log.String("err", err.Error()),
			)
			return nil, err
		}
	}
	return &Store{
		AuthRepo: authRepo,
		CharRepo: charRepo,
		SaitRepo: saitRepo,
	}, nil
}

func connect(dsn, name string) (*bun.DB, error) {
	db, err := bun.Dial(dsn)
	if err != nil {
		return nil, err
	}

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(cfg.DB.Verbose)))
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName(name),
		bunotel.WithFormattedQueries(true),
	))

	return db, nil
}
