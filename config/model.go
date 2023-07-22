package config

// IConfig is an interface that helps you interact with the config module
type IConfig interface {
	Get() *Config
}

// IConfigModel used as the instance to the IConfig Interface
type IConfigModel struct {
	model *Config
}

// Config is a model that is used to pass the configuration through out the project
type Config struct {
	AppVersion string `yaml:"appVersion"`
	Server     Server `yaml:"server"`
	DB         DB     `yaml:"db"`
}

// Server contains server related configurations
type Server struct {
	HTTP HTTP `yaml:"http"`
}

// HTTP contains http related configurations
type HTTP struct {
	Address string `yaml:"address"`
}

type DB struct {
	Tables []Table `yaml:"tables"`
}

type Table struct {
	Schema string `yaml:"schema"`
	Name   string `yaml:"name"`
}
