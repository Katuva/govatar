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
		Server   string
		Port     uint
		Username string
		Password string
		Database string
	}
}

// LoadConfig Load our config file
func LoadConfig(confPath string) TomlConf {
	if confPath == "" {
		confPath = "config/default.toml"
	}

	f, err := os.Open(confPath)
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
