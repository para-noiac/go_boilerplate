package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "localhost",
			Port:     "5432",
			User:     "leekahwei",
			Password: "1",
			Dbname:   "test",
		},
	}
}
