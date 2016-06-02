package config

import (
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

// TomlConf Our app config
type TomlConf struct {
	Server struct {
		Port uint
	}

	App struct {
		Secret      string
		TokenSecret string
		BcryptCost  int
		TemplateDir string
		StaticDirs  []string
	}

	Database struct {
		Type     string
		Server   string
		Port     uint
		Username string
		Password string
		Database string
	}
}

// LoadConfig Load our config file
func LoadConfig(confPath ...string) TomlConf {
	var confFile string

	if len(confPath) == 0 {
		confFile = "config/default.toml"
	} else {
		confFile = confPath[0]
	}

	f, err := os.Open(confFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var config TomlConf
	if err := toml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}

	return config
}
