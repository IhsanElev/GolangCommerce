package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DbConfig  `yaml:"db"`
}

type DbConfig struct {
	Host           string           `yaml:"host"`
	Port           string           `yaml:"port"`
	User           string           `yaml:"user"`
	Password       string           `yaml:"password"`
	Name           string           `yaml:"name"`
	ConnectionPool DbConnectionPool `yaml:"connection_pool"`
}

// Connection Pool berguna untuk membatasi connection ke database, contoh misalnya postgres
// memiliki maks 100 connection jika melebihi dari itu akan mendapatkan error too many connection
// connection pool untuk mencegah hal itu
type DbConnectionPool struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnection     uint8 `yaml:"max_open_connection"`
	MaxLifeTimeConnection uint8 `yaml:"max_life_time_connection"`
	MaxIdleTimeConnection uint8 `yaml:"max_idle_time_connection"`
}
type AppConfig struct {
	Name       string           `yaml:"name"`
	Port       string           `yaml:"port"`
	Encryption EncryptionConfig `yaml:"encryption"`
}

type EncryptionConfig struct {
	Salt      uint8  `yaml:"salt"`
	JWTSecret string `yaml:"jwt_secret"`
}

var Cfg Config

func LoadConfig(filename string) (err error) {
	configByte, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	return yaml.Unmarshal(configByte, &Cfg)

}
