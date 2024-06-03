# Description: Скрипт для тестирования API с помощью curl
# Author: Alexey Pavlinov
#
# Примеры использования:
# Регистрация
curl -X POST http://localhost:8084/auth/register -H "Content-Type: application/json" -d '{"username": "user1", "password": "password"}'

# Авторизация
curl -X POST http://localhost:8084/auth/login -H "Content-Type: application/json" -d '{"username": "user1", "password": "password"}'

# Создание статьи
curl -X POST http://localhost:8084/articles/ -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title": "Article 1", "content": "Content of Article 1"}'

# Установка предпочтений
curl -X POST http://localhost:8084/preferences -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"article_id": 1, "liked": true}'

# Получение предпочтений
curl -X GET http://localhost:8084/preferences -H "Authorization: Bearer <token>"


