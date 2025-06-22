package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	_ "go.uber.org/zap/zapcore"
	"log"
	"os"
)

var Log *zap.Logger

func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Can't open log file: %v", err)
	}

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(file),
		zap.InfoLevel,
	)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	core := zapcore.NewTee(fileCore, consoleCore)
	Log = zap.New(core, zap.AddCaller())
}
