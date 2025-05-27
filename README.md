
# Котошоп — крутецкий интернет-магазин товаров, а это его сервер



### Backend

- **Go** (язык программирования)
- **Gin** (веб-фреймворк)
- **GORM** (ORM для работы с БД)
- **JWT** (аутентификация)

### База данных

- **PostgreSQL** (основное хранилище)

## Установка

### Требования

- Go 1.20+
- PostgreSQL 14+

### Шаги

1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/shagalinD/kotoshop-server
   cd kotoshop-server
   ```


2. Настроить БД:

   ```bash
   createdb kotoshop
   ```
   
## Конфигурация

Создайте файл `.env` в папке `server` по примеру в .env.example


3. Запустить сервер (из папки `backend`):

   ```bash
   cd server
   go mod tidy
   go run main.go
   ```



## Структура проекта

```
kotoshop-server/
├── server/               # Backend
│   ├── models/           # Модели БД
│   ├── handlers/      # Обработчики запросов
│   └── postgres/           # Подключение к БД
```

## Примеры API-запросов

### Получить товары

```http
GET /api/products/get_all
```

### Добавить в корзину

```http
POST /api/cart/add_product
{
  "product_id": 5,
  "quantity": 2
}
```

### Создать заказ

```http
POST /api/orders/create
{
  "address": "Россия, Мрляндия, Котовый переулок, 15"
}
```

### Скрипт Postman 

Для того, чтобы проверить все запросы, склонируйте и запустите Fork для скрипта в Postman по данной ссылке: https://www.postman.com/dmitriy-8345409/workspace/studing/collection/45000778-cc0bb7d1-8652-4ecb-a840-8e2cd81c281e?action=share&creator=45000778 
