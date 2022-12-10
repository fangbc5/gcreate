package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

type Project struct {
	Name          string
	ModuleName    string
	InterfaceName string
}
type Mysql struct {
	Username string
	Password string
	Driver   string
	Url      string
	Schema   string
}

type Table struct {
	Names  []string
	Prefix string
}

type Dir struct {
	Tmpl string
	Out  string
}
type Configuration struct {
	Project
	Table
	Mysql
	Dir
}

var model *Configuration
var lock sync.Mutex

func MakeConfig() *Configuration {
	if model == nil {
		lock.Lock()
		if model == nil {
			model = &Configuration{}
		}
		lock.Unlock()
	}
	return model
}

func (c *Configuration) Load() *Configuration {
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
