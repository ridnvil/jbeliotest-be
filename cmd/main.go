package main

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"jubeliotesting/internal/api"
	"jubeliotesting/internal/domain"
	"jubeliotesting/internal/repository"
	"jubeliotesting/internal/service"
	"jubeliotesting/internal/socketserver"
	"jubeliotesting/internal/worker"
	"jubeliotesting/pkg/config"
	"jubeliotesting/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cnfEnv config.GetEnvConfig
	if err := env.Parse(&cnfEnv); err != nil {
		log.Fatal(err)
	}

	db, err := config.CreateConnection(cnfEnv)
	if err != nil {
		log.Fatal(err)
	}

	if errmigg := db.AutoMigrate(&domain.Product{}, &domain.SubCategory{}, &domain.Category{}, &domain.InventoryMovement{}, &domain.Sale{}); errmigg != nil {
		log.Fatal(err)
	}

	logger.InitLogger()
	app := fiber.New()

	rdb := config.NewRedisClient(cnfEnv)
	channel := "process_dataset:insert"

	socket := socketserver.NewWebsocketServer(cnfEnv, rdb)
	socket.Start(app)

	salesRepo := repository.NewSalesRepository(db)
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	subCatRepo := repository.NewSubCategoryRepository(db)
	inventRepo := repository.NewInventoryMoveRepository(db)

	salesService := service.NewSalesService(salesRepo)
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	subCatService := service.NewSubCategoryService(subCatRepo)
	inventMoveService := service.NewInventoryMoveService(inventRepo)

	publishService := service.NewPublisherService(rdb, channel)

	subs := worker.NewSubcriber(salesService, productService, categoryService, subCatService, inventMoveService)
	go subs.StartSubscriber(ctx, rdb, channel, cnfEnv)

	// Route goes here
	publishHandler := api.NewPublishHandler(publishService, cnfEnv)
	publishHandler.PublishRoute(app)

	salesHandler := api.NewSalesHandler(salesService, cnfEnv)
	salesHandler.SalesRoute(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	if err = app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

	<-quit

	cancel()
	if err = app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
