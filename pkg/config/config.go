package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	LISTEN_TYPE_SOCK = "sock"
	LISTEN_TYPE_PORT = "port"
)

type AppConfig struct {
	Port      int    `env:"PORT" env-default:"8080"`
	Host      string `env:"HOST" env-default:"localhost"`
	Env       string `env:"ENV" env-default:"development"`
	AdminUser struct {
		Email    string `env:"ADMIN_EMAIL" env-default:"admin"`
		Password string `env:"ADMIN_PWD" env-default:"admin"`
	}
}

type LoggerConfig struct {
	LogFilePath string `env:"LOG_FILE" env-default:"./tmp/log.log"`
	MaxSize     int    `env:"LOG_MAXSIZE" env-default:"100"`
	MaxBackups  int    `env:"LOG_MAXBACKUP" env-default:"3"`
	MaxAge      int    `env:"LOG_MAXAGE" env-default:"24"`
	Compress    bool   `env:"LOG_COMPRESS" env-default:"true"`
}

type MySQLConfig struct {
	DSN string `env:"WEB_DB_DSN" env-default:"file_cloud:Todor1990@tcp(localhost:3306)/file_cloud?parseTime=true"`
}

type TlsConfig struct {
	KeyPath  string `env:"TLS_KEY_PATH" env-default:"./tls/key.pem"`
	CertPath string `env:"TLS_CERT_PATH" env-default:"./tls/cert.pem"`
}

type SessionConfig struct {
	Secret string `env:"SESSION_SECRET" env-default:"s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge"`
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Sender   string
}

type Config struct {
	App     AppConfig
	Logger  LoggerConfig
	MySQL   MySQLConfig
	Session SessionConfig
	TLS     TlsConfig
	SMTP    SMTPConfig
}

// Singleton pattern
var instance *Config
var once sync.Once

// Return instance of config
func New() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "File cloud"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}