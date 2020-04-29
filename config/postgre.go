package config

type PostgreConfig struct {
	Host string `json:"host"`
	Database string `json:"database"`
	User string`json:"user"`
	Password string `json:"password"`
	Port string `json:"port"`
}
