# 🚀 Быстрый старт FastSpot Database

## За 3 шага к готовой базе данных!

### Шаг 1️⃣: Установка зависимостей

```bash
npm install
```

### Шаг 2️⃣: Настройка подключения

Создайте файл `.env` в корне проекта:

```bash
echo "MONGODB_URI=mongodb://localhost:27017/fastspot" > .env
```

Или для MongoDB Atlas:
```bash
echo "MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/fastspot" > .env
```

### Шаг 3️⃣: Запуск seed скрипта

```bash
npm run seed
```

---

## ✅ Проверка

Убедитесь, что всё работает:

```bash
npm run test:db
```

Должен вывести:
```
✅ Подключение успешно!
📊 База данных: fastspot
📦 Найдено коллекций: 9
```

---

## 🎯 Что дальше?

### Вход в систему

Используйте учетные данные администратора:
- **Email**: `admin@local`
- **Password**: `Admin123!`

### Просмотр данных через MongoDB Compass

1. Скачайте [MongoDB Compass](https://www.mongodb.com/products/compass)
2. Подключитесь используя ваш `MONGODB_URI`
3. Откройте базу `fastspot`
4. Исследуйте коллекции!

### Структура данных

- **11 продуктов** в 4 категориях
- **3 активных акции**
- **8 вопросов** для mood quiz
- **4 правила** AI-рекомендаций

---

## 🐛 Проблемы?

### MongoDB не запущен?

**macOS:**
```bash
brew services start mongodb-community
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

