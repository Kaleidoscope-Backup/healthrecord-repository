package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Cfg ...
var Cfg Configuration

// PubAMQP ...
var PubAMQP *AMQPPublisher

//AMQPUri ...
var AMQPUri string

//IPFSUri ...
var IPFSUri string

// Configuration ...
type Configuration struct {
	PUBLISHER struct {
		Type string `yaml:"AMQP"`
	} `yaml:"PUBLISHER"`
	AMQP struct {
		Host                       string `yaml:"Host"`
		Port                       int    `yaml:"Port"`
		Username                   string `yaml:"Username"`
		Password                   string `yaml:"Password"`
		HealthRecordTopic          string `yaml:"HealthRecordTopic"`
		HealthRecordKnowledgeTopic string `yaml:"HealthRecordKnowledgeTopic"`
		HealthRecordHashTopic      string `yaml:"HealthRecordHashTopic"`
		VHost                      string `yaml:"VHost"`
	} `yaml:"AMQP"`
	IPFS struct {
		Host string `yaml:"Host"`
		Port int    `yaml:"Port"`
	} `yaml:"IPFS"`
}

// InitializeConfig ...
func InitializeConfig() {
	var file string
	flag.StringVar(&file, "Config", "", "Application configuration as Yml")
	flag.Parse()

	if file == "" {
		fmt.Println("Config file missing. Program is terminated ...")
		return
	}

	fileByte, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := yaml.Unmarshal(fileByte, &Cfg); err != nil {
		log.Fatal(err.Error())
	}

	if Cfg.PUBLISHER.Type == "AMQP" {
		PubAMQP = &AMQPPublisher{}
	}

	if Cfg.AMQP.Port == 0 {
		AMQPUri = fmt.Sprintf("amqp://%s:%s@%s/%s", Cfg.AMQP.Username, Cfg.AMQP.Password, Cfg.AMQP.Host, Cfg.AMQP.VHost)
	} else {
		AMQPUri = fmt.Sprintf("amqp://%s:%s@%s:%d/", Cfg.AMQP.Username, Cfg.AMQP.Password, Cfg.AMQP.Host, Cfg.AMQP.Port)
	}

	IPFSUri = fmt.Sprintf("%s:%d", Cfg.IPFS.Host, Cfg.IPFS.Port)

}
