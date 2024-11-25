package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"project/config"
	"strings"
)

func InitLog(config config.AppConfig) *zap.Logger {
	if !strings.HasPrefix(config.LogPath, "/") {
		log.Fatal("invalid log path")
		return nil
	}
	
	if _, err := os.Stat(".." + config.LogPath); os.IsNotExist(err) {
		err = os.MkdirAll(".."+config.LogPath, 0750)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	logLevel := map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
	}

	file, err := os.Create(".." + config.LogPath + "/app.log")
	if err != nil {
		log.Fatal(err)
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(file),
		logLevel[config.LogLevel],
	)
	return zap.New(core)
}
