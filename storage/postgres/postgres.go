package postgres

import (
	"context"
	"ecommerce/configs"
	log "ecommerce/pkg/logger"
	"ecommerce/storage"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
	_ "github.com/lib/pq"
)

type Store struct {
	db  *pgxpool.Pool
	cfg configs.Config
	log log.ILogger
}

func New(ctx context.Context, cfg configs.Config, log log.ILogger) (storage.IStorage, error) {
	url := fmt.Sprintf(
		`host=%s port=%s user=%s dbname=%s password=%s sslmode=disable`,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPassword,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return &Store{}, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return &Store{}, err
	}

	// migration
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		return &Store{}, err
	}

	if err := m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			fmt.Println("entered", err)
			version, dirty, err := m.Version()
			if err != nil {
				return &Store{}, err
			}

			if dirty {
				version--
				if err := m.Force(int(version)); err != nil {
					return &Store{}, err
				}
			}
			return &Store{}, err
		}
	}

	return &Store{
		db:  pool,
		cfg: cfg,
		log: log,
	}, nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Admin() storage.IAdminsStorage {
	return NewAdminRepo(s.db, s.log)
}
