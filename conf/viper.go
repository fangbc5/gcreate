package conf

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Mysql struct {
	Address  string
	Port     string
	Username string
	Password string
	Database string
}

type Dir struct {
	Tmpl        string
	ProjectName string
}

type Configuration struct {
	Mysql
	Dir
}

var model *Configuration
var lock sync.Mutex

func GetConfig() *Configuration {
	if model == nil {
		lock.Lock()
		if model == nil {
			model = &Configuration{}
		}
		lock.Unlock()
	}
	return model
}

func Init() *Configuration {
	GetConfig()
	viper.SetConfigType("yaml")
	viper.SetConfigFile("app.yaml")
	viper.SetConfigName("app")
	viper.AddConfigPath("./")
	viper.AddConfigPath("conf")
	viper.AddConfigPath("../conf")
	viper.AddConfigPath(os.Getenv("GCREATE_CONFIG_DIR"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("config load error: %v", err)
	}
	err = viper.Unmarshal(model)
	if err != nil {
		log.Panicf("config load error: %v", err)
	}
	if os.Getenv("GCREATE_TMPL_PATH") != "" {
		model.Tmpl = os.Getenv("GCREATE_TMPL_PATH")
	}
	if os.Getenv("GCREATE_OUT_PATH") != "" {
		model.Tmpl = os.Getenv("GCREATE_OUT_PATH")
	}
	return model
}
