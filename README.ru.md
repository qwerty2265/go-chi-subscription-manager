# go-chi-subscription-manager

## Возможности

- CRUD-операции с подписками
- Расчёт общей стоимости подписок за период
- Swagger-документация
- Docker-контейнеризация

## Структура проекта

```
.
├── app/                # Инициализация приложения и маршрутизация
├── cmd/server/         # Точка входа
├── docs/               # Swagger-документация
├── internal/           # Внутренние пакеты (common, subscription)
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
```

## Начало работы

### Необходимые компоненты

- [Go 1.24+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://www.docker.com/products/docker-desktop/)

### Конфигурация

Скопируйте `.env.example` в `.env` и заполните переменные окружения:

```sh
cp .env.example .env
```

### Запуск через Docker

Соберите и запустите приложение и базу данных PostgreSQL:

```sh
docker-compose up --build
```

API будет доступен по адресу `http://localhost:7070`.

### Локальный запуск

1. Запустите базу данных PostgreSQL
2. Настройте файл `.env`.
3. Запустите сервер:

```sh
go run /cmd/server/main.go
```

## Документация API

Swagger UI доступен по адресу:  
[http://localhost:7070/swagger/index.html](http://localhost:7070/swagger/index.html)

## Список эндпоинтов

- `GET /api/subscriptions?user-id={uuid}` - Список подписок пользователя
- `POST /api/subscriptions` - Создать подписку
- `GET /api/subscriptions/{id}` — Получить подписку по ID
- `PUT /api/subscriptions/{id}` — Обновить подписку
- `DELETE /api/subscriptions/{id}` — Удалить подписку
- `GET /api/subscriptions/total-price?user-id={uuid}&service-name={name}&from=MM-YYYY&to=MM-YYYY` — Рассчитать общую стоимость с фильтрами