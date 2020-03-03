package core

import (
	"os"
)

//Config struct
type Config struct {
	APIPrefix    string
	DBType       string
	DBConnection string
}

//Load from env variables
func (c *Config) Load() {
	c.APIPrefix = os.Getenv("API_PREFIX")
	c.DBType = os.Getenv("DB_TYPE")
	c.DBConnection = os.Getenv("DB_CONNECTION")
}
