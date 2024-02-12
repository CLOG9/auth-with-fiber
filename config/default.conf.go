package config

type DefaultConfig struct {
	ApiVersion string
}

type RouteEndpoints struct {
	Home     string
	Login    string
	Register string
	Admin    string
}

var Defaults DefaultConfig = DefaultConfig{
	ApiVersion: "/v1",
}
var RouteEndpts RouteEndpoints = RouteEndpoints{
	Home:     "/",
	Login:    "/login",
	Register: "/register",
	Admin:    "/admin",
}
