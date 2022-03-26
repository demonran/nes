package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Conf struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Host string `yaml:"host"`
	Name string `yaml:"name"`
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "cmd/server/app.yml", "--conf app.yaml")
	flag.Parse()
}

func GetConfig(conf *Conf) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	buf = []byte(os.ExpandEnv(string(buf)))
	if err = yaml.Unmarshal(buf, conf); err != nil {
		panic(err)
	}

}
