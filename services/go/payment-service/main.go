package main

import (
	"context"
	"net/http"

	"monorepo-demonstration/services/go/common/pkg/config"
	"monorepo-demonstration/services/go/common/pkg/logger"
	"monorepo-demonstration/services/go/common/pkg/middleware"
	"monorepo-demonstration/services/go/internal/database"
	"monorepo-demonstration/services/go/internal/observability"
	paymenthandler "monorepo-demonstration/services/go/payment-service/handler"
	paymentrepo "monorepo-demonstration/services/go/payment-service/repository"
	paymentsvc "monorepo-demonstration/services/go/payment-service/service"
)

func main() {
	cfg := config.Load("payment-service")
	log := logger.New(cfg.ServiceName)

	observability.SetupTracing()

	if _, err := database.Connect(context.Background()); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := paymentrepo.NewInMemory()
	svc := paymentsvc.New(repo)
	handler := paymenthandler.New(svc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Printf("starting %s on :%s", cfg.ServiceName, cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, middleware.Logging(log, mux)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
