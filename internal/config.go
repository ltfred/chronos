package internal

import (
	"bytes"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var Cfg Config
var OriginCfg string

type Config struct {
	Days []Day `json:"days"`
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
	s := viper.AllSettings()
	marshal, _ := toml.Marshal(s)
	OriginCfg = string(marshal)
	cobra.CheckErr(viper.Unmarshal(&Cfg))
}

func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/chronos/config.toml"
}

var config = `
name = "chronos"

days = [
{name = "My birthday", date = "03-04", isLunar = false},
{name = "Mom birthday", date = "03-07",isLunar = true}
]
`

func InitConfigFile() {
	viper.SetConfigType("toml")
	viper.SetConfigFile(GetConfigPath())
	cobra.CheckErr(viper.ReadConfig(bytes.NewBuffer([]byte(config))))
	cobra.CheckErr(viper.WriteConfig())
}
