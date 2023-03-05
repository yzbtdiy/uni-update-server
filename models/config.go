package models

// 配置文件结构体
type Config struct {
	Server ServerConfig `yaml:"server"`
	Client ClientConfig `yaml:"client"`
}

type ServerConfig struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type ClientConfig struct {
	UpdateId int `yaml:"update_id"`
}
