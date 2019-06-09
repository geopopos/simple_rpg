// Package config will create the configuration object
package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Banner is our logo
const Banner = `
__   __  __   __  ______    ______    ___   _______  _______  __    _  _______   
|  | |  ||  | |  ||    _ |  |    _ |  |   | |       ||   _   ||  |  | ||       |  
|  |_|  ||  | |  ||   | ||  |   | ||  |   | |       ||  |_|  ||   |_| ||    ___|  
|       ||  |_|  ||   |_||_ |   |_||_ |   | |       ||       ||       ||   |___   
|       ||       ||    __  ||    __  ||   | |      _||       ||  _    ||    ___|  
|   _   ||       ||   |  | ||   |  | ||   | |     |_ |   _   || | |   ||   |___   
|__| |__||_______||___|  |_||___|  |_||___| |_______||__| |__||_|  |__||_______|                                                                                  
`

// Consts for DB types
const (
	DBTypeMemory    = "memory"
	DBTypeCassandra = "cassandra"
)

// Configuration is the root struct for config type
type Configuration struct {
	Server     ServiceConfiguration
	Subscriber SubscriberConfiguration
	Database   DatabaseConfiguration
}

// ServiceConfiguration configuration of service
type ServiceConfiguration struct {
	Name string
}

// SubscriberConfiguration is the config for the subscriber
type SubscriberConfiguration struct {
	Topic string
}

// DatabaseConfiguration config for mongodb
type DatabaseConfiguration struct {
	Type  string
	Hosts string
	Port  int
}

// NewConfig Create a new configuration
func NewConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	viper := viper.New()
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	configPath := getEnv("API_CONFIG_DIR", "/etc/player-service/")
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		if strings.Contains(err.Error(), "Not Found in") {
			return viper, fmt.Errorf("Config not found")
		}
		return nil, err
	}

	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })
	return viper, nil
}

func getEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
