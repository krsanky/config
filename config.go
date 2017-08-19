package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var GOPATH string
var baseDir string

func init() {
	viper.BindEnv("GOPATH")
	GOPATH = viper.GetString("GOPATH")
	baseDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("config.init() err:%s", err))
	}
	fmt.Printf("baseDir:%s\n", baseDir)

	viper.AddConfigPath(baseDir)
	viper.SetConfigName("settings")
	Read()

	viper.SetDefault("template_dir", fmt.Sprintf("%s/template", baseDir))
	viper.SetDefault("static_dir", fmt.Sprintf("%s/static", baseDir))
	viper.SetDefault("work_dir", fmt.Sprintf("%s/work", baseDir))
	viper.SetDefault("log_file", fmt.Sprintf("%s/log.txt", baseDir))
	viper.SetDefault("db_file", fmt.Sprintf("%s/database.db", baseDir))
}

func Get(name string) string {
	return viper.GetString(name)
}

func GetBool(name string) bool {
	return viper.GetBool(name)
}

func GetInt(name string) int {
	return viper.GetInt(name)
}

func Read() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config.Read() err:%s", err))
	}
}
