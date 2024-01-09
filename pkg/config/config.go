package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port      int    `env:"PORT" env-default:"443"`
	Host      string `env:"HOST" env-default:"localhost"`
	Env       string `env:"ENV" env-default:"development"`
	AdminUser struct {
		Email    string `env:"ADMIN_EMAIL" env-default:"admin"`
		Password string `env:"ADMIN_PWD" env-default:"admin"`
	}
}

type LoggerConfig struct {
	LogLevel    string `env:"LOG_LEVEL" env-default:"development"`
	LogFilePath string `env:"LOG_FILE" env-default:"./tmp/log.log"`
	MaxSize     int    `env:"LOG_MAXSIZE" env-default:"100"`
	MaxBackups  int    `env:"LOG_MAXBACKUP" env-default:"3"`
	MaxAge      int    `env:"LOG_MAXAGE" env-default:"24"`
	Compress    bool   `env:"LOG_COMPRESS" env-default:"true"`
}

type MySQLConfig struct {
	Driver string `env:"MYSQL_DRIVER" env-default:"mysql"`
	DSN    string `env:"MYSQL_DSN" env-default:"root:486464@tcp(localhost:3306)/neuronews?parseTime=true"`
}

type TLSConfig struct {
	KeyPath  string `env:"TLS_KEY_PATH" env-default:"./tls/key.pem"`
	CertPath string `env:"TLS_CERT_PATH" env-default:"./tls/cert.pem"`
}

type SessionConfig struct {
	Secret string `env:"SESSION_SECRET" env-default:"s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge"`
}

type SMTPConfig struct {
	Host     string `env:"SMTP_HOST"`
	Port     int    `env:"SMTP_PORT"`
	Username string `env:"SMTP_USERNAME"`
	Password string `env:"SMTP_PASSWORD"`
	Sender   string `env:"SMTP_SENDER"`
}

type Config struct {
	App     AppConfig
	Logger  LoggerConfig
	MySQL   MySQLConfig
	Session SessionConfig
	TLS     TLSConfig
	SMTP    SMTPConfig
}

// Singleton pattern
// var cfg *Config
// var once sync.Once

// New return instance of config
func New() (*Config, error) {
	// once.Do(func() {
	// instance = &Config{}

	// if err := cleanenv.ReadEnv(instance); err != nil {
	// 	helpText := "Neuro news"
	// 	help, _ := cleanenv.GetDescription(instance, &helpText)
	// 	log.Print(help)
	// 	log.Fatal(err)
	// }
	// })

	// var cfg *Config
	// load environments from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	// Загрузка переменных окружения из .env файла
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
