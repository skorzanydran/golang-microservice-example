package env

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"os"
)

type Profile int64

const (
	prod Profile = iota
	test
	local
)

func (p Profile) ToString() string {
	switch p {
	case prod:
		return "prod"
	case test:
		return "prod"
	case local:
		return "local"
	}
	return "prod"
}

func GetProfile(profile string) Profile {
	switch profile {
	case "prod":
		return prod
	case "test":
		return test
	case "local":
		return local
	}
	return prod
}

func Init() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	profile := GetProfile(os.Getenv("PROFILE"))

	err := config.LoadFiles(fmt.Sprintf("internal/env/config-%s.yml", profile.ToString()))
	fmt.Println(fmt.Sprintf("Run with [%s] profile", profile.ToString()))
	if err != nil {
		panic(err)
	}
}

func Load(key string, dst interface{}) {
	err := config.BindStruct(key, &dst)
	if err != nil {
		panic(err)
	}
}
