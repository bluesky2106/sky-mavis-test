package config

import (
	"fmt"
	"os"
)

// Config : gateway configuration
type Config struct {
	Host         string `json: "host"`
	Port         string `json: "port"`
	Env          string `json: "env"`
	MySQLConnURL string `json: "mysql_conn_url"`
}

var config *Config

func init() {
	env := os.Getenv("env")
	host := os.Getenv("host")
	port := os.Getenv("port")
	mysqlConnURL := os.Getenv("mysql_conn_url")

	config = &Config{
		Host:         host,
		Port:         port,
		Env:          env,
		MySQLConnURL: mysqlConnURL,
	}
}

// GetConfig :
func GetConfig() *Config {
	return config
}

// Print configurations
func (conf *Config) Print() {
	fmt.Println("------------ Gateway configurations --------------")
	fmt.Printf("Env:\t\t\t%s\n", conf.Env)
	fmt.Printf("Host:\t\t\t%s\n", conf.Host)
	fmt.Printf("Port:\t\t\t%s\n", conf.Port)
	fmt.Printf("MySQL Connection URL:\t%s\n", conf.MySQLConnURL)
}
