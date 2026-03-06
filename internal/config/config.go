package config

type Config struct {
	App   AppConfig      `yaml:"app"`
	Mysql PostgresConfig `yaml:"mysql"`
}

type AppConfig struct {
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
