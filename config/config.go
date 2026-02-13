package config

type DBConfig struct {
	Host          string
	Port          int
	User          string
	Password      string
	DBName        string
	EnableSSLMode bool
}

type AppConfig struct {
	Version     string
	ServiceName string
	AppUrl      string
	HttpPort    int
	JwtSecret   string
	DB          *DBConfig
}
