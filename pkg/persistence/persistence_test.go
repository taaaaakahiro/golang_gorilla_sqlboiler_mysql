package persistence

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/config"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/io"
)

var (
	ctx        context.Context
	reviewRepo *ReviewRepository
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
	}
	db, _ := io.NewDatabase(sqlSetting)

	reviewRepo = NewReviewRepository(db)

	res := m.Run()
	// after

	os.Exit(res)
}
