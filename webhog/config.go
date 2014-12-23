package webhog

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"strings"
)

type configuration struct {
	mongodb   string
	ApiKey    string
	AwsKey    string
	AwsSecret string
	AwsRegion string
	bucket    string
}

var Config = new(configuration)

func LoadConfig() error {
	conf, err := yaml.ReadFile("webhog.yml")
	if err != nil {
		return err
	}

	Config.ApiKey, _ = conf.Get(getEnv() + ".api_key")
	Config.bucket, _ = conf.Get(getEnv() + ".bucket")

	key, _ := conf.Get(getEnv() + ".aws_key")
	secret, _ := conf.Get(getEnv() + ".aws_secret")
	region, _ := conf.Get(getEnv() + ".aws_region")
	mongo, _ := conf.Get(getEnv() + ".mongodb")

	Config.AwsKey = os.Getenv(key)
	Config.AwsSecret = os.Getenv(secret)
	Config.AwsRegion = os.Getenv(region)
	Config.mongodb = os.Getenv(mongo)

	return err
}

func getEnv() string {
	env := strings.ToLower(os.Getenv("MARTINI_ENV"))
	if env == "" || env == "development" {
		return "development"
	}

	return env
}
