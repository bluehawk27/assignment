package config

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func getProjectPath() string {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Println("Warning, cannot get current path")
		return ""
	}
	// Traverse back from current directory until service base dir is reached and add to config path
	for !strings.HasSuffix(dir, "assignment") && dir != "/" {
		dir, err = filepath.Abs(dir + "/..")
		if err != nil {
			break
		}
	}
	return dir
}

func getHomePath() string {
	home, _ := homedir.Dir()
	return home
}

func Init() {

	// Find home directory.
	viper.SetEnvPrefix("assignment")
	viper.BindEnv("configFile")
	viper.BindEnv("configPath")

	viper.SetDefault("configFile", "config")
	viper.SetDefault("configPath", "/tmp")

	viper.SetDefault("logging.level", "DEBUG")
	viper.SetDefault("logging.errorLogFile", "error.log")

	// Search config in home directory with name ".rs-collabs-brand-test" (without extension).
	viper.AddConfigPath(getHomePath())
	viper.AddConfigPath(getProjectPath())
	viper.AddConfigPath(viper.GetString("configPath"))
	viper.SetConfigName(viper.GetString("configFile"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error using config file:", viper.ConfigFileUsed())
		log.Println(err.Error())
	}

}

// GetCCacheConfig : Gets the in Memory Configuration
func GetCCacheConfig() CCache {
	c := CCache{
		Capacity: viper.GetInt64("ccache.capacity"),
		Expiry:   viper.GetInt64("ccache.expiry"),
	}

	return c
}

// GetRedisConfig : Gets the Redis Configuration
func GetRedisConfig() Redis {
	r := Redis{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.DB"),
		Expiry:   viper.GetInt64("redis.expiry"),
	}

	return r
}

// GetProxyConfig : Gets the Service Configuration
func GetProxyConfig() Proxy {
	p := Proxy{
		Host: viper.GetString("proxy.host"),
		Port: viper.GetString("proxy.port"),
	}

	return p
}

// GetProxyConnectionString : Sets the Connection string for the project
func GetProxyConnectionString() string {
	p := GetProxyConfig()
	address := fmt.Sprintf("%s:%s", p.Host, p.Port)

	return address
}

// GetRedisConnectionString : Sets the Redis db Connection string
func GetRedisConnectionString() string {
	r := GetRedisConfig()
	address := fmt.Sprintf("%s:%s", r.Host, r.Port)

	return address
}
