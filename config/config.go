package imemo

import (
	"github.com/BurntSushi/toml"
)

var config *Config

type Config struct {
	Sys SysConfig `toml:Sys`
	App AppConfig `toml:App`
}

type SysConfig struct {
	FileName string `toml:fileName`
}

type AppConfig struct {
	Name    string `toml:name`
	Version string `toml:version`
}

func init() {
	toml.DecodeFile("config.toml", &config)
}

func GetConfig() *Config {
	return config
}
