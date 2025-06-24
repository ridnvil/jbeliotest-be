package main

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"
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

func initTracer(ctx context.Context, cnfg config.GetEnvConfig) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(cnfg.TempoEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	res, _ := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("gofiber-app"),
		),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func main() {
	var cnfEnv config.GetEnvConfig
	if err := env.Parse(&cnfEnv); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp, err := initTracer(ctx, cnfEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = tp.Shutdown(ctx)
	}()

	otel.SetTracerProvider(tp)

	db, err := config.CreateConnection(cnfEnv)
	if err != nil {
		log.Fatal(err)
	}

	if errmigg := db.AutoMigrate(&domain.Product{}, &domain.SubCategory{}, &domain.Category{}, &domain.InventoryMovement{}, &domain.Sale{}); errmigg != nil {
		log.Fatal(err)
	}

	logger.InitLogger()
	app := fiber.New(fiber.Config{
		AppName: "gofiber-app",
	})
	app.Use(otelfiber.Middleware())

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

	go func() {
		// Lakukan pekerjaan yang memerlukan trace di sini
		if err = app.Listen(":3000"); err != nil {
			log.Fatal(err)
		}
	}()

	<-quit

	cancel()
	if err = app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
