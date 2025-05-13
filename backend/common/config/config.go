package config

import (
	"flag"
	"github.com/bsthun/gut"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Environment      *uint8    `yaml:"environment" validate:"gte=1,lte=2"`
	WebListen        *string   `yaml:"webListen" validate:"required"`
	WebRoot          *string   `yaml:"webRoot" validate:"required"`
	Cors             []*string `yaml:"cors" validate:"required"`
	Secret           *string   `yaml:"secret" validate:"required"`
	FrontendUrl      *string   `yaml:"frontendUrl" validate:"required"`
	AuthEndpoint     *string   `yaml:"authEndpoint" validate:"required"`
	AuthClientId     *string   `yaml:"authClientId" validate:"required"`
	AuthClientSecret *string   `yaml:"authClientSecret" validate:"required"`
}

func Init() *Config {
	// * parse arguments
	path := flag.String("config", ".local/config.yml", "Path to config file")
	flag.Parse()

	// * declare struct
	config := new(Config)

	// * read config
	yml, err := os.ReadFile(*path)
	if err != nil {
		gut.Fatal("unable to read configuration file", err)
	}

	// * parse config
	if err := yaml.Unmarshal(yml, config); err != nil {
		gut.Fatal("unable to parse configuration file", err)
	}

	// * validate config
	if err := gut.Validate(config); err != nil {
		gut.Fatal("invalid configuration", err)
	}

	// * apply secret key
	var bytes = []byte(*config.Secret)
	if len(bytes) < 16 {
		for i := len(bytes); i < 16; i++ {
			bytes = append(bytes, 0)
		}
	}
	if err := gut.SetIdEncoderKey(bytes); err != nil {
		gut.Fatal("unable to set secret key", err)
	}

	return config
}
