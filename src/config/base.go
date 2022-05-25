package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	Params *Config
)

type Config struct {
	Port               int    `required:"true"`
	ServiceFrontend    string `required:"true"`
	MongoURI           string `required:"true"`
	CORSPermitted      string `required:"true"`
	AirbringrDomain    string `required:"true"`
	MongoDbName        string `required:"true"`
	JWTSignatureSecret string `required:"true"`
}

func New() {
	config := Config{}
	err := envconfig.Process("banner", &config)

	if err != nil {
		log.Fatal(err)
	}
	Params = &config
}
