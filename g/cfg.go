package g

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var config *Config
var rwLock sync.RWMutex

type Config struct {
	ChatGpt `json:"chatGpt" yaml:"chatGpt"`
	Server  `json:"server" yaml:"server"`
	Lark    `json:"lark" yaml:"lark"`
}

type ChatGpt struct {
	AuthToken string
}

type Server struct {
	Addr string
	Path string
}

type Lark struct {
	AppId             string
	AppSecret         string
	VerificationToken string
	EventEncryptKey   string
}

func GetConfig() *Config {
	rwLock.RLock()
	defer rwLock.RUnlock()
	return config
}

func SetConfig(c *Config) {
	rwLock.Lock()
	defer rwLock.Unlock()
	config = c
}

func ParseWithPath(path string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Println("error", err.Error())
		panic(err)
	}
	config = &c
}

func Parse() {
	cfg := flag.String("c", "./cfg.yaml", "-c xxx.yaml")
	flag.Parse()
	ParseWithPath(*cfg)
}
