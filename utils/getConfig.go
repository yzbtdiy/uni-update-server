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
