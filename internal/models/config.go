package models

type Config struct {
	Env  string `yaml:"env" env-default:"local"`
	Port string `yaml:"port"`
}
