package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"sqlite-agent/api"
	"sqlite-agent/config"
	"sqlite-agent/util/log"

	"go.uber.org/zap"
)

func main() {
	var (
		configPath string
		dev        bool
	)

	flag.StringVar(&configPath, "config", "", "Path to the configuration file")
	flag.BoolVar(&dev, "dev", false, "Run in development mode")
	flag.Parse()

	if dev {
		config.RunModel = config.RunModelDevValue
	}

	if err := config.LoadConfig(configPath); err != nil {
		zap.L().Fatal("Failed to load configuration file", zap.Error(err))
	}

	log.SetLogger()

	r := api.InitApi()
	go r.Run(fmt.Sprintf(":%d", config.ListenPort))

	// 监听停止信号
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-sigs
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			zap.L().Info("Received shutdown signal, exiting...", zap.String("signal", sig.String()))
			return
		}
	}
}
