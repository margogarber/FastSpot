# 📊 Database Structure# 📊 Database Overview# 📊 Database Diagram



A simple visual guide to how FastSpot's database is organized.



---Quick visual guide to FastSpot's database structure.## How Collections Connect



## The Big Picture



```## Main Collections & Relationships```mermaid

┌─────────────┐

│   USERS     │  (Customers & Admins)erDiagram

└──────┬──────┘

       │```mermaid    users ||--o{ carts : "has"

       ├─────────────┐

       │             │erDiagram    users ||--o{ orders : "places"

       ▼             ▼

┌─────────┐    ┌─────────┐    users ||--o{ carts : "has"    users ||--o{ ai_sessions : "creates"

│  CARTS  │    │ ORDERS  │

└────┬────┘    └────┬────┘    users ||--o{ orders : "places"    

     │              │

     │    ┌─────────┴─────────┐    categories ||--o{ products : "contains"    categories ||--o{ products : "contains"

     │    │                   │

     ▼    ▼                   ▼    products }o--o{ carts : "added_to"    

┌──────────────┐       ┌─────────────┐

│   PRODUCTS   │◄──────┤ CATEGORIES  │    products }o--o{ orders : "part_of"    products }o--o{ promotions : "appliesTo"

└──────┬───────┘       └─────────────┘

       │```    products }o--o{ carts : "items"

       ▼

┌──────────────┐    products }o--o{ orders : "cartSnapshot"

│  PROMOTIONS  │

└──────────────┘## Collections    products }o--o{ mood_rules : "recommendedProducts"

```

    products }o--o{ ai_sessions : "recommends"

---

**users** → Admin & guest accounts      

## How It Works

**categories** → Menu sections (Burgers, Drinks, etc.)      mood_questions }o--o{ ai_sessions : "answers"

### 🛒 Shopping Flow

**products** → Menu items with options      

```

1. Browse → 2. Add to Cart → 3. Checkout → 4. Create Order**carts** → Active shopping carts      users {

   │            │               │              │

Categories   Products        Cart           Order**orders** → Completed purchases          ObjectId _id PK

```

**promotions** → Special deals          string role

### 👤 Users Can Be

**mood_questions** → Quiz for AI recommendations          string name

- **Guest** - Browse and order without account (uses session ID)

- **Admin** - Manage menu, view orders, create promotions**ai_sessions** → AI recommendation history        string email UK



---        string phone



## Main Collections Explained## How It Works        string passwordHash



### 1. 📂 Categories        date createdAt

Menu sections that organize products.

### 🛒 Ordering Flow    }

```

Examples:Browse Menu → Add to Cart → Checkout → Order Created    

- Burgers 🍔

- Drinks 🥤    categories {

- Desserts 🍰

- Snacks 🍟### 🤖 AI Recommendations        ObjectId _id PK

```

Take Quiz → AI Analyzes Mood → Suggests Products → Add to Cart        string name

### 2. 🍔 Products

Individual menu items.        string slug UK

        string image

```        boolean isActive

Each product has:    }

- Name & description    

- Price    products {

- Image        ObjectId _id PK

- Category (Burger, Drink, etc.)        ObjectId categoryId FK

- Customization options (size, sauce, extras)        string name

- Tags (vegetarian, spicy, popular)        string slug UK

```        string description

        number priceUSD

**Example:**        string image

```        boolean isActive

Big Mac        array ingredients

├── Base Price: $5.99        array options

├── Category: Burgers        array tags

├── Ingredients: pickles, lettuce, onions (removable)    }

├── Options:    

│   ├── Size: Small, Medium, Large (+$1.00)    promotions {

│   └── Sauce: BBQ, Spicy, Classic        ObjectId _id PK

└── Tags: comfort-food, popular        string title

```        string description

        date startsAt

### 3. 🛒 Carts        date endsAt

Active shopping carts (not checked out yet).        string bannerImage

        boolean isActive

```        array appliesTo

Cart Items:    }

├── Product: Big Mac    

├── Quantity: 2    carts {

├── Customizations:        ObjectId _id PK

│   ├── No pickles        ObjectId userId FK

│   ├── Large size        string sessionId

│   └── BBQ sauce        array items

└── Total: $13.98        number totalUSD

```        string currency

        date updatedAt

Each cart belongs to either:    }

- A logged-in **user** (userId)    

- A **guest** (sessionId)    orders {

        ObjectId _id PK

### 4. 📦 Orders        ObjectId userId FK

Completed purchases.        object cartSnapshot

        string status

```        object payment

Order includes:        object delivery

├── Cart snapshot (what was ordered)        number totalUSD

├── Status (pending → preparing → ready → completed)        date createdAt

├── Payment info (card, ApplePay, cash)    }

├── Delivery details (pickup or delivery)    

└── Total price    mood_questions {

```        ObjectId _id PK

        number order

### 5. 🎉 Promotions        string question

Special deals and discounts.        string type

        array options

```        boolean isActive

Examples:    }

- "Buy 1 Get 1 Free Burgers"    

- "20% Off All Drinks"    mood_rules {

- "Combo Deal: Burger + Fries + Drink"        ObjectId _id PK

```        string name

        object conditions

Each promotion:        array recommendedProducts

- Has start and end dates        number priority

- Applies to specific products        boolean isActive

- Can be active or inactive    }

    

### 6. 😊 Mood Quiz & AI    ai_sessions {

        ObjectId _id PK

**mood_questions** - Quiz to understand customer mood        ObjectId userId FK

```        string sessionId

Question: "How are you feeling?"        array answers

Answers:        object aiResponse

- Happy 😊        boolean fallbackUsed

- Tired 😴        date createdAt

- Hungry 😋    }

- Stressed 😰```

```

---

**ai_sessions** - AI recommends products based on mood

```## Основные потоки данных

User takes quiz → AI analyzes → Suggests products

Example: Feeling stressed → Recommends comfort food### 1. 🛒 Процесс заказа

```

```mermaid

---flowchart LR

    A[User] -->|browse| B[Products]

## Relationships    B -->|add to| C[Cart]

    C -->|checkout| D[Order]

### One-to-Many    D -->|snapshot| C

    

```    E[Categories] -->|organize| B

1 Category → Many Products    F[Promotions] -.->|discount| B

└─ "Burgers" category contains:```

   - Big Mac

   - Cheeseburger### 2. 🤖 AI-рекомендации

   - Veggie Burger

``````mermaid

flowchart TD

```    A[User начинает Quiz] --> B[Mood Questions]

1 User → Many Orders    B --> C[User отвечает]

└─ Customer "John" has placed:    C --> D[AI Session создан]

   - Order #1 (yesterday)    D --> E{Gemini API доступен?}

   - Order #2 (today)    E -->|Да| F[Gemini анализирует]

```    E -->|Нет| G[Mood Rules применяются]

    F --> H[Рекомендации]

### Many-to-Many    G --> H

    H --> I[Products добавлены в Cart]

``````

Products ↔ Promotions

├─ "Buy 1 Get 1" promotion applies to:### 3. 👤 Типы пользователей

│  - Big Mac

│  - Cheeseburger```mermaid

└─ "Big Mac" can be in multiple promotionsflowchart LR

```    A[User] --> B{Авторизован?}

    B -->|Да| C[userId]

---    B -->|Нет| D[sessionId]

    C --> E[Cart/Orders привязаны к userId]

## Data Flow Examples    D --> F[Cart привязана к sessionId]

    

### Example 1: Customer Makes an Order    style C fill:#90EE90

    style D fill:#FFB6C1

``````

Step 1: Customer browses Menu

        ↓---

     Products (filtered by Category)

        ## Детальная структура коллекций

Step 2: Customer adds items to Cart

        ↓### Products - самая сложная структура

     Cart (stores productId + customizations)

        ```mermaid

Step 3: Customer checks outgraph TB

        ↓    A[Product] --> B[Basic Info]

     Order created (with cart snapshot)    A --> C[Pricing]

        ↓    A --> D[Customization]

     Cart is cleared    A --> E[Relations]

```    

    B --> B1[name, slug, description]

### Example 2: AI Recommendation    B --> B2[image, isActive]

    

```    C --> C1[priceUSD]

Step 1: Customer takes Mood Quiz    C --> C2[currency: USD]

        ↓    

     mood_questions shown    D --> D1[Ingredients]

            D --> D2[Options]

Step 2: Customer answers    D --> D3[Tags]

        ↓    

     Answers saved in ai_sessions    D1 --> D1A[pickles, onions, lettuce]

            D1 --> D1B[defaultIncluded: true/false]

Step 3: AI analyzes mood    

        ↓    D2 --> D2A[Size: small/medium/large]

     AI recommends Products    D2 --> D2B[Sauce: bbq/spicy/classic]

        ↓    D2 --> D2C[extraPriceUSD per choice]

     Products added to Cart    

```    D3 --> D3A[comfort-food, healthy]

    D3 --> D3B[vegetarian, spicy]

---    D3 --> D3C[для AI matching]

    

## Field Types Quick Reference    E --> E1[categoryId → Categories]

    E --> E2[используется в Carts]

| Type | Example | Used For |    E --> E3[используется в Orders]

|------|---------|----------|    E --> E4[используется в Promotions]

| **String** | "Big Mac" | Names, descriptions, text |```

| **Number** | 5.99 | Prices, quantities |

| **Boolean** | true/false | Active/inactive, yes/no |### Cart Items - структура элемента корзины

| **Date** | 2024-01-15 | Created dates, timestamps |

| **ObjectId** | 507f1f77... | Unique IDs, references |```json

| **Array** | ["tag1", "tag2"] | Lists, multiple values |{

| **Object** | {name: "...", price: ...} | Complex data structures |  "items": [

    {

---      "productId": "ObjectId",

      "qty": 2,

## Quick Stats      "chosenIngredients": ["pickles", "onions"],

      "chosenOptions": {

After running `seed.js`:        "size": "large",

        "sauce": "bbq"

```      },

📊 Database contains:      "unitPriceUSD": 7.99,

├─ 1 admin user      "totalUSD": 15.98

├─ 4 categories    }

├─ 10 products  ],

├─ 3 promotions  "totalUSD": 15.98

├─ 8 quiz questions}

└─ 0 orders (starts empty)```

```

### Order - полная информация

---

```json

## Visual Schema{

  "cartSnapshot": { /* копия cart */ },

```  "status": "preparing",

users {  "payment": {

  _id: ObjectId    "method": "card",

  role: "admin" | "guest"    "status": "paid"

  email: "admin@local"  },

  passwordHash: "..."  "delivery": {

  name: "Admin"    "type": "courier",

}    "address": "123 Main St",

    "eta": "2024-01-15T18:30:00Z",

categories {    "tracking": [

  _id: ObjectId      {

  name: "Burgers"        "ts": "2024-01-15T17:00:00Z",

  slug: "burgers"        "status": "preparing",

  image: "url"        "note": "Ваш заказ готовится"

  isActive: true      }

}    ]

  }

products {}

  _id: ObjectId```

  categoryId: → categories._id

  name: "Big Mac"---

  slug: "big-mac"

  priceUSD: 5.99## Индексы для производительности

  image: "url"

  ingredients: [...]### Критические индексы

  options: [...]

  tags: ["comfort-food"]| Коллекция | Поле | Тип | Цель |

  isActive: true|-----------|------|-----|------|

}| **users** | email | unique | Быстрый поиск по email |

| **categories** | slug | unique | URL-friendly поиск |

carts {| **products** | slug | unique | URL-friendly поиск |

  _id: ObjectId| **products** | categoryId | index | Фильтр по категории |

  userId: → users._id (or null for guests)| **products** | tags | index | AI matching |

  sessionId: "guest_abc123"| **carts** | userId | index | Поиск корзины пользователя |

  items: [| **carts** | sessionId | index | Гостевые корзины |

    {| **orders** | createdAt | index (desc) | Сортировка новых заказов |

      productId: → products._id| **orders** | status | index | Фильтр по статусу |

      qty: 2

      chosenIngredients: ["pickles"]---

      chosenOptions: {"size": "large"}

      totalUSD: 13.98## Примеры запросов

    }

  ]### Найти все бургеры дешевле $7

  totalUSD: 13.98

}```javascript

db.products.find({

orders {  categoryId: burgersCategory._id,

  _id: ObjectId  priceUSD: { $lt: 7 },

  userId: → users._id (or null)  isActive: true

  cartSnapshot: {...}}).sort({ priceUSD: 1 })

  status: "pending"```

  payment: {method: "card", status: "paid"}

  delivery: {type: "pickup", eta: Date}### Найти продукты для "комфортной еды"

  totalUSD: 13.98

  createdAt: Date```javascript

}db.products.find({

  tags: { $in: ["comfort-food", "indulgent"] },

promotions {  isActive: true

  _id: ObjectId})

  title: "Buy 1 Get 1 Free"```

  startsAt: Date

  endsAt: Date### Создать заказ из корзины

  appliesTo: [→ products._id]

  isActive: true```javascript

}// 1. Получить корзину

```const cart = await db.carts.findOne({ userId: userId })



---// 2. Создать заказ

await db.orders.insertOne({

**That's the complete structure! Simple, right? 🎯**  userId: userId,

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

