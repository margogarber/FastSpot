# FastSpot Database Setup

Инструкции по настройке и работе с базой данных MongoDB для приложения FastSpot.

## 📋 Содержание

- [Требования](#требования)
- [Установка](#установка)
- [Конфигурация](#конфигурация)
- [Заполнение данными](#заполнение-данными)
- [Структура данных](#структура-данных)
- [Полезные команды](#полезные-команды)

---

## 🔧 Требования

- **Node.js** >= 16.x
- **MongoDB** >= 5.0
- **npm** или **yarn**

---

## 📦 Установка

### 1. Установите зависимости

```bash
npm install mongodb bcryptjs
```

Или если используете yarn:

```bash
yarn add mongodb bcryptjs
```

### 2. Убедитесь, что MongoDB запущен

#### Локальная установка:
```bash
# macOS (через Homebrew)
brew services start mongodb-community

# Linux (systemd)
sudo systemctl start mongod

# Проверка статуса
mongosh --eval "db.runCommand({ ping: 1 })"
```

#### Docker:
```bash
docker run -d \
  --name fastspot-mongo \
  -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=admin \
  -e MONGO_INITDB_ROOT_PASSWORD=password \
  mongo:7
```

---

## ⚙️ Конфигурация

### Переменные окружения

Создайте файл `.env` в корне проекта:

```env
# MongoDB Connection
MONGODB_URI=mongodb://localhost:27017/fastspot

# Для MongoDB Atlas (облачная версия)
# MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/fastspot?retryWrites=true&w=majority

# Для Docker с авторизацией
# MONGODB_URI=mongodb://admin:password@localhost:27017/fastspot?authSource=admin
```

### Настройка MongoDB Atlas (опционально)

Если вы хотите использовать облачную MongoDB:

1. Зарегистрируйтесь на [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
2. Создайте бесплатный кластер
3. Настройте Network Access (добавьте свой IP или `0.0.0.0/0` для разработки)
4. Создайте пользователя базы данных
5. Получите строку подключения (Connection String)
6. Вставьте её в `MONGODB_URI` в файле `.env`

---

## 🌱 Заполнение данными

### Запуск скрипта seed

```bash
# Из корня проекта
cd db
node seed.js
```

Или с указанием MongoDB URI напрямую:

```bash
MONGODB_URI=mongodb://localhost:27017/fastspot node seed.js
```

### Что создает скрипт?

Скрипт `seed.js` создает:

- ✅ **1 пользователь-администратор**
  - Email: `admin@local`
  - Password: `Admin123!`
  
- ✅ **4 категории**: Бургеры, Напитки, Десерты, Снэки

- ✅ **11 продуктов** с:
  - Настраиваемыми ингредиентами
  - Дополнительными опциями (размер, соус и т.д.)
  - Тегами для AI-рекомендаций
  - Ценами в USD

- ✅ **3 акции** (промо)

- ✅ **8 вопросов** для mood quiz

- ✅ **4 правила** для fallback рекомендаций

- ✅ **Индексы** для оптимизации запросов

---

## 📚 Структура данных

Подробное описание всех коллекций и их связей смотрите в [schema.md](./schema.md).

### Основные коллекции:

1. **users** - Пользователи (админы и гости)
2. **categories** - Категории меню
3. **products** - Продукты с ингредиентами и опциями
4. **promotions** - Акции и спецпредложения
5. **carts** - Корзины пользователей
6. **orders** - Заказы с доставкой и оплатой
7. **mood_questions** - Вопросы для mood quiz
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

