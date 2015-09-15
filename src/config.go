package main

import (
	"fmt"
	"flag"
	"github.com/cloudfoundry-incubator/candiedyaml"
	"net/url"
	"io/ioutil"
)

var configStr string

type RedisConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}

var defaultRedisConfig = RedisConfig{
	Host: "127.0.0.1",
	Port: 6379,
	User: "",
	Password: "",
}

type Config struct {
	RedisServs []RedisConfig `yaml:"redisServs"`
}

var defaultConfig = Config{
	RedisServs: []RedisConfig{defaultRedisConfig},
}

func DefaultConfig() *Config {
	c := defaultConfig
	c.Process()
	return &c
}

func (c *Config) Process() {
	//do some process init
}

func (c *Config) RedisServices() []string {
	var redisServices []string
	for _, info := range c.RedisServs {
		uri := url.URL{
			Scheme: "redis",
			User: url.UserPassword(info.User, info.Password),
			Host: fmt.Sprintf("%s:%d", info.Host, info.Port),
		}
		redisServices = append(redisServices, uri.String())
	}
	return redisServices
}

func (c *Config) Initialize(configYaml []byte) error {
	c.RedisServs = []RedisConfig{}
	return candiedyaml.Unmarshal(configYaml, &c)
}

func InitConfigFromFile(path string) *Config{
	var c *Config = DefaultConfig()
	var e error
	b,e := ioutil.ReadFile(path)
	if e != nil {
		panic(e.Error())
	}
	e = c.Initialize(b)
	if e != nil {
		panic(e.Error())
	}

	//process
	c.Process()

	return c
}

func init(){
	flag.StringVar(&configStr, "c", "","This is config filepath")
	flag.Parse()
}

func main(){

	if configStr !="" {
		fmt.Println(configStr)
	}else {
		fmt.Println("filepath is not flag.")
	}

	//if config is package must import config
	c := DefaultConfig()
	if configStr != "" {
		c = InitConfigFromFile(configStr)
	}

	//[{192.168.172.128 6379 guest guest} {127.0.0.1 6379 redis password}]
	fmt.Println(c.RedisServs)
}
