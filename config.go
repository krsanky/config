package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var GOPATH string
var configPath string

func init() {
	fmt.Println("init config")

	viper.BindEnv("GOPATH")
	GOPATH = viper.GetString("GOPATH")

	viper.SetDefault("template_dir", fmt.Sprintf("%s/src/oldcode.org/webplay/template", GOPATH))
	viper.SetDefault("static_dir", fmt.Sprintf("%s/src/oldcode.org/webplay/static", GOPATH))
	viper.SetDefault("work_dir", fmt.Sprintf("%s/src/oldcode.org/webplay/work", GOPATH))
	viper.SetDefault("log_file", fmt.Sprintf("%s/src/oldcode.org/webplay/log.txt", GOPATH))
	viper.SetDefault("db_file", fmt.Sprintf("%s/src/oldcode.org/webplay/database.db", GOPATH))

	//read_settings()
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

func SetDir(path string) {
	configPath = path
}

func Read() {
	//_, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//_, filename, _, ok := runtime.Caller(1)
	//if !ok {
	//	panic("error finding settings file")
	//}
	//fmt.Println(filename)
	//filepath := path.Dir(path.Join(filename, ".."))
	//fmt.Println(filepath)
	viper.AddConfigPath(configPath)

	viper.SetConfigName("settings")

	err := viper.ReadInConfig()
	if err != nil {
		panic("error finding settings file 2")
	}
	template_dir := viper.GetString("template_dir")
	fmt.Printf("template_dir: %s\n", template_dir)

}
