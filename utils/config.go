package utils

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Configuration struct {
	KONG_HOST string `mapstructure:"KONG_HOST" yaml:"KONG_HOST"`
	KONG_PORT int    `mapstructure:"KONG_PORT" yaml:"KONG_PORT"`
}

func (c Configuration) GetUrl() string {
	return fmt.Sprintf("http://%s:%v", viper.Get("KONG_HOST"), viper.Get("KONG_PORT"))
}

func (c Configuration) GetConfig() Configuration {
	return c
}

func (c Configuration) SetConfigProperty(property string, value string) {
	r := reflect.ValueOf(&c)
	f := reflect.Indirect(r).FieldByName(property)

	switch f.Kind() {
	case reflect.Int:
		intVal, err := strconv.ParseInt(value, 0, 64)
		PrintErr(err)
		f.SetInt(intVal)
	case reflect.String:
		f.SetString(value)
	// case reflect.Bool:
	// 	value = strconv.FormatBool(f.Bool())
	// case reflect.Struct, reflect.Interface, reflect.Array, reflect.Slice:
	// 	value = fmt.Sprintf("%v", f)
	case reflect.Invalid:
	}

	yamlData, err := yaml.Marshal(&c)
	PrintErr(err)
	err = ioutil.WriteFile("./config/config.yaml", yamlData, 0644)
	PrintErr(err)
	Config, err = LoadConfig("./config")
	PrintErr(err)
}
func (c Configuration) SetKONG_PORT(value int) {
	c.KONG_PORT = value

	yamlData, err := yaml.Marshal(&c)
	PrintErr(err)
	err = ioutil.WriteFile("./config/config.yaml", yamlData, 0644)
	PrintErr(err)
	Config, err = LoadConfig("./config")
	PrintErr(err)
}

func (c Configuration) SetKONG_HOST(value string) {
	c.KONG_HOST = value

	yamlData, err := yaml.Marshal(&c)
	PrintErr(err)
	err = ioutil.WriteFile("./config/config.yaml", yamlData, 0644)
	PrintErr(err)
	Config, err = LoadConfig("./config")
	PrintErr(err)
}

// var Config = Configuration{
// 	KONG_HOST: "localhost",
// 	// KONG_PORT: 9569,
// 	KONG_PORT: 8001,
// }

func LoadConfig(path string) (config Configuration, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.BindEnv("KONG_PORT")
	viper.BindEnv("KONG_HOST")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}

var Config, err = LoadConfig("./config")
