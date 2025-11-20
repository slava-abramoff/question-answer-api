# Question Answer API

Простой API-сервис для работы с вопросами и ответами на Go + PostgreSQL.

## Требования

- Docker и Docker Compose
- Go

---

## Структура проекта

```
.
├── application/ # Исходный код приложения
├── migrations/ # SQL миграции для базы данных (goose)
├── docker-compose.yml # Docker Compose для запуска приложения и БД
├── Dockerfile # Dockerfile для сборки приложения
├── go.mod
└── go.sum
```

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

5. **Остановка приложения**

```bash
docker compose down
```

---

## Вопросы

### Создать вопрос
**POST** `/questions`

**Body (JSON):**
```json
{
  "Text": "Откуда растут цветы?"
}
```

**Response (JSON):**
```json
{
  "ID": 1,
  "Text": "Откуда растут цветы?",
  "CreatedAt": "2025-11-20T09:41:55Z",
  "UpdatedAt": "2025-11-20T09:41:55Z"
}
```

### Получить вопрос по ID
**GET** `/questions/:id`

**Response (JSON):**
```json
{
  "ID": 1,
  "Text": "Откуда растут цветы?",
  "CreatedAt": "2025-11-20T09:41:55Z",
  "UpdatedAt": "2025-11-20T09:41:55Z"
}
```

### Получить все вопросы
**GET** `/questions`

```json
[
  {
    "ID": 1,
    "Text": "Откуда растут цветы?",
    "CreatedAt": "2025-11-20T09:41:55Z",
    "UpdatedAt": "2025-11-20T09:41:55Z"
  },
  {
    "ID": 2,
    "Text": "Как растут деревья?",
    "CreatedAt": "2025-11-20T09:45:10Z",
    "UpdatedAt": "2025-11-20T09:45:10Z"
  }
]
```

### Обновить вопрос
**PUT** `/questions/:id`

```json
{
  "Text": "Откуда растут цветы?"
}
```

**Response (JSON):**
```json
{
  "ID": 1,
  "Text": "Откуда растут цветы?",
  "CreatedAt": "2025-11-20T09:41:55Z",
  "UpdatedAt": "2025-11-20T09:41:55Z"
}
```

### Удалить вопрос
**DELETE** `/questions/:id`

**Response:**
```
204 No Content
```

## Ответы


### Создать ответ к вопросу
**POST** `/questions/:id/answers`

**Body (JSON):**
```json
{
  "UserID": "acde070d-8c4c-4f0d-9d8a-162843c10333",
  "Text": "Цветы растут из семян."
}
```

**Response (JSON):**
```json
{
  "ID": 1,
  "QuestionID": 1,
  "UserID": "acde070d-8c4c-4f0d-9d8a-162843c10333",
  "Text": "Цветы растут из семян.",
  "CreatedAt": "2025-11-20T09:50:00Z",
  "UpdatedAt": "2025-11-20T09:50:00Z"
}
```

### Получить ответ по ID
**GET** `/answers/:id`

**Response (JSON):**
```json
{
  "ID": 1,
  "QuestionID": 1,
  "UserID": "acde070d-8c4c-4f0d-9d8a-162843c10333",
  "Text": "Цветы растут из семян.",
  "CreatedAt": "2025-11-20T09:50:00Z",
  "UpdatedAt": "2025-11-20T09:50:00Z"
}
```

### Удалить ответ
**DELETE** `/answers/:id`

**Response:**
```
204 No Content
```

