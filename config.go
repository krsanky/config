package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var GOPATH string
var baseDir string

func init() {
	viper.BindEnv("GOPATH")
	GOPATH = viper.GetString("GOPATH")
}

func Init(path string) {
	baseDir = fmt.Sprintf("%s/src/%s", GOPATH, path)
	fmt.Printf("baseDir:%s\n", baseDir)
	Read()
}

func InitWeb(path string) {
	Init(path)
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
	viper.AddConfigPath(baseDir)
	viper.SetConfigName("settings")
	err := viper.ReadInConfig()
	if err != nil {
		panic("error finding settings file 2")
	}
}
