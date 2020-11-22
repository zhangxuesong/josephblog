package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Config struct {
	Runmode string `yaml:"runmode"`

	Service struct {
		Port string `yaml:"port"`
	}

	Jwt struct {
		SignKey string        `yaml:"signKey"`
		TimeOut time.Duration `yaml:"timeOut"`
	}

	Mysql struct {
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"db_name"`
	} `yaml:"mysql"`

	Redis struct {
		Key  string `yaml:"key"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Auth string `yaml:"auth"`
		Db   int    `yaml:"db"`
	} `yaml:"redis"`

	Elastic struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"elastic"`
}

func init() {
	log.Println("加载配置文件开始。。。")
	basePath, err := os.Getwd()
	if err != nil {
		log.Panic("get base path error!!!")
	}

	fileName := filepath.Join(basePath, "config", "config.yaml")
	config, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic("can't read %s", fileName)
	}

	err = yaml.Unmarshal(config, &Config)
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("加载配置文件成功。。。")
}
