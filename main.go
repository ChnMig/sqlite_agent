package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"sqlite-agent/api/app"
	"sqlite-agent/config"
	"sqlite-agent/db"
	"sqlite-agent/util/log"

	"go.uber.org/zap"
)

func main() {
	var (
		configPath string
		dev        bool
		check      bool // 新增 check 参数
	)

	flag.StringVar(&configPath, "config", "", "Path to the configuration file")
	flag.BoolVar(&dev, "dev", false, "Run in development mode")
	flag.BoolVar(&check, "check", false, "Check DB connection and exit") // 新增 check 参数
	flag.Parse()

	if dev {
		config.RunModel = config.RunModelDevValue
		fmt.Println("Running in development mode")
	}

	if err := config.LoadConfig(configPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	if check {
		config.RunModel = config.RunModelDevValue
		log.SetLogger()
		if err := db.Init(); err != nil {
			fmt.Fprintf(os.Stderr, "DB connection check failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("DB connection check succeeded")
		return
	}

	log.SetLogger()

	r := app.InitApi()
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
