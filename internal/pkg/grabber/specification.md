### Техническое задание:

#### Идентификация новостного сайта:

1. Необходимо определить новостной сайт, с которого будут браться новости.
2. Убедиться, что сайт предоставляет данные в формате, который можно парсить.

#### Парсинг новостей:

1. Написать код для извлечения списка новостей с сайта.
2. Выбрать последнюю новость из списка.
3. Извлечь заголовок и содержимое новости.

#### Взаимодействие с API для генерации изображений:

1. Найти подходящий API или сервис, который может генерировать изображения на основе текста.
2. Реализовать код для отправки заголовка новости API и получения ссылки на изображение.

#### Загрузка изображения:

1. Скачать изображение по полученной ссылке.
2. Убедиться, что изображение успешно загружено и доступно для дальнейшей обработки.

#### Сохранение новости в базу данных:

1. Создать структуру базы данных для хранения информации о новостях (например, заголовок, содержание, ссылка на изображение).
2. Написать код для добавления новости в базу данных.

#### Запланированное выполнение задач:

1. Настроить запуск процесса обновления новостей два раза в час.

### Этапы разработки:

#### Исследование и планирование:

1. Выбор новостного сайта.
2. Поиск подходящего API для генерации изображений.
3. Проектирование базы данных.

#### Написание кода:

1. Разработка кода для парсинга новостей.
2. Написание кода для взаимодействия с API генератора изображений.
3. Реализация функционала загрузки изображений.
4. Настройка взаимодействия с базой данных.

#### Тестирование:

1. Проведение модульного тестирования отдельных компонентов.
2. Тестирование взаимодействия между компонентами.
3. Проверка корректности сохранения данных в базе данных.

#### Интеграция и настройка расписания:

1. Интеграция всех компонентов в единое приложение.
2. Настройка запуска задач по расписанию (например, с использованием пакета time в Go).

### Инструменты и паттерны:

- **Go:** Язык программирования Go для написания приложения.
- **HTML Parsing:** Для парсинга HTML страниц можно использовать пакет golang.org/x/net/html.
- **HTTP Requests:** Для отправки HTTP запросов на сайт можно использовать стандартную библиотеку net/http.
- **API Requests:** Для взаимодействия с API генератора изображений можно использовать net/http или популярный пакет github.com/go-resty/resty.
- **База данных:** Для работы с базой данных в Go можно использовать стандартный пакет database/sql, а также популярные библиотеки-обертки, такие как github.com/jmoiron/sqlx или ORM-фреймворки типа github.com/jinzhu/gorm.
- **Паттерны:**
  - **Singleton (Одиночка):** Можно использовать паттерн Singleton для создания единственного экземпляра клиента HTTP для работы с API и базой данных.
  - **Scheduler (Планировщик):** Для запуска задач по расписанию можно использовать пакеты типа github.com/robfig/cron или стандартную библиотеку time.
