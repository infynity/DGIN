package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//系统配置
var AppConfig *Config

type Config struct {
	Env      string `json:"Env"`      //环境配置
	Port     string `json:"Port"`     //程序监听端口
	MySqlDsn string `json:"MySqlDsn"` //数据库连接字符串
}

func InitConfig() {
	filename := "config/config.json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("InitConfig ReadFile Failed: %v", err)
		panic(err)
	}

	conf := Config{}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Println("InitConfig json.Unmarshal Failed: %v", err)
		panic(err)
	}

	AppConfig = &conf
}
