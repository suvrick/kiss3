package db

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Timeout  int
}
