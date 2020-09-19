package main

import (
	"encoding/json"
	"github.com/phachon/go-logger"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeOut  int64
	WriteTimeOut int64
	Static       string
}

var config Configuration
var logger = go_logger.NewLogger()

func init() {
	setLogger()
	loadConfig()
}

func setLogger() {
	// TODO: set logger parameters here
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		logger.Error("config json read failed!")
		os.Exit(2)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		logger.Error("config json parse failed!")
		os.Exit(2)
	}
	logger.Debugf("config json parse: \t%s", config)
}
