# 🚀 Quick Start# 🚀 Быстрый старт FastSpot Database



Get your database up and running in 3 steps!## За 3 шага к готовой базе данных!



## Step 1: Install Dependencies### Шаг 1️⃣: Установка зависимостей



```bash```bash

npm installnpm install

``````



## Step 2: Configure Connection### Шаг 2️⃣: Настройка подключения



Create `.env` file:Создайте файл `.env` в корне проекта:



```bash```bash

echo "MONGODB_URI=mongodb://admin:Admin123!@localhost:27017/fastspot?authSource=admin" > .envecho "MONGODB_URI=mongodb://localhost:27017/fastspot" > .env

``````



Or for MongoDB Atlas (cloud):Или для MongoDB Atlas:

```bash```bash

echo "MONGODB_URI=mongodb+srv://user:pass@cluster.mongodb.net/fastspot" > .envecho "MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/fastspot" > .env

``````



## Step 3: Seed the Database### Шаг 3️⃣: Запуск seed скрипта



```bash```bash

node seed.jsnpm run seed

``````



✨ Done! Your database now has:---

- 1 admin user

- 4 categories## ✅ Проверка

- 10 products

- 3 promotionsУбедитесь, что всё работает:

- 8 quiz questions

```bash

## 🔑 Admin Loginnpm run test:db

```

```

Email: admin@localДолжен вывести:

Password: Admin123!```

```✅ Подключение успешно!

📊 База данных: fastspot

## ✅ Test Connection📦 Найдено коллекций: 9

```

```bash

node test-connection.js---

```

## 🎯 Что дальше?

Should show:

```### Вход в систему

✅ Connected successfully!

📊 Database: fastspotИспользуйте учетные данные администратора:

📦 Collections: 9- **Email**: `admin@local`

```- **Password**: `Admin123!`



## 🐛 Troubleshooting### Просмотр данных через MongoDB Compass



### MongoDB not running?1. Скачайте [MongoDB Compass](https://www.mongodb.com/products/compass)

2. Подключитесь используя ваш `MONGODB_URI`

```bash3. Откройте базу `fastspot`

# macOS4. Исследуйте коллекции!

brew services start mongodb-community

### Структура данных

# Docker

docker-compose up -d mongodb- **11 продуктов** в 4 категориях

```- **3 активных акции**

- **8 вопросов** для mood quiz

### Connection error?- **4 правила** AI-рекомендаций



1. Check MongoDB is running---

2. Verify `.env` file exists

3. Check port 27017 is available## 🐛 Проблемы?



## 📚 Next Steps### MongoDB не запущен?



- [README.md](./README.md) - Full documentation**macOS:**

- [schema.md](./schema.md) - Database structure```bash

- [QUERIES.md](./QUERIES.md) - Useful queriesbrew services start mongodb-community

```

**Linux:**
```bash
sudo systemctl start mongod
```

**Docker:**
```bash
docker run -d --name fastspot-mongo -p 27017:27017 mongo:7
```

### Ошибка подключения?

Проверьте:
1. MongoDB запущен: `mongosh --eval "db.runCommand({ ping: 1 })"`
2. `.env` файл создан и содержит правильный `MONGODB_URI`
3. Порт 27017 доступен: `lsof -i :27017` (macOS/Linux)

---

## 📚 Дополнительная информация

- [README.md](./README.md) - Подробная документация
- [schema.md](./schema.md) - Структура базы данных
- [seed.js](./seed.js) - Скрипт для заполнения данных

---

**Готово! Теперь можно начинать разработку! 💪**

