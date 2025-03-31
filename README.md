```md
# URL Shortener

Этот проект представляет собой простой сервис для сокращения URL с возможностью отслеживания статистики кликов по сокращённым ссылкам. Он включает в себя:

- Сокращение длинных URL.
- Перенаправление на оригинальные ссылки.
- Отслеживание статистики кликов по сокращённым ссылкам.

## Установка

Для работы с проектом вам нужно иметь установленный Docker и Docker Compose.

### Клонировать репозиторий:

```bash
git clone https://github.com/yourusername/URL-Shortener.git
cd URL-Shortener
```

### Сборка и запуск с Docker

1. Для запуска проекта используйте Docker Compose:

```bash
docker-compose up -d
```

2. Это создаст и запустит три контейнера:
   - **PostgreSQL** для хранения данных.
   - **Redis** для кэширования.
   - **Приложение** с сервером для обработки запросов.

### Обновление базы данных

Если вы изменяли структуру базы данных, например, добавили новые поля в таблицу или индексы, запустите команду для перезапуска контейнера:

```bash
docker-compose down
docker-compose up -d
```

## Использование

После запуска сервиса на вашем локальном сервере будет доступен API для взаимодействия с сокращёнными URL.

### Эндпоинты:

#### 1. Сокращение URL

**Метод**: `POST`  
**URL**: `/shorten`  
**Тело запроса**:

```json
{
	"url": "https://example.com/long-url"
}
```

**Ответ**:

```json
{
	"short_url": "http://localhost:8080/EAaArVRs"
}
```

#### 2. Перенаправление на оригинальный URL

**Метод**: `GET`  
**URL**: `/r/{short_url}`  
Замените `{short_url}` на сокращённый URL.

Пример:

```bash
GET http://localhost:8080/r/EAaArVRs
```

Этот запрос перенаправит вас на оригинальный URL.

#### 3. Статистика кликов по URL

**Метод**: `GET`  
**URL**: `/stats/{short_url}`  
Замените `{short_url}` на сокращённый URL.

Пример:

```bash
GET http://localhost:8080/stats/EAaArVRs
```

**Ответ**:

```json
{
	"clicks": 123
}
```

## Структура проекта

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   └── url.go
│   ├── service/
│   │   └── url.go
│   └── storage/
│       └── db.go
├── migrations/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── init.sql
├── README.md
└── LICENSE
```

### Описание структуры:

- **cmd/api/main.go**: Основной файл для запуска приложения.
- **internal/handler/url.go**: Обработчики HTTP запросов (сокращение URL, перенаправление, статистика).
- **internal/service/url.go**: Логика обработки и генерации сокращённых URL.
- **internal/storage/db.go**: Взаимодействие с базой данных (PostgreSQL).
- **init.sql**: Скрипт для инициализации базы данных.

## Требования

- Docker
- Docker Compose

## Лицензия

Этот проект лицензирован под лицензией MIT.

```
MIT License

Copyright (c) 2025 scrameee

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
