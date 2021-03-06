package fileCache

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	SavePath string
	ExtName  string
}

var config Config

func (c Config) Init(configPath string) {
	file, _ := os.Open(configPath)
	var _config Config
	detail, _ := ioutil.ReadAll(file)
	err := json.Unmarshal(detail, &_config)

	if err != nil {
		config.SavePath = "/tmp/"
		config.ExtName = ".gofc"
	} else {
		config.SavePath = string(_config.SavePath)
		config.ExtName = string(_config.ExtName)
	}
}
