package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"sync"

	"github.com/calmlow/go-places/internal/types"
	"gopkg.in/yaml.v3"
)

var (
	lock         sync.Mutex
	cachedConfig LocalConfig
	configLock   sync.Mutex
	configLoaded bool
)

type LocalConfig struct {
	Places          []types.Place `yaml:"places"`
	ReferenceRepo   string        `yaml:"reference-repo"`
	BackgroundColor string        `yaml:"background-color"`
}

func ReadYamlConfigFile(params ...string) (LocalConfig, error) {
	configLock.Lock()
	defer configLock.Unlock()

	if configLoaded {
		log.Println("Return cached config")
		return cachedConfig, nil
	}

	log.Println("Remove me. Readme YAML from fs")
	var config LocalConfig
	configFile := getLocalConfigPath()
	if len(params) == 1 {
		configFile = params[0]
	} else if len(params) > 1 {
		return config, fmt.Errorf("call ReadYamlConfigFile with one param, the string config filename of choice")
	}

	fileBytes, err := readFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config yaml file: %v", err)
	}

	err = yaml.Unmarshal(fileBytes, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config yaml '%s' file: %v", getLocalConfigPath(), err)
	}

	cachedConfig = config
	configLoaded = true
	return config, nil
}

func readFile(filePath string) ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		if filepath.Ext(filePath) == ".yml" {
			filePath = filePath[:len(filePath)-4] + ".yaml"
			return os.ReadFile(filePath)
		}
		return bytes, err
	}

	return bytes, nil
}

func getLocalConfigPath() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home := os.Getenv("HOME")
		configHome = home + "/.config"
	}

	return configHome + "/reposelector/reposelector-config.yaml"
}
