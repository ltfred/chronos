package internal

import (
	"bytes"
	"github.com/spf13/cobra"
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
	cobra.CheckErr(viper.ReadInConfig())
	cobra.CheckErr(viper.Unmarshal(&Cfg))
}

func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/chronos/config.toml"
}

var config = `
name = "chronos"

[importantDay]
birthdays = [{name = "My birthday", date = "03-04", isLunar = false},{name = "Mom birthday", date = "03-07",isLunar = true}]
memorialDays = [{name = "Wedding anniversary", date = "12-25"}]
`

func InitConfigFile() {
	viper.SetConfigType("toml")
	viper.SetConfigFile(GetConfigPath())
	cobra.CheckErr(viper.ReadConfig(bytes.NewBuffer([]byte(config))))
	cobra.CheckErr(viper.WriteConfig())
}
