package main

import (
	"context"
	"net/http"

	"monorepo-demonstration/services/go/common/pkg/config"
	"monorepo-demonstration/services/go/common/pkg/logger"
	"monorepo-demonstration/services/go/common/pkg/middleware"
	"monorepo-demonstration/services/go/internal/database"
	"monorepo-demonstration/services/go/internal/observability"
	userhandler "monorepo-demonstration/services/go/user-service/handler"
	userrepo "monorepo-demonstration/services/go/user-service/repository"
	usersvc "monorepo-demonstration/services/go/user-service/service"
)

func main() {
	cfg := config.Load("user-service")
	log := logger.New(cfg.ServiceName)

	observability.SetupTracing()

	if _, err := database.Connect(context.Background()); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := userrepo.NewInMemory()
	svc := usersvc.New(repo)
	handler := userhandler.New(svc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Printf("starting %s on :%s", cfg.ServiceName, cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, middleware.Logging(log, mux)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
