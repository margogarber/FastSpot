# MongoDB Schema для FastSpot

## Коллекции и их структура

### 1. users
Хранит информацию о пользователях системы (админы и гости).

```javascript
{
  _id: ObjectId,
  role: String,              // "admin" | "guest"
  name: String,
  email: String,             // уникальный
  phone: String,
  passwordHash: String,      // bcrypt hash
  createdAt: Date
}
```

**Индексы:**
- `email` (unique)
- `role`

---

### 2. categories
Категории меню (бургеры, напитки, десерты и т.д.).

```javascript
{
  _id: ObjectId,
  name: String,              // "Бургеры", "Напитки"
  slug: String,              // "burgers", "drinks"
  image: String,             // URL изображения
  isActive: Boolean          // доступна ли категория
}
```

**Индексы:**
- `slug` (unique)
- `isActive`

---

### 3. products
Товары меню с ингредиентами, опциями и тегами.

```javascript
{
  _id: ObjectId,
  categoryId: ObjectId,      // ссылка на categories._id
  name: String,              // "Биг Мак", "Кока-Кола"
  slug: String,              // "big-mac", "coca-cola"
  description: String,
  priceUSD: Number,          // цена в долларах (например, 5.99)
  image: String,             // URL изображения
  isActive: Boolean,
  ingredients: [             // настраиваемые ингредиенты
    {
      key: String,           // "pickles", "onions"
      label: String,         // "Маринованные огурцы"
      defaultIncluded: Boolean
    }
  ],
  options: [                 // дополнительные опции
    {
      key: String,           // "size", "sauce"
      label: String,         // "Размер", "Соус"
      type: String,          // "single" | "multiple"
      choices: [
        {
          value: String,     // "large", "medium"
          label: String,     // "Большой", "Средний"
          extraPriceUSD: Number // доплата в USD (0 если нет)
        }
      ]
    }
  ],
  tags: [String]             // ["spicy", "vegetarian", "comfort-food"]
}
```

**Индексы:**
- `categoryId`
- `slug` (unique)
- `isActive`
- `tags`

---

### 4. promotions
Акции и специальные предложения.

```javascript
{
  _id: ObjectId,
  title: String,             // "Летняя распродажа!"
  description: String,
  startsAt: Date,
  endsAt: Date,
  bannerImage: String,       // URL баннера
  isActive: Boolean,
  appliesTo: [ObjectId]      // массив productId, на которые действует акция
}
```

**Индексы:**
- `isActive`
- `startsAt`, `endsAt`

---

### 5. carts
Корзины пользователей (как авторизованных, так и гостей).

```javascript
{
  _id: ObjectId,
  userId: ObjectId | null,   // null для гостей
  sessionId: String,         // для гостей - ID сессии
  items: [
    {
      productId: ObjectId,
      qty: Number,
      chosenIngredients: [String],  // ["pickles", "onions"]
      chosenOptions: {              // {"size": "large", "sauce": "bbq"}
        [key: String]: String | [String]
      },
      unitPriceUSD: Number,   // цена за единицу с учетом опций
      totalUSD: Number        // unitPriceUSD * qty
    }
  ],
  totalUSD: Number,           // сумма всех items[].totalUSD
  currency: String,           // "USD"
  updatedAt: Date
}
```

**Индексы:**
- `userId`
- `sessionId`

---

### 6. orders
Завершенные заказы с полной информацией о доставке и оплате.

```javascript
{
  _id: ObjectId,
  userId: ObjectId | null,   // null для гостей
  cartSnapshot: {            // копия корзины на момент заказа
    items: [/* структура как в carts.items */],
    totalUSD: Number
  },
  status: String,            // "pending" | "preparing" | "ready" | "delivering" | "completed" | "cancelled"
  payment: {
    method: String,          // "card" | "applepay" | "googlepay" | "cash"
    status: String           // "pending" | "paid" | "failed"
  },
  delivery: {
    type: String,            // "pickup" | "courier"
    address: String,         // адрес (если courier)
    eta: Date,               // ожидаемое время
    tracking: [              // история статусов
      {
        ts: Date,
        status: String,
        note: String
      }
    ]
  },
  totalUSD: Number,
  createdAt: Date
}
```

**Индексы:**
- `userId`
- `status`
- `createdAt`

---

### 7. mood_questions
Вопросы для "mood quiz" (тест настроения для AI-рекомендаций).

```javascript
{
  _id: ObjectId,
  order: Number,             // порядок вопроса
  question: String,          // "Как вы себя чувствуете сегодня?"
  type: String,              // "single" | "multiple" | "scale"
  options: [
    {
      value: String,         // "happy", "tired", "hungry"
      label: String,         // "Счастлив 😊"
      weight: Object         // {"energy": 1, "comfort": 0.5}
    }
  ],
  isActive: Boolean
}
```

**Индексы:**
- `order`
- `isActive`

---

### 8. mood_rules
Правила для fallback рекомендаций (если AI недоступен).

```javascript
{
  _id: ObjectId,
  name: String,              // "Comfort Food Lover"
  conditions: {              // условия для срабатывания правила
    tags: [String],          // требуемые теги в ответах
    minScore: Object         // {"comfort": 3, "energy": 1}
  },
  recommendedProducts: [ObjectId], // массив productId
  priority: Number,          // приоритет правила (больше = выше)
  isActive: Boolean
}
```

**Индексы:**
- `priority`
- `isActive`

---

### 9. ai_sessions
История сессий AI-рекомендаций для аналитики.

```javascript
{
  _id: ObjectId,
  userId: ObjectId | null,
  sessionId: String,
  answers: [                 // ответы пользователя на quiz
    {
      questionId: ObjectId,
      selectedOptions: [String]
    }
  ],
  aiResponse: {              // ответ от Gemini API
    raw: String,             // сырой текст ответа
    recommendedProducts: [ObjectId],
    reasoning: String        // объяснение рекомендации
  },
  fallbackUsed: Boolean,     // использовались ли правила вместо AI
  createdAt: Date
}
```

**Индексы:**
- `userId`
- `sessionId`
- `createdAt`

---

## Связи между коллекциями

```
users (1) ──────────> (n) carts
users (1) ──────────> (n) orders
users (1) ──────────> (n) ai_sessions

categories (1) ─────> (n) products

products (n) ──────> (n) promotions.appliesTo
products (n) ──────> (n) carts.items[].productId
products (n) ──────> (n) orders.cartSnapshot.items[].productId
products (n) ──────> (n) mood_rules.recommendedProducts
products (n) ──────> (n) ai_sessions.aiResponse.recommendedProducts

mood_questions (n) ─> (n) ai_sessions.answers[].questionId
```

---

## Примечания

1. **Валюта**: Вся система работает в USD (доллары США)
2. **Гости**: Для гостевых заказов используется `sessionId` вместо `userId`
3. **Корзина**: Сохраняется в `cartSnapshot` при создании заказа для истории
4. **AI Fallback**: Если Gemini API недоступен, используются `mood_rules`
5. **Индексы**: Обязательно создать для производительности

