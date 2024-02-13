package config

type DefaultConfig struct {
	ApiVersion string
}

var Defaults DefaultConfig = DefaultConfig{
	ApiVersion: "/v1",
}
