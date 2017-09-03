package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

var GOPATH string
var baseDir string

func init() {
	viper.BindEnv("GOPATH")
	GOPATH = viper.GetString("GOPATH")
	var err error
	baseDir, err = os.Getwd()
	if err != nil {
		panic(fmt.Errorf("config.init() err:%s", err))
	}
	//fmt.Printf("baseDir:%s\n", baseDir)

	viper.AddConfigPath(baseDir)
	// this came up in testing.  this package can load from anywhere (pwd speaking)
	viper.AddConfigPath(path.Clean(fmt.Sprintf("%s/../", baseDir)))
	viper.SetConfigName("settings")
	read()

	//fmt.Printf("baseDir2:%s\n", baseDir)
	viper.SetDefault("template_dir", fmt.Sprintf("%s/template", baseDir))
	viper.SetDefault("static_dir", fmt.Sprintf("%s/static", baseDir))
	viper.SetDefault("work_dir", fmt.Sprintf("%s/work", baseDir))
	viper.SetDefault("log_file", fmt.Sprintf("%s/log.txt", baseDir))
	viper.SetDefault("db_file", fmt.Sprintf("%s/database.db", baseDir))
	viper.SetDefault("pongo2_debug", true)
}

func settingsDir() string {
	return "asd"
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

func read() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config.Read() err:%s", err))
	}
	//fmt.Printf("config:%s\n", viper.ConfigFileUsed())
	//fmt.Printf("config:%s\n", path.Dir(viper.ConfigFileUsed()))
	baseDir = path.Dir(viper.ConfigFileUsed())
	//fmt.Printf("baseDir3:%s\n", baseDir)
}
