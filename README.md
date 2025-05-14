# Auth Service (Go + PostgreSQL + Docker)

Тестовое задание на позицию Junior Backend Developer.  
Реализован сервис авторизации с использованием JWT и refresh токенов.

---

## Функциональность

- Получение access и refresh токенов по GUID пользователя
- Обновление токенов с защитой от повторного использования
- Защищённый эндпоинт для получения информации о пользователе
- Деавторизация (logout)
- Валидация по IP и User-Agent
- Отправка webhook при подозрительной активности
- База данных — PostgreSQL
- Документация через Swagger

---

## Технологии

- Go (Golang)
- PostgreSQL
- JWT (SHA512)
- bcrypt
- Docker & Docker Compose
- Swagger / OpenAPI

---

