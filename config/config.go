package config

import (
	"fmt"
	"os"
	"path/filepath"

	pathtool "sqlite-agent/util/path-tool"
)

// Here are some basic configurations
// These configurations are usually generic
var (
	// listen
	ListenPort = 10086 // api listen port
	// run model
	RunModelKey      = "model"
	RunModel         = ""
	RunModelDevValue = "dev"
	RunModelRelease  = "release"
	// path
	SelfName = filepath.Base(os.Args[0])      // own file name
	AbsPath  = pathtool.GetCurrentDirectory() // current directory
	// log
	LogDir        = filepath.Join(pathtool.GetCurrentDirectory(), "log")   // log directory
	LogPath       = filepath.Join(LogDir, fmt.Sprintf("%s.log", SelfName)) // self log path
	LogMaxSize    = 50                                                     // M
	LogMaxBackups = 3                                                      // backups
	LogMaxAge     = 30                                                     // days
	LogModelDev   = "dev"                                                  // dev model
	// api
	ApiAuth string
	// default config file path
	DefaultConfigPath = filepath.Join(AbsPath, "config.yaml")
)

// 修改配置结构以支持分层的 YAML 配置
var Config struct {
	Server struct {
		ListenPort int    `yaml:"listen_port"`
		ApiAuth    string `yaml:"api_auth"`
	} `yaml:"server"`
	Log struct {
		MaxSize    int `yaml:"max_size"`
		MaxBackups int `yaml:"max_backups"`
		MaxAge     int `yaml:"max_age"`
	} `yaml:"log"`
}

// InitConfig initializes the configuration
func init() {
	pathtool.CreateDir(LogDir)
}
