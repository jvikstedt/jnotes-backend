package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Required driver
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// DB database connection
var DB *sqlx.DB

// Config Contains database configurations for different environments
type Config map[string]map[string]string

// Setup database connection
func Setup(configPath string) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}
	DB, err = sqlx.Connect(config[env]["driver"], config[env]["open"])
	if err != nil {
		panic(err)
	}
}
