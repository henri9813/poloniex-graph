package main

import (
	"log"
	"os"

	influxdb "github.com/influxdata/influxdb1-client/v2"
)

var influxClient influxdb.Client
var influxDatabase string

func init() {
	var influxURL, influxUsername, influxPassword string

	if influxURL = os.Getenv("INFLUXDB_URL"); influxURL == "" {
		log.Fatal("environment variable: INFLUXDB_URL is not defined")
	}

	if influxDatabase = os.Getenv("INFLUXDB_DB"); influxDatabase == "" {
		log.Fatal("environment variable: INFLUXDB_DB is not defined")
	}

	if influxUsername = os.Getenv("INFLUXDB_USER"); influxUsername == "" {
		log.Fatal("environment variable: INFLUXDB_USER is not defined")
	}

	if influxPassword = os.Getenv("INFLUXDB_USER_PASSWORD"); influxPassword == "" {
		log.Fatal("environment variable: INFLUXDB_USER_PASSWORD is not defined")
	}

	c, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr:     influxURL,
		Username: influxUsername,
		Password: influxPassword,
	})

	if err != nil {
		log.Fatalf("Error creating InfluxDB Client: %s", err)
	}

	influxClient = c
}
