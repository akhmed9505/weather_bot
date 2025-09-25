## Weather Bot

Telegram-бот, который моментально показывает текущую погоду в любом городе.

## Функциональность

- Установка города с помощью команды `/city [название]`
- Получение температуры установленного города командой `/weather`
- Данные пользователей и их городов хранятся в PostgreSQL

## Технологии

- Go
- PostgreSQL
- Telegram Bot API
- OpenWeatherMap API

## Запуск

Требуется PostgreSQL и переменные окружения в `.env`:

```
BOT_TOKEN=
OPENWEATHERAPI_KEY=

DATABASE_URL=

GOOSE_DRIVER=
GOOSE_DBSTRING=
GOOSE_MIGRATION_DIR=
```

```
go run main.go
```