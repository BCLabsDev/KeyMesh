package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Log
var Log_Level = 0

// DB
var (
	DB_Host = ""
	DB_Port = 0
	DB_User = ""
	DB_Pass = ""
	DB_Name = ""
)

// config 结构体映射 TOML
type config struct {
	Log struct {
		Level int `toml:"level"`
	} `toml:"log"`
	DB struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		Name string `toml:"name"`
	} `toml:"db"`
}

func init() {
	var cfg config
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "无法解析 TOML 配置文件: %v\n", err)
		os.Exit(1)
	}
	// log
	Log_Level = cfg.Log.Level
	// DB
	DB_Host = cfg.DB.Host
	DB_Port = cfg.DB.Port
	DB_Name = cfg.DB.Name
	DB_User = cfg.DB.User
	DB_Pass = cfg.DB.Pass
}
