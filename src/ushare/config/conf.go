package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const ROOT_PTAH  = "/Users/tustar/Documents/iProject/Coding/API"

type config struct {
	// site
	Site struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	}
	// mysql
	MySql struct {
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Port     string `yaml:"port"`
	}
	// api
	Api struct {
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	}
	// token
	Token struct {
		Key string `yaml:"key"`
	}
	// build
	Build struct {
		Debug bool `yaml:"debug"`
	}
}

var Conf config

func init() {

	file, err := ioutil.ReadFile(ROOT_PTAH + "/src/ushare/config/conf.yaml")
	if err != nil {
		log.Fatalf("ReadFile: %v ", err)
	}

	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
