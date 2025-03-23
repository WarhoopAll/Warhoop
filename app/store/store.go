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
	AuthRepo  *bun.AuthRepo
	NexusRepo *bun.NexusRepo
	CharRepo  *bun.CharRepo
}

func NewBun(logger *log.Logger) (*Store, error) {
	cfg := config.Get()

	if cfg == nil {
		panic("configuration is nil")
	}

	auth, err := connect(cfg.DBAuth, "auth")
	if err != nil {
		logger.Error("store.NewBun.auth",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	char, err := connect(cfg.DBCharacters, "characters")
	if err != nil {
		logger.Error("store.NewBun.characters",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	nexus, err := connect(cfg.DBNexus, "nexus")
	if err != nil {
		logger.Error("store.NewBun.nexus",
			log.String("error", err.Error()),
		)
		return nil, err
	}

	nexusRepo := bun.NewNexusRepo(nexus, logger)
	if nexusRepo == nil {
		logger.Error("store.NewBun.NexusRepo",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	authRepo := bun.NewAuthRepo(auth, logger, nexusRepo)
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
		AuthRepo:  authRepo,
		CharRepo:  charRepo,
		NexusRepo: nexusRepo,
	}, nil
}

func connect(dsn, name string) (*bun.DB, error) {
	db, err := bun.Dial(dsn)
	if err != nil {
		return nil, err
	}

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(config.Get().DBVerbose)))
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName(name),
		bunotel.WithFormattedQueries(true),
	))

	return db, nil
}
