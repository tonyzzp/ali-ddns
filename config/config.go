package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ConfigAli struct {
	Region    string `yaml:"region"`
	KeyId     string `yaml:"key_id"`
	KeySecret string `yaml:"key_secret"`
}

type ConfigCloudFlare struct {
	Token string `yaml:"token"`
}

type appConfig struct {
	Ali        ConfigAli        `yaml:"ali"`
	CloudFlare ConfigCloudFlare `yaml:"cloudflare"`
}

var Config *appConfig

func findConfigFile(file string) string {
	if file != "" {
		return file
	}

	// 查找工作目录
	name := "xddns-config.yaml"
	fi, e := os.Stat(name)
	if e == nil && fi != nil && !fi.IsDir() {
		return name
	}

	// 查找exe目录
	dir, e := os.Executable()
	if e != nil {
		return name
	}
	file = filepath.Join(filepath.Dir(dir), name)
	return file
}

func Init(file string) error {
	file = findConfigFile(file)
	Config = &appConfig{}
	r, e := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if e != nil {
		return e
	}
	e = yaml.NewDecoder(r).Decode(Config)
	r.Close()
	log.Println(Config)
	return e
}
