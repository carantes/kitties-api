package core

import (
	"os"
)

//Config struct
type Config struct {
	APIPrefix string
}

//Load from env variables
func (c *Config) Load() {
	c.APIPrefix = os.Getenv("API_PREFIX")
}
