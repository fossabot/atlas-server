package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

var appConfig *AppConfig

var once sync.Once

// GetAppConfig retrieves the current application configuration
func GetAppConfig() *AppConfig {
	once.Do(func() {
		loadAppConfig()
	})
	return appConfig
}

func loadAppConfig() {
	yamlFile, err := ioutil.ReadFile("app.yml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
