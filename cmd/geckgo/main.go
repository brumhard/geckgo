package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"google.golang.org/grpc/reflection"

	geckgov1 "github.com/brumhard/geckgo/pkg/pb/geckgo/v1"
	"github.com/brumhard/geckgo/pkg/service"
	"google.golang.org/grpc"

	"github.com/brumhard/alligotor"

	"github.com/brumhard/geckgo/db"
	"github.com/golang-migrate/migrate/v4/source/httpfs"

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
			Addr string
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
			Addr string
		}{
			Addr: ":8080",
		},
	}

	cfgLoader := alligotor.New(alligotor.NewEnvSource("GECKGO"))
	if err := cfgLoader.Get(&cfg); err != nil {
		return err
	}

	connectString := fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		strconv.Itoa(cfg.DB.Port), cfg.DB.User, cfg.DB.Password, cfg.DB.DBName,
	)

	dbConnection, err := sqlx.Connect("pgx", connectString)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(dbConnection.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	source, err := httpfs.New(http.FS(db.Migrations), "migrations")
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

	// TODO: use propper logger
	logger := zap.NewNop()
	repo := service.NewRepository(dbConnection, logger)

	grpcServer := grpc.NewServer()
	geckgov1.RegisterGeckgoServiceServer(grpcServer, service.NewServer(repo, logger))
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", cfg.API.Addr)
	if err != nil {
		return err
	}

	return grpcServer.Serve(lis)
}
