package util

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Config struct
type Config struct {
	maps map[string]interface{}
}

// Get the value in the config.json
// Example:
// config := new(config.Config)
// config.Read("config/config.json")
// fmt.Println(config.Get("host"))
// fmt.Println(config.Get("headers.Accept"))
func (config *Config) Get(name string) interface{} {
	if config.maps == nil {
		return nil
	}

	keys := strings.Split(name, ".")
	length := len(keys)

	if length == 1 {
		return config.maps[name]
	}

	var value interface{}
	for i := 0; i < length; i++ {
		if i == 0 {
			value = config.maps[keys[i]]
		} else {
			if submap, ok := value.(map[string]interface{}); ok {
				value = submap[keys[i]]
			}
		}
	}
	return value
}

// Read the file, config.json from specific path
func (config *Config) Read(file string) {
	file, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &config.maps); err != nil {
		panic(err)
	}
}
