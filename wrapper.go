package wrapper

import (
	"github.com/spf13/viper"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

type viperConfig struct {
	v *viper.Viper
}

func NewViperConfig() *viperConfig {
	v := viper.New()
	v.Set("loud", true)
	// if err := v.ReadInConfig(); err != nil {
	// 	log.Fatalf("Failed to read config: %v", err)
	// }

	return &viperConfig{v: v}
}

func (c *viperConfig) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *viperConfig) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *viperConfig) GetBool(key string) bool {
	return c.v.GetBool(key)
}
