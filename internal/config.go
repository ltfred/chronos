package internal

import (
	"github.com/spf13/viper"
	"os"
)

var Cfg Config

type Config struct {
	ImportantDay ImportantDay `json:"importantDay"`
}

type ImportantDay struct {
	Birthdays    []Day `json:"birthdays"`
	MemorialDays []Day `json:"memorialDays"`
}

type Day struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	IsLunar bool   `json:"isLunar"`
}

func SetupConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigFile(GetConfigPath())
	if err := viper.ReadInConfig(); err != nil {

	}
	if err := viper.Unmarshal(&Cfg); err != nil {

	}
}

func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/chronos/config.toml"
}
