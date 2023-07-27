package config

type ServerConfType struct {
	HostName string `yaml:"hostname" json:"hostname"`
	Port     int    `yaml:"port" json:"port"`
	Env      string `yaml:"env" json:"env"` // release/debug/test
}

// Configuration mapping to app-config.yml
type AppConfig struct {
	Server ServerConfType `yaml:"server" json:"server"`
}

// app config instance
var APP_CONFIG AppConfig = AppConfig{
	// Set some default value here
	Server: ServerConfType{
		HostName: "127.0.0.1",
		Port:     80,
		Env:      "debug",
	},
}
