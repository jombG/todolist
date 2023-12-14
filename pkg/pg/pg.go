package pg

import (
	"database/sql"
	ent "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"

	_ "github.com/lib/pq"
)

type (
	// ConnectionDescriber описание конфига подключения к БД, является универсальным представлением параметров.
	ConnectionDescriber interface {
		DataSourceName() string // Возвращает DSN для создания подключения к БД, пригодный для использования в sql.Open().
		DriverName() string
		MigrationSource() string // Возвращает источник миграционных скриптов.
	}
)

func EntDriver(cfg ConnectionDescriber) (*ent.Driver, error) {
	conn, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	if err = migrations(cfg); err != nil {
		return nil, err
	}

	client := ent.OpenDB(cfg.DriverName(), conn)

	return client, nil
}

func migrations(cfg ConnectionDescriber) error {
	conn, err := goose.OpenDBWithDriver("postgres", cfg.DataSourceName())
	if err != nil {
		return err
	}

	goose.SetBaseFS(nil)

	if err := goose.Up(conn, cfg.MigrationSource()); err != nil {
		log.Fatalf("goose: failed goose up migration: %v", err)
		return err
	}

	return nil
}

func connect(cfg ConnectionDescriber) (*sql.DB, error) {
	db, err := sql.Open(cfg.DriverName(), cfg.DataSourceName())
	if err != nil {
		return nil, fmt.Errorf("open %s connection (%v): %w", cfg.DriverName(), cfg, err)
	}

	db.SetConnMaxIdleTime(0)
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't check db connection (%v): %w", cfg, err)
	}

	return db, nil
}
