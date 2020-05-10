package cfg

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"runtime"
)

func Load() error {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "pp"
	}
	log.Println("Environment:", env)

	relPath := getCurrentDir()
	viper.AddConfigPath(relPath)
	viper.SetConfigName(env)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("error loading config: ", err.Error())
	}
	return nil
}

func getCurrentDir() string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get config path")
	}
	return path.Dir(thisFile)
}
