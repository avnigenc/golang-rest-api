package shared

import "time"

type Config struct {
	ServerHost     	string `yaml:"ServerHost"`
	ServerPort     	string `yaml:"ServerPort"`
	DatabaseHost    string `yaml:"DatabaseHost"`
	DatabasePort	string `yaml:"DatabasePort"`
	JwtSecret		string `yaml:"JwtSecret"`
	JwtExpireTime	time.Duration `yaml:"JwtExpireTime"`
}
