# 📊 Диаграмма структуры базы данных FastSpot

## Общая схема связей

```mermaid
erDiagram
    users ||--o{ carts : "has"
    users ||--o{ orders : "places"
    users ||--o{ ai_sessions : "creates"
    
    categories ||--o{ products : "contains"
    
    products }o--o{ promotions : "appliesTo"
    products }o--o{ carts : "items"
    products }o--o{ orders : "cartSnapshot"
    products }o--o{ mood_rules : "recommendedProducts"
    products }o--o{ ai_sessions : "recommends"
    
    mood_questions }o--o{ ai_sessions : "answers"
    
    users {
        ObjectId _id PK
        string role
        string name
        string email UK
        string phone
        string passwordHash
        date createdAt
    }
    
    categories {
        ObjectId _id PK
        string name
        string slug UK
        string image
        boolean isActive
    }
    
    products {
        ObjectId _id PK
        ObjectId categoryId FK
        string name
        string slug UK
        string description
        number priceUSD
        string image
        boolean isActive
        array ingredients
        array options
        array tags
    }
    
    promotions {
        ObjectId _id PK
        string title
        string description
        date startsAt
        date endsAt
        string bannerImage
        boolean isActive
        array appliesTo
    }
    
    carts {
        ObjectId _id PK
        ObjectId userId FK
        string sessionId
        array items
        number totalUSD
        string currency
        date updatedAt
    }
    
    orders {
        ObjectId _id PK
        ObjectId userId FK
        object cartSnapshot
        string status
        object payment
        object delivery
        number totalUSD
        date createdAt
    }
    
    mood_questions {
        ObjectId _id PK
        number order
        string question
        string type
        array options
        boolean isActive
    }
    
    mood_rules {
        ObjectId _id PK
        string name
        object conditions
        array recommendedProducts
        number priority
        boolean isActive
    }
    
    ai_sessions {
        ObjectId _id PK
        ObjectId userId FK
        string sessionId
        array answers
        object aiResponse
        boolean fallbackUsed
        date createdAt
    }
```

---

## Основные потоки данных

### 1. 🛒 Процесс заказа

```mermaid
flowchart LR
    A[User] -->|browse| B[Products]
    B -->|add to| C[Cart]
    C -->|checkout| D[Order]
    D -->|snapshot| C
    
    E[Categories] -->|organize| B
    F[Promotions] -.->|discount| B
```

### 2. 🤖 AI-рекомендации

```mermaid
flowchart TD
    A[User начинает Quiz] --> B[Mood Questions]
    B --> C[User отвечает]
    C --> D[AI Session создан]
    D --> E{Gemini API доступен?}
    E -->|Да| F[Gemini анализирует]
    E -->|Нет| G[Mood Rules применяются]
    F --> H[Рекомендации]
    G --> H
    H --> I[Products добавлены в Cart]
```

### 3. 👤 Типы пользователей

```mermaid
flowchart LR
    A[User] --> B{Авторизован?}
    B -->|Да| C[userId]
    B -->|Нет| D[sessionId]
    C --> E[Cart/Orders привязаны к userId]
    D --> F[Cart привязана к sessionId]
    
    style C fill:#90EE90
    style D fill:#FFB6C1
```

---

## Детальная структура коллекций

### Products - самая сложная структура

```mermaid
graph TB
    A[Product] --> B[Basic Info]
    A --> C[Pricing]
    A --> D[Customization]
    A --> E[Relations]
    
    B --> B1[name, slug, description]
    B --> B2[image, isActive]
    
    C --> C1[priceUSD]
    C --> C2[currency: USD]
    
    D --> D1[Ingredients]
    D --> D2[Options]
    D --> D3[Tags]
    
    D1 --> D1A[pickles, onions, lettuce]
    D1 --> D1B[defaultIncluded: true/false]
    
    D2 --> D2A[Size: small/medium/large]
    D2 --> D2B[Sauce: bbq/spicy/classic]
    D2 --> D2C[extraPriceUSD per choice]
    
    D3 --> D3A[comfort-food, healthy]
    D3 --> D3B[vegetarian, spicy]
    D3 --> D3C[для AI matching]
    
    E --> E1[categoryId → Categories]
    E --> E2[используется в Carts]
    E --> E3[используется в Orders]
    E --> E4[используется в Promotions]
```

### Cart Items - структура элемента корзины

```json
{
  "items": [
    {
      "productId": "ObjectId",
      "qty": 2,
      "chosenIngredients": ["pickles", "onions"],
      "chosenOptions": {
        "size": "large",
        "sauce": "bbq"
      },
      "unitPriceUSD": 7.99,
      "totalUSD": 15.98
    }
  ],
  "totalUSD": 15.98
}
```

### Order - полная информация

```json
{
  "cartSnapshot": { /* копия cart */ },
  "status": "preparing",
  "payment": {
    "method": "card",
    "status": "paid"
  },
  "delivery": {
    "type": "courier",
    "address": "123 Main St",
    "eta": "2024-01-15T18:30:00Z",
    "tracking": [
      {
        "ts": "2024-01-15T17:00:00Z",
        "status": "preparing",
        "note": "Ваш заказ готовится"
      }
    ]
  }
}
```

---

## Индексы для производительности

### Критические индексы

| Коллекция | Поле | Тип | Цель |
|-----------|------|-----|------|
| **users** | email | unique | Быстрый поиск по email |
| **categories** | slug | unique | URL-friendly поиск |
| **products** | slug | unique | URL-friendly поиск |
| **products** | categoryId | index | Фильтр по категории |
| **products** | tags | index | AI matching |
| **carts** | userId | index | Поиск корзины пользователя |
| **carts** | sessionId | index | Гостевые корзины |
| **orders** | createdAt | index (desc) | Сортировка новых заказов |
| **orders** | status | index | Фильтр по статусу |

---

## Примеры запросов

### Найти все бургеры дешевле $7

```javascript
db.products.find({
  categoryId: burgersCategory._id,
  priceUSD: { $lt: 7 },
  isActive: true
}).sort({ priceUSD: 1 })
```

### Найти продукты для "комфортной еды"

```javascript
db.products.find({
  tags: { $in: ["comfort-food", "indulgent"] },
  isActive: true
})
```

### Создать заказ из корзины

```javascript
// 1. Получить корзину
const cart = await db.carts.findOne({ userId: userId })

// 2. Создать заказ
await db.orders.insertOne({
  userId: userId,
  cartSnapshot: {
    items: cart.items,
    totalUSD: cart.totalUSD
  },
  status: "pending",
  payment: { method: "card", status: "pending" },
  delivery: { type: "pickup", tracking: [] },
  totalUSD: cart.totalUSD,
  createdAt: new Date()
})

// 3. Очистить корзину
await db.carts.updateOne(
  { userId: userId },
  { $set: { items: [], totalUSD: 0 } }
)
```

### AI-рекомендации с fallback

```javascript
// 1. Создать AI сессию
const session = await db.ai_sessions.insertOne({
  userId: userId,
  sessionId: sessionId,
  answers: userAnswers,
  createdAt: new Date()
})

// 2. Попытка получить рекомендации от Gemini
try {
  const aiResponse = await callGeminiAPI(userAnswers)
  await db.ai_sessions.updateOne(
    { _id: session.insertedId },
    { 
      $set: { 
        aiResponse: aiResponse,
        fallbackUsed: false 
      } 
    }
  )
} catch (error) {
  // 3. Fallback на правила
  const rules = await db.mood_rules
    .find({ isActive: true })
    .sort({ priority: -1 })
    .toArray()
  
  const matchedProducts = matchRulesToAnswers(rules, userAnswers)
  
  await db.ai_sessions.updateOne(
    { _id: session.insertedId },
    { 
      $set: { 
        aiResponse: { recommendedProducts: matchedProducts },
        fallbackUsed: true 
      } 
    }
  )
}
```

---

## Масштабирование

### Шардинг (для больших нагрузок)

Рекомендуемые ключи шардинга:

- **orders**: `userId` + `createdAt`
- **carts**: `userId` или `sessionId`
- **ai_sessions**: `userId` + `createdAt`

### Репликация

Для высокой доступности используйте MongoDB Replica Set:

```
Primary (Write) ← → Secondary (Read)
                ← → Secondary (Read)
```

---

**Эта диаграмма помогает визуализировать всю архитектуру базы данных FastSpot! 🎨**

