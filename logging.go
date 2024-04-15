package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("this is an info message", zap.String("username", "johndoe"))
	// logger.Info("this is an info message", "username", "johndoe")
	sugar := logger.Sugar()
	defer logger.Sync()

	sugar.Debug("this is a debug message")
	sugar.Infof("this is an info message", 1234, "johndoe")
	sugar.Info("this is an info message")
	sugar.Warn("this is a warn message")
	// sugar.Error("this is an error message")
	// sugar.Fatal("this is a fatal message")
}
