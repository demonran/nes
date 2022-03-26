package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Conf struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (db *Database) Source() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Name)
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "cmd/server/app.yml", "--conf app.yaml")
	flag.Parse()
}

func GetConfig(conf interface{}) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	buf = []byte(os.ExpandEnv(string(buf)))
	if err = yaml.Unmarshal(buf, conf); err != nil {
		panic(err)
	}

}
