package config

import (
	"fmt"
	"os"
)

// Config : gateway configuration
type Config struct {
	Host           string `json: "host"`
	Port           string `json: "port"`
	Env            string `json: "env"`
	TokenSecretKey string `json: "token_secret_key"`

	// gcloud storage
	GCStorageCredentials string `json: "gcstorage_credentials"`
	GCStorageBucketName  string `json: "gcstorage_bucket_name"`
	GCStorageBaseURL     string `json: "gcstorage_base_url"`

	MySQLConnURL string `json: "mysql_conn_url"`
}

var config *Config

func init() {
	env := os.Getenv("env")
	host := os.Getenv("host")
	port := os.Getenv("port")
	tokenSecretKey := os.Getenv("token_secret_key")

	gcStorageCredentials := os.Getenv("gcstorage_credentials")
	gcStorageBucketName := os.Getenv("gcstorage_bucket_name")
	gcStorageBaseURL := os.Getenv("gcstorage_base_url")

	mysqlConnURL := os.Getenv("mysql_conn_url")

	config = &Config{
		Host:                 host,
		Port:                 port,
		Env:                  env,
		TokenSecretKey:       tokenSecretKey,
		GCStorageBaseURL:     gcStorageBaseURL,
		GCStorageBucketName:  gcStorageBucketName,
		GCStorageCredentials: gcStorageCredentials,
		MySQLConnURL:         mysqlConnURL,
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
	fmt.Printf("TokenSecretKey:\t\t\t%s\n", conf.TokenSecretKey)

	fmt.Printf("GCStorageBaseURL:\t%s\n", conf.GCStorageBaseURL)
	fmt.Printf("GCStorageBucketName:\t%s\n", conf.GCStorageBucketName)
	fmt.Printf("GCStorageCredentials:\t%s\n", conf.GCStorageCredentials)

	fmt.Printf("MySQL Connection URL:\t%s\n", conf.MySQLConnURL)
}
