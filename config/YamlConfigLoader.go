package config

import (
	"github.com/myste1tainn/gojenkins-cli/models"
	"github.com/myste1tainn/gojenkins-core/helpers"
	"gopkg.in/yaml.v2"
)

var YamlConfigFileName = "./gojenkins.yaml"

type YamlConfigLoader interface {
	Load() models.YamlConfig
}

type DefaultYamlConfigLoader struct {
	FileReader helpers.FileReader
}

func (this DefaultYamlConfigLoader) Load() models.YamlConfig {
	data := helpers.PanicData(this.FileReader.ReadFile("./" + YamlConfigFileName))
	config := models.YamlConfig{}
	helpers.Panic(yaml.Unmarshal(data, &config))
	return config
}

func NewDefaultYamlConfigLoader(reader helpers.FileReader) *DefaultYamlConfigLoader {
	return &DefaultYamlConfigLoader{FileReader: reader}
}
