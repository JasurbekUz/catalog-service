package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/JasurbekUz/catalog-service/config"
	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
	"github.com/JasurbekUz/catalog-service/pkg/db"
	"github.com/JasurbekUz/catalog-service/pkg/logger"
	"github.com/JasurbekUz/catalog-service/service"
	"github.com/JasurbekUz/catalog-service/storage"
)

func main() {
	// run config load function
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "catalog-service")
	defer func(l logger.Logger) {
		err:= logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)
	log.Info("main:sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)
	// connect postgres
	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	pgStorage := storage.NewStoragePg(connDB)
	//run service
	catalogService := service.NewCatalogService(pgStorage, log)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
	s := grpc.NewServer()
	cs.RegisterCatalogServiceServer(s, catalogService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
