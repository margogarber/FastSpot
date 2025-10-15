# FastSpot Database# FastSpot Database Setup



Setup and usage guide for FastSpot's MongoDB database.Инструкции по настройке и работе с базой данных MongoDB для приложения FastSpot.



## Requirements## 📋 Содержание



- Node.js >= 16- [Требования](#требования)

- MongoDB >= 5.0- [Установка](#установка)

- npm- [Конфигурация](#конфигурация)

- [Заполнение данными](#заполнение-данными)

## Installation- [Структура данных](#структура-данных)

- [Полезные команды](#полезные-команды)

### 1. Install Dependencies

---

```bash

npm install## 🔧 Требования

```

- **Node.js** >= 16.x

### 2. Start MongoDB- **MongoDB** >= 5.0

- **npm** или **yarn**

**Local:**

```bash---

# macOS

brew services start mongodb-community## 📦 Установка



# Linux### 1. Установите зависимости

sudo systemctl start mongod

``````bash

npm install mongodb bcryptjs

**Docker:**```

```bash

docker-compose up -d mongodbИли если используете yarn:

```

```bash

## Configurationyarn add mongodb bcryptjs

```

Create `.env` file:

### 2. Убедитесь, что MongoDB запущен

```env

MONGODB_URI=mongodb://admin:Admin123!@localhost:27017/fastspot?authSource=admin#### Локальная установка:

``````bash

# macOS (через Homebrew)

For MongoDB Atlas (cloud):brew services start mongodb-community

```env

MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/fastspot# Linux (systemd)

```sudo systemctl start mongod



## Seeding Data# Проверка статуса

mongosh --eval "db.runCommand({ ping: 1 })"

Run the seed script to populate the database:```



```bash#### Docker:

node seed.js```bash

```docker run -d \

  --name fastspot-mongo \

This creates:  -p 27017:27017 \

- **1 admin user**  -e MONGO_INITDB_ROOT_USERNAME=admin \

  - Email: `admin@local`  -e MONGO_INITDB_ROOT_PASSWORD=password \

  - Password: `Admin123!`  mongo:7

- **4 categories** (Burgers, Drinks, Desserts, Snacks)```

- **10 products** with customization options

- **3 promotions**---

- **8 mood quiz questions**

- **Indexes** for better performance## ⚙️ Конфигурация



## Database Structure### Переменные окружения



See [schema.md](./schema.md) for detailed collection schemas.Создайте файл `.env` в корне проекта:



### Main Collections```env

# MongoDB Connection

- **users** - Admin and guest accountsMONGODB_URI=mongodb://localhost:27017/fastspot

- **categories** - Menu categories

- **products** - Menu items with options# Для MongoDB Atlas (облачная версия)

- **promotions** - Special offers# MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/fastspot?retryWrites=true&w=majority

- **carts** - Shopping carts

- **orders** - Completed orders# Для Docker с авторизацией

- **mood_questions** - AI quiz questions# MONGODB_URI=mongodb://admin:password@localhost:27017/fastspot?authSource=admin

- **mood_rules** - AI recommendation rules```

- **ai_sessions** - AI recommendation history

### Настройка MongoDB Atlas (опционально)

## Useful Commands

Если вы хотите использовать облачную MongoDB:

### Connect to MongoDB

1. Зарегистрируйтесь на [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)

```bash2. Создайте бесплатный кластер

# Local3. Настройте Network Access (добавьте свой IP или `0.0.0.0/0` для разработки)

mongosh fastspot4. Создайте пользователя базы данных

5. Получите строку подключения (Connection String)

# With auth6. Вставьте её в `MONGODB_URI` в файле `.env`

mongosh -u admin -p Admin123! --authenticationDatabase admin fastspot

```---



### View Data## 🌱 Заполнение данными



```javascript### Запуск скрипта seed

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

