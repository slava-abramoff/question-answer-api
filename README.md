# Question Answer API

Простой API-сервис для работы с вопросами и ответами на Go + PostgreSQL.

## Требования

- Docker и Docker Compose
- Go (если будешь запускать миграции локально)

---

## Структура проекта

.
├── application/ # Исходный код приложения
├── migrations/ # SQL миграции для базы данных (goose)
├── docker-compose.yml # Docker Compose для запуска приложения и БД
├── Dockerfile # Dockerfile для сборки приложения
├── go.mod
└── go.sum

## Быстрый старт с Docker

1. **Запускаем PostgreSQL**

```bash
docker compose up -d db
```

2. **Применяем миграции**

```bash
goose -dir ./migrations postgres "host=localhost user=app password=app dbname=app sslmode=disable" up
```

3. **Собираем приложение**

```bash
docker compose build app
```

4. **Запускаем приложение**

```bash
docker compose up app
```

4. **Остановка приложения**

```bash
docker compose down
```