# PickPoint Service

Сервис для управления пунктами выдачи заказов (ПВЗ) и приемкой товаров.

## Структура проекта
```
pickpoint/
│
├── cmd/
│   └── main.go                  # Точка входа в приложение
│
├── internal/
│   ├── api/                     # Слой HTTP API
│   │   ├── handler/             # Обработчики HTTP-запросов
│   │   │   ├── intake/
│   │   │   ├── pickpoint/
│   │   │   └── user/
│   │   └── router/              # Инициализация роутера (маршрутизация)
│
│   ├── app/                     # Инициализация зависимостей и запуск приложения
│   │   ├── app.go
│   │   └── service_provider.go
│
│   ├── client/                  # Клиенты внешних систем и обёртки над ними
│   │   └── db/                  # Работа с БД
│   │       ├── pg/             # Инициализация pgx подключения
│   │       └── prettier/       # Форматирование SQL-запросов
│   │           └── query_prettier.go
│
│   ├── closer/                  # Утилита для graceful shutdown
│   ├── config/                  # Загрузка и работа с конфигурацией
│   ├── model/                   # Общие модели проекта (DTO/Entity)
│   ├── repository/             # Реализация доступа к данным (БД)
│   │   ├── intake/
│   │   ├── pickpoint/
│   │   └── user/
│   └── service/                # Бизнес-логика (usecases)
│       ├── intake/
│       ├── pickpoint/
│       └── user/
│
├── migrations/                 # SQL-скрипты для инициализации базы
│   └── init.sql
│
├── .env                        # Переменные окружения
├── .gitignore
├── Dockerfile                  # Docker-образ приложения
├── docker-compose.yml          # Docker Compose для поднятия зависимостей
├── go.mod
└── README.md                   # Документация проекта
```


# Инструкция по запуску сервиса
### но оно не работает, хахаха)

- Клонировать репозиторий в рабочую директорию:
```
https://github.com/hiphopzeliboba/pick-point.git
```
- Перейти в директорию go_service_api_tarantool:
```
cd pick-point
```
- Собрать и запустить сервис + инстанс тарантула Docker (должен быть установлен docker-compose):
```
docker-compose up --build -d
```
4. Сервис работает и доступен по адресу: `http://localhost:8080`
---