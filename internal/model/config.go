package model

type Config struct {
	RestPort int
	DBConfig DBConfig
}

type DBConfig struct {
	Username   string
	Password   string
	Host       string
	Port       string
	Database   string
	ConnString string
}
