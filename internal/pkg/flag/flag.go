package flag

import (
	"flag"
	"os"

	"github.com/alekslesik/neuro-news/pkg/config"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

// flags init
func Init(config *config.Config) error {


	// create flagset
	flagSet := flag.NewFlagSet("flag", flag.ContinueOnError)

	// define flags in flagset
	flagSet.StringVar(&config.App.Env, "env", string(logger.DEVELOPMENT), "Environment (development|staging|production)")
	flagSet.IntVar(&config.App.Port, "port", 443, "API server port")
	flagSet.StringVar(&config.SMTP.Host, "smtp-host", "app.debugmail.io", "SMTP host")
	flagSet.IntVar(&config.SMTP.Port, "smtp-port", 25, "SMTP port")
	flagSet.StringVar(&config.SMTP.Username, "smtp-username", "d40e021c-f8d5-49af-a118-81f40f7b84b7", "SMTP username")
	flagSet.StringVar(&config.SMTP.Password, "smtp-password", "a8c960ed-d3ad-44e6-8461-37d40f15e569", "SMTP password")
	flagSet.StringVar(&config.SMTP.Sender, "smtp-sender", "alekslesik@gmail.com", "SMTP sender")

	// take transferred to application arguments using os.Args slice
	args := os.Args[1:]
	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	return nil
}

// Области для Улучшения

//     Жестко Закодированные Значения по Умолчанию: Флаги инициализируются со значением по умолчанию (например, SMTP-сервер, имя пользователя и пароль). В идеале эти значения следует задавать в конфигурационных файлах или переменных окружения, чтобы избежать жесткой привязки к конкретным значениям в коде.
//     Обработка Чувствительных Данных: Чувствительные данные, такие как имя пользователя и пароль SMTP, должны быть обрабатываться с особой осторожностью, возможно, через безопасные каналы (например, переменные окружения), а не через флаги командной строки.
//     Проверка Валидности Флагов: Добавление логики для проверки корректности введенных значений флагов поможет предотвратить ошибки на ранних этапах выполнения приложения.
//     Документация и Комментарии: Хорошей практикой является добавление комментариев к коду, особенно к публичным функциям и основным блокам логики. Это помогает другим разработчикам быстрее понять назначение кода.