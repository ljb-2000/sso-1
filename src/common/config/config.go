package config

import(
	"sync"
	"github.com/astaxie/beego/logs"
)

var once sync.Once
var instance *Config

func Instance() *Config {
	once.Do(func() {
		instance = new(Config)
	})
	return instance
}

type Config struct {
	DbHost string `json:"DbHost"`
	DbUser string `json:"DbUser"`
	DbPass string `json:"DbPass"`
	DbName string `json:"DbName"`
}

func (conf *Config) Instance() {
	conf.DbHost = "10.151.31.125:3306"
	conf.DbUser = "root"
	conf.DbPass = "Alert1234!"
	conf.DbName = "configportal"
}
func (conf *Config) GetDbEndpoint() string {
	endpoint := conf.DbUser + ":" + conf.DbPass + "@tcp(" + conf.DbHost + ")/" + conf.DbName + DB_CONNECTION_SUFFIX
	logs.Info("Config MySQL Endpoint: endpoint=%s", endpoint)
	return endpoint
}

