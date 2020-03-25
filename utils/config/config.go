package config

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config interface {
	Get(key string) interface{}
	GetString(key string) string
	GetStringOrDefault(key string, defaultVal string) string
	MustGetString(key string) string
	GetBool(key string) bool
	GetBoolOrDefault(key string, defaultVal bool) bool
	GetInt(key string) int
	GetInt64(key string) int64
	GetIntOrDefault(key string, defaultVal int) int
	GetDuration(key string) time.Duration
	GetDurationOrDefault(key string, defaultVal time.Duration) time.Duration
	GetStringSlice(key string) []string
	Unmarshal(key string, val interface{}) error
	TryUnmarshal(key string, val interface{}) bool
	MustUnmarshal(key string, val interface{})
	GetStringMapString(key string) map[string]string
	GetFloat64(key string) float64
}

var (
	config Config
)

func init() {
	if config != nil {
		return
	}

	configDir := constants.CONFIG_PATH
	viper.AddConfigPath(configDir)
	viper.SetConfigName("default")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	config = &viperConfig{}
	c.config = config
}

func GetConfig() Config {
	return config
}
