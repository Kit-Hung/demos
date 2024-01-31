/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/30 19:53
 * @Description： zap 使用示例
 */
package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := getConfig()
	logger := zap.Must(config.Build())
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}(logger)
	logger.Info("testing zap: ", zap.String("test key", "test value"))
}

func getConfig() zap.Config {
	productionConfig := zap.NewProductionConfig()
	productionConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	productionConfig.OutputPaths = []string{"stdout"}
	productionConfig.EncoderConfig = getEncoderConfig()
	productionConfig.Encoding = "console"
	return productionConfig
}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zapcore.NewConsoleEncoder(encoderConfig)
	return encoderConfig
}
