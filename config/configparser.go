/*
    ToDD Configuration

	Copyright 2016 Matt Oswalt. Use or modification of this
	source code is governed by the license provided here:
	https://github.com/Mierdin/todd/blob/master/LICENSE
*/

package config

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/gcfg.v1"
)

type API struct {
	Host string
	Port string
}

type Assets struct {
	IP   string
	Port string
}

type Comms struct {
	Plugin   string
	User     string
	Password string
	Host     string
	Port     string
}

type DB struct {
	Host         string
	Port         string
	Plugin       string
	DatabaseName string
}

type TSDB struct {
	Host         string
	Port         string
	Plugin       string
	DatabaseName string
}

type Testing struct {
	Timeout int // seconds
}

type Grouping struct {
	Interval int // seconds
}

type LocalResources struct {
	DefaultInterface string
	OptDir           string
	IPAddrOverride   string
}

type Config struct {
	API            API
	Assets         Assets
	Comms          Comms
	DB             DB
	TSDB           TSDB
	Testing        Testing
	Grouping       Grouping
	LocalResources LocalResources
}

func GetConfig(cfgpath string) Config {
	var cfg Config

	err := gcfg.ReadFileInto(&cfg, cfgpath)
	if err != nil {
		log.Error("Error retrieving configuration")
		panic("Error retrieving configuration")
	}

	return cfg
}
