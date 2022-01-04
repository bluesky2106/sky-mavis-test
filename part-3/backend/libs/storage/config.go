package storage

import (
	"fmt"
)

// Config : google storage conf
type Config struct {
	Credentials string
	BucketName  string
	BaseURL     string
}

// Print configurations
func (conf *Config) Print() {
	fmt.Println("\t\tGCloud Storage configurations:")
	fmt.Printf("\t\tCredentials:\t\t\t%s\n", conf.Credentials)
	fmt.Printf("\t\tBucketName:\t\t%s\n", conf.BucketName)
	fmt.Printf("\t\tBucketName:\t\t%s\n", conf.BaseURL)
}
