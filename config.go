package main

import (
	"gopkg.in/yaml.v2"
	"flag"
	"io/ioutil"
	"time"
//	"github.com/gdexlab/go-render/render"
)

type Config struct {
	Listen 			string 				`yaml:"listen"`
	Prefix 			string 				`yaml:"prefix"`
	Forwarders 		map[string]string 	`yaml:"forwarders"`
	Default    		string 				`yaml:"default"`
	Static 			map[string]string 	`yaml:"static"`
	Cache 			struct {
		ExpTime 	time.Duration 		`yaml:"expiration"`
		PurgeTime 	time.Duration  		`yaml:"purge"`
	} 									`yaml:"cache"`
	LogLevel        string 				`yaml:"log-level"`
}

func InitConfig() (Config, error) {
	fileName := flag.String("file", "config.yml", "config filename")
	flag.Parse()

	Configs, err := parseFile(*fileName)
	if err != nil {
		return Config{}, err
	}

	return *Configs, nil
}

func parseFile(filePath string) (*Config, error) {
	cfg := new(Config)
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	cfg.Cache.ExpTime = 0
	cfg.Cache.PurgeTime = 0
	cfg.LogLevel = "info"
	if err := yaml.UnmarshalStrict(body, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validateForwarders() {

}
