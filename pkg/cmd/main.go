package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/config"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/server"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	port := fmt.Sprintf(":%s", strconv.Itoa(cfg.Port))

	err = server.NewServer(port)

	log.Fatal(err)
}
