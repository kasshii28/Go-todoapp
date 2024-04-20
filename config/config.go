package config

type ConfigList struct {
	Post      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList