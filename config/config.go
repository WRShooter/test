package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm"
	"log"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// config.toml 的数据结构
type Config1 struct {
	DbType         string
	DbHost         string
	DbUserName     string
	DbUserPassword string
	DbName         string
	DbCharset      string
}

func (c *Config1) GetConfig1() {
	if _, err := toml.DecodeFile("./conf/config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

type DBConfig struct {
	Dialect  string
	Host     string
	Username string
	Password string
	Name     string
	Charset  string
}
type Config struct {
	DB *DBConfig
}

func GetConfig() *Config {
	var cfg1 = Config1{}
	cfg1.GetConfig1()
	return &Config{
		DB: &DBConfig{
			Dialect:  cfg1.DbType,
			Host:     cfg1.DbHost,
			Username: cfg1.DbUserName,
			Password: cfg1.DbUserPassword,
			Name:     cfg1.DbName,
			Charset:  cfg1.DbCharset,
		},
	}
}

func ReturnUrl() string {
	var dbCfg DBConfig = *((*GetConfig()).DB)
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.Charset)
	return dbURI
}
