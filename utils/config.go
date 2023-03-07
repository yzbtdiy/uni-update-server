package utils

import (
	"log"
	"os"

	"github.com/yzbtdiy/uni-update-server/models"

	"gopkg.in/yaml.v3"
)

// 读取 yaml 文件
func ReadYaml[T models.YamlParse](path string, data T) T {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Read file error: ", err)
	}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("Yaml unmarshal error:  ", err)
	}
	return data
}

func InitConfig() {
	if isExist("./config") {
		if isDir("./config") {
			if isExist("./config/config.yaml") {
				return
			} else {
				GenerateDefaultConfig()
			}
		} else {
			log.Fatal("config is not a folder")
		}
	} else {
		os.Mkdir("./config", 0755)
		GenerateDefaultConfig()
	}
}

func GenerateDefaultConfig() {
	serverConfig := models.ServerConfig{
		Ip:   "0.0.0.0",
		Port: "88",
	}
	clientConfig := models.ClientConfig{
		UpdateId: 1,
	}
	defaultConfig := models.Config{
		Server: serverConfig,
		Client: clientConfig,
	}

	config, err := yaml.Marshal(defaultConfig)
	if err != nil {
		panic(err)
	}
	os.WriteFile("./config/config.yaml", config, 0644)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		log.Println(err)
		return false
	}
	return true
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
