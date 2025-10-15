# 🗄️ FastSpot Database# FastSpot Database# FastSpot Database Setup



Simple guide to understanding and setting up the FastSpot database.



---Setup and usage guide for FastSpot's MongoDB database.Инструкции по настройке и работе с базой данных MongoDB для приложения FastSpot.



## What's This?



FastSpot uses **MongoDB** to store everything:## Requirements## 📋 Содержание

- Menu items (burgers, drinks, etc.)

- User orders

- Shopping carts

- Admin accounts- Node.js >= 16- [Требования](#требования)

- Promotions

- MongoDB >= 5.0- [Установка](#установка)

---

- npm- [Конфигурация](#конфигурация)

## Quick Setup (3 Steps!)

- [Заполнение данными](#заполнение-данными)

### 1️⃣ Install Packages

```bash## Installation- [Структура данных](#структура-данных)

npm install

```- [Полезные команды](#полезные-команды)



### 2️⃣ Create `.env` File### 1. Install Dependencies

```bash

MONGODB_URI=mongodb://admin:Admin123!@localhost:27017/fastspot?authSource=admin---

```

```bash

### 3️⃣ Fill Database with Sample Data

```bashnpm install## 🔧 Требования

node seed.js

``````



✅ **Done!** You now have:- **Node.js** >= 16.x

- 1 admin account

- 4 categories (Burgers, Drinks, Desserts, Snacks)### 2. Start MongoDB- **MongoDB** >= 5.0

- 10 menu items

- 3 special promotions- **npm** или **yarn**



---**Local:**



## 🔑 Login Info```bash---



After running the setup, you can log in as admin:# macOS



```brew services start mongodb-community## 📦 Установка

Email: admin@local

Password: Admin123!

```

# Linux### 1. Установите зависимости

---

sudo systemctl start mongod

## What's in the Database?

``````bash

### 📦 Main Collections (Tables)

npm install mongodb bcryptjs

| Collection | What It Stores |

|------------|----------------|**Docker:**```

| **users** | Admin and customer accounts |

| **categories** | Menu sections (Burgers, Drinks, etc.) |```bash

| **products** | Menu items with prices & options |

| **carts** | Active shopping carts |docker-compose up -d mongodbИли если используете yarn:

| **orders** | Completed purchases |

| **promotions** | Special deals & discounts |```

| **mood_questions** | Quiz questions for recommendations |

| **ai_sessions** | AI recommendation history |```bash



---## Configurationyarn add mongodb bcryptjs



## How Things Connect```



```Create `.env` file:

Categories → Products → Carts → Orders

                ↓### 2. Убедитесь, что MongoDB запущен

            Promotions

``````env



**Example:**MONGODB_URI=mongodb://admin:Admin123!@localhost:27017/fastspot?authSource=admin#### Локальная установка:

1. "Burgers" is a **category**

2. "Big Mac" is a **product** in that category``````bash

3. Customer adds Big Mac to their **cart**

4. Customer checks out → Creates an **order**# macOS (через Homebrew)

5. **Promotions** can apply discounts to products

For MongoDB Atlas (cloud):brew services start mongodb-community

---

```env

## Common Tasks

MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/fastspot# Linux (systemd)

### View All Menu Items

```bash```sudo systemctl start mongod

mongosh fastspot

db.products.find().pretty()

```

## Seeding Data# Проверка статуса

### Check Orders

```bashmongosh --eval "db.runCommand({ ping: 1 })"

db.orders.find().pretty()

```Run the seed script to populate the database:```



### See Active Promotions

```bash

db.promotions.find({ isActive: true })```bash#### Docker:

```

node seed.js```bash

### Reset Database

```bash```docker run -d \

node seed.js

```  --name fastspot-mongo \

This will clear everything and start fresh with sample data.

This creates:  -p 27017:27017 \

---

- **1 admin user**  -e MONGO_INITDB_ROOT_USERNAME=admin \

## Troubleshooting

  - Email: `admin@local`  -e MONGO_INITDB_ROOT_PASSWORD=password \

### ❌ "Can't connect to MongoDB"

  - Password: `Admin123!`  mongo:7

**Make sure MongoDB is running:**

- **4 categories** (Burgers, Drinks, Desserts, Snacks)```

```bash

# macOS- **10 products** with customization options

brew services start mongodb-community

- **3 promotions**---

# Or use Docker

docker-compose up -d mongodb- **8 mood quiz questions**

```

- **Indexes** for better performance## ⚙️ Конфигурация

### ❌ "Authentication failed"



Check your `.env` file has the correct connection string:

```## Database Structure### Переменные окружения

MONGODB_URI=mongodb://admin:Admin123!@localhost:27017/fastspot?authSource=admin

```



### ❌ "Port 27017 in use"See [schema.md](./schema.md) for detailed collection schemas.Создайте файл `.env` в корне проекта:



Something else is using MongoDB's port:

```bash

lsof -i :27017### Main Collections```env

```

# MongoDB Connection

---

- **users** - Admin and guest accountsMONGODB_URI=mongodb://localhost:27017/fastspot

## Useful Tools

- **categories** - Menu categories

**MongoDB Compass** - Visual database browser

- Download: https://www.mongodb.com/products/compass- **products** - Menu items with options# Для MongoDB Atlas (облачная версия)

- Connect with your `MONGODB_URI`

- Browse and edit data visually- **promotions** - Special offers# MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/fastspot?retryWrites=true&w=majority



**mongosh** - Command-line tool- **carts** - Shopping carts

```bash

mongosh fastspot- **orders** - Completed orders# Для Docker с авторизацией

show collections

db.products.find()- **mood_questions** - AI quiz questions# MONGODB_URI=mongodb://admin:password@localhost:27017/fastspot?authSource=admin

```

- **mood_rules** - AI recommendation rules```

---

- **ai_sessions** - AI recommendation history

## Need More Info?

### Настройка MongoDB Atlas (опционально)

- **schema.md** - Detailed structure of each collection

- **QUERIES.md** - Useful database queries## Useful Commands

- **QUICKSTART.md** - Fast setup guide

- **seed.js** - See what sample data looks likeЕсли вы хотите использовать облачную MongoDB:



---### Connect to MongoDB



## Sample Data Included1. Зарегистрируйтесь на [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)



After running `seed.js`, you'll have:```bash2. Создайте бесплатный кластер



**Products:**# Local3. Настройте Network Access (добавьте свой IP или `0.0.0.0/0` для разработки)

- 🍔 Big Mac, Cheeseburger, Veggie Burger

- 🥤 Coca-Cola, Sprite, Orange Juicemongosh fastspot4. Создайте пользователя базы данных

- 🍰 Chocolate Cake, Apple Pie

- 🍟 French Fries, Onion Rings5. Получите строку подключения (Connection String)



**Categories:**# With auth6. Вставьте её в `MONGODB_URI` в файле `.env`

- Burgers

- Drinks  mongosh -u admin -p Admin123! --authenticationDatabase admin fastspot

- Desserts

- Snacks```---



**Promotions:**

- Buy 1 Get 1 Free

- 20% Off Drinks### View Data## 🌱 Заполнение данными

- Combo Deals



---

```javascript### Запуск скрипта seed

**That's it! You're ready to start developing! 🚀**

// Show collections

show collections```bash

# Из корня проекта

// View productscd db

db.products.find().pretty()node seed.js

```

// Active categories

db.categories.find({ isActive: true })Или с указанием MongoDB URI напрямую:



// Find by tags```bash

db.products.find({ tags: "vegetarian" })MONGODB_URI=mongodb://localhost:27017/fastspot node seed.js

``````



### Clear Database### Что создает скрипт?



```javascriptСкрипт `seed.js` создает:

// Clear all data

db.users.deleteMany({})- ✅ **1 пользователь-администратор**

db.categories.deleteMany({})  - Email: `admin@local`

db.products.deleteMany({})  - Password: `Admin123!`

db.orders.deleteMany({})  

db.carts.deleteMany({})- ✅ **4 категории**: Бургеры, Напитки, Десерты, Снэки

```

- ✅ **11 продуктов** с:

## MongoDB Compass  - Настраиваемыми ингредиентами

  - Дополнительными опциями (размер, соус и т.д.)

For visual database management:  - Тегами для AI-рекомендаций

  - Ценами в USD

1. Download [MongoDB Compass](https://www.mongodb.com/products/compass)

2. Connect using your `MONGODB_URI`- ✅ **3 акции** (промо)

3. Browse and edit data visually

- ✅ **8 вопросов** для mood quiz

## Troubleshooting

- ✅ **4 правила** для fallback рекомендаций

### Can't connect?

- ✅ **Индексы** для оптимизации запросов

- Check MongoDB is running: `mongosh --eval "db.runCommand({ ping: 1 })"`

- Verify `.env` file exists---

- Check port 27017: `lsof -i :27017`

## 📚 Структура данных

### Authentication errors?

Подробное описание всех коллекций и их связей смотрите в [schema.md](./schema.md).

- Ensure MongoDB auth is set up correctly

- Check username/password in `MONGODB_URI`### Основные коллекции:

- Verify `authSource=admin` is included

1. **users** - Пользователи (админы и гости)

## More Info2. **categories** - Категории меню

3. **products** - Продукты с ингредиентами и опциями

- [QUICKSTART.md](./QUICKSTART.md) - Quick setup guide4. **promotions** - Акции и спецпредложения

- [schema.md](./schema.md) - Database schema details5. **carts** - Корзины пользователей

- [QUERIES.md](./QUERIES.md) - Useful queries6. **orders** - Заказы с доставкой и оплатой

- [DIAGRAM.md](./DIAGRAM.md) - Visual overview7. **mood_questions** - Вопросы для mood quiz

8. **mood_rules** - Правила для AI-рекомендаций
9. **ai_sessions** - История AI-рекомендаций

---

## 🛠️ Полезные команды

### Подключение к MongoDB через CLI

```bash
# Локальная MongoDB
mongosh fastspot

# С авторизацией
mongosh -u admin -p password --authenticationDatabase admin fastspot

# MongoDB Atlas
mongosh "mongodb+srv://cluster.mongodb.net/fastspot" --username your-username
```

### Просмотр данных

```javascript
// Показать все коллекции
show collections

// Посмотреть продукты
db.products.find().pretty()

// Посмотреть активные категории
db.categories.find({ isActive: true })

// Найти продукты по тегу
db.products.find({ tags: "comfort-food" })

// Посмотреть все акции
db.promotions.find()
```

### Очистка базы данных

```javascript
// Удалить все данные из всех коллекций
db.users.deleteMany({})
db.categories.deleteMany({})
db.products.deleteMany({})
db.promotions.deleteMany({})
db.carts.deleteMany({})
db.orders.deleteMany({})
db.mood_questions.deleteMany({})
db.mood_rules.deleteMany({})
db.ai_sessions.deleteMany({})

// Или полностью удалить базу данных
use fastspot
db.dropDatabase()
```

### Обновление данных

После изменения `seed.js` просто запустите его снова:

```bash
node seed.js
```

Скрипт автоматически:
1. Удалит все существующие данные
2. Создаст новые данные
3. Пересоздаст индексы

---

## 🔐 Безопасность

### Для продакшена:

1. **Используйте сильные пароли**
   ```javascript
   // Измените пароль администратора в seed.js
   const adminPassword = await bcrypt.hash('YourStrongPassword123!@#', 10);
   ```

2. **Настройте авторизацию MongoDB**
   ```bash
   # Создайте пользователя с ограниченными правами
   mongosh
   use admin
   db.createUser({
     user: "fastspot_app",
     pwd: "secure_password",
     roles: [{ role: "readWrite", db: "fastspot" }]
   })
   ```

3. **Используйте переменные окружения**
   - Никогда не коммитьте `.env` файл в Git
   - Добавьте `.env` в `.gitignore`

4. **Для MongoDB Atlas**
   - Ограничьте IP-адреса в Network Access
   - Используйте сильные пароли
   - Включите шифрование в transit и at rest

---

## 🧪 Тестирование

### Проверка подключения

```bash
# Создайте тестовый скрипт test-connection.js
cat << 'EOF' > test-connection.js
const { MongoClient } = require('mongodb');

async function test() {
  const uri = process.env.MONGODB_URI || 'mongodb://localhost:27017/fastspot';
  const client = new MongoClient(uri);
  
  try {
    await client.connect();
    console.log('✅ Подключение успешно!');
    
    const db = client.db();
    const collections = await db.listCollections().toArray();
    console.log(`📊 Найдено коллекций: ${collections.length}`);
    collections.forEach(c => console.log(`   - ${c.name}`));
    
  } catch (error) {
    console.error('❌ Ошибка подключения:', error.message);
  } finally {
    await client.close();
  }
}

test();
EOF

# Запустите тест
node test-connection.js
```

---

## 📞 Поддержка

Если возникли проблемы:

1. Убедитесь, что MongoDB запущен
2. Проверьте `MONGODB_URI` в `.env`
3. Проверьте версию Node.js: `node --version`
4. Проверьте логи MongoDB

### Частые ошибки

**Ошибка: "MongoServerError: Authentication failed"**
- Проверьте логин/пароль в `MONGODB_URI`
- Убедитесь, что пользователь создан в MongoDB

**Ошибка: "MongoNetworkError: connect ECONNREFUSED"**
- MongoDB не запущен
- Неверный хост/порт в `MONGODB_URI`

**Ошибка: "Cannot find module 'mongodb'"**
- Установите зависимости: `npm install`

---

## 📄 Лицензия

Этот проект создан для учебных целей.

---

**Готово! 🚀** 

Теперь у вас есть полностью настроенная база данных для FastSpot с демо-данными!

