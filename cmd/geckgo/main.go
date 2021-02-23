package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	db2 "github.com/brumhard/geckgo/db"
	"github.com/golang-migrate/migrate/v4/source/httpfs"

	"github.com/brumhard/alligotor"

	"github.com/brumhard/geckgo/pkg"
	kitlog "github.com/go-kit/kit/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := struct {
		DB struct {
			Port     int
			User     string
			Password string
			DBName   string
		}
		API struct {
			Port int
		}
	}{
		DB: struct {
			Port     int
			User     string
			Password string
			DBName   string
		}{
			Port:     5432,
			DBName:   "postgres",
			User:     "postgres",
			Password: "Pass2020!",
		},
		API: struct {
			Port int
		}{
			Port: 8080,
		},
	}

	cfgLoader := alligotor.New(alligotor.FromEnvVars("GECKGO"))
	if err := cfgLoader.Get(&cfg); err != nil {
		return err
	}

	logger := kitlog.NewJSONLogger(os.Stdout)

	connectString := fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		strconv.Itoa(cfg.DB.Port), cfg.DB.User, cfg.DB.Password, cfg.DB.DBName,
	)

	db, err := sqlx.Connect("pgx", connectString)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	source, err := httpfs.New(http.FS(db2.Migrations), "migrations")
	if err != nil {
		return err
	}

	defer source.Close()

	migrations, err := migrate.NewWithInstance("httpfs", source, cfg.DB.DBName, driver)
	if err != nil {
		return err
	}

	if err := migrations.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}

	repo := pkg.NewRepository(db, logger)
	service := pkg.NewService(repo, logger)

	return http.ListenAndServe(":"+strconv.Itoa(cfg.API.Port), pkg.MakeHandler(service, logger))
}
