package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App *App `yaml:"app"`
	//Database map[string]xdb.Config
	//Redis    map[string]xrds.Config
}

type App struct {
	Name string   `yaml:"name"`
	Addr *AppAddr `yaml:"addr"`
}

type AppAddr struct {
	Grpc string `yaml:"grpc"`
	Http string `yaml:"http"`
}

func NewConfig(env string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(env)
	if err != nil {
		fmt.Println(err.Error())
	}
	var conf *Config
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("config.app: %#v\n", conf.App)
	
	return conf, nil
}
