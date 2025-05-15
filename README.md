
# Котошоп — Интернет-магазин товаров для котиков

Невероятно модный онлайн магазин с невероятным функционалом. Ваш ждут:

- 📦 Просмотр каталога товаров с фильтрацией и поиском
- 🛒 Добавление товаров в корзину
- 💳 Оформление заказа с выбором адреса доставки
- 📝 Просмотр истории заказов
- 🔐 Авторизация и регистрация пользователей
- ✏️ Редактирование профиля (имя, телефон)

## Технологии

### Frontend

- **React** + **React Router**
- **Redux Toolkit** (управление состоянием)
- **Material-UI** (стилизация)
- **Axios** (HTTP-запросы)

### Backend

- **Go** (язык программирования)
- **Gin** (веб-фреймворк)
- **GORM** (ORM для работы с БД)
- **JWT** (аутентификация)

### База данных

- **PostgreSQL** (основное хранилище)

## Установка

### Требования

- Node.js v16+
- Go 1.20+
- PostgreSQL 14+

### Шаги

1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/shagalinD/MyCompletedShop
   cd MyCompletedShop
   ```


2. Настроить БД:

   ```bash
   createdb kotoshop
   ```

3. Запустить сервер (из папки `backend`):

   ```bash
   cd server
   go mod tidy
   go run main.go
   ```

4. Запустить клиент (из корневой папки):
   ```bash
   cd client
   npm install
   npm start
   ```

## Конфигурация

Создайте файл `.env` в папке `backend`:

```ini
POSTGRES_STRING=user=user password=12345678 dbname=kotoshop port=5432
SECRET_KEY=h3co2iy523y4c1adf34c24rc23c234c234c234c249uyc103uc193yc19
```

## Структура проекта

```
MyCompletedShop/
│   # Frontend
│   src/
│   ├── components/   # React-компоненты
│   ├── features/     # Redux slices
│   └── pages/        # Страницы приложения
│
├── backend/               # Backend
│   ├── models/           # Модели БД
│   ├── handlers/      # Обработчики запросов
│   └── routes/           # Маршруты API
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

🐾 **Приятных покупок вашим котикам!** 🐾
