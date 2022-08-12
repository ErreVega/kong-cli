package config

import "fmt"

type Configuration struct {
	KONG_HOST string
	KONG_PORT int
}

func (c Configuration) GetUrl() string {
	return fmt.Sprintf("http://%s:%v", c.KONG_HOST, c.KONG_PORT)
}

var Config = Configuration{
	KONG_HOST: "localhost",
	KONG_PORT: 8001,
}
