# Database Schema# MongoDB Schema для FastSpot



MongoDB collections and structure for FastSpot.## Коллекции и их структура



## Collections### 1. users

Хранит информацию о пользователях системы (админы и гости).

### users

User accounts (admins and guests).```javascript

{

```javascript  _id: ObjectId,

{  role: String,              // "admin" | "guest"

  _id: ObjectId,  name: String,

  role: String,              // "admin" or "guest"  email: String,             // уникальный

  name: String,  phone: String,

  email: String,             // unique  passwordHash: String,      // bcrypt hash

  phone: String,  createdAt: Date

  passwordHash: String,      // bcrypt hashed}

  createdAt: Date```

}

```**Индексы:**

- `email` (unique)

**Indexes:** `email` (unique), `role`- `role`



------



### categories### 2. categories

Menu categories.Категории меню (бургеры, напитки, десерты и т.д.).



```javascript```javascript

{{

  _id: ObjectId,  _id: ObjectId,

  name: String,              // "Burgers", "Drinks"  name: String,              // "Бургеры", "Напитки"

  slug: String,              // "burgers", "drinks"    slug: String,              // "burgers", "drinks"

  image: String,             // image URL  image: String,             // URL изображения

  isActive: Boolean  isActive: Boolean          // доступна ли категория

}}

``````



**Indexes:** `slug` (unique), `isActive`**Индексы:**

- `slug` (unique)

---- `isActive`



### products---

Menu items with customization options.

### 3. products

```javascriptТовары меню с ингредиентами, опциями и тегами.

{

  _id: ObjectId,```javascript

  categoryId: ObjectId,      // → categories._id{

  name: String,              // "Big Mac", "Coca-Cola"  _id: ObjectId,

  slug: String,              // "big-mac", "coca-cola"  categoryId: ObjectId,      // ссылка на categories._id

  description: String,  name: String,              // "Биг Мак", "Кока-Кола"

  priceUSD: Number,          // base price  slug: String,              // "big-mac", "coca-cola"

  image: String,  description: String,

  isActive: Boolean,  priceUSD: Number,          // цена в долларах (например, 5.99)

    image: String,             // URL изображения

  ingredients: [             // customizable ingredients  isActive: Boolean,

    {  ingredients: [             // настраиваемые ингредиенты

      key: String,           // "pickles", "onions"    {

      label: String,         // "Pickles"      key: String,           // "pickles", "onions"

      defaultIncluded: Boolean      label: String,         // "Маринованные огурцы"

    }      defaultIncluded: Boolean

  ],    }

    ],

  options: [                 // extra options  options: [                 // дополнительные опции

    {    {

      key: String,           // "size", "sauce"      key: String,           // "size", "sauce"

      label: String,         // "Size", "Sauce"      label: String,         // "Размер", "Соус"

      type: String,          // "single" or "multiple"      type: String,          // "single" | "multiple"

      choices: [      choices: [

        {        {

          value: String,     // "large", "medium"          value: String,     // "large", "medium"

          label: String,     // "Large", "Medium"          label: String,     // "Большой", "Средний"

          extraPriceUSD: Number          extraPriceUSD: Number // доплата в USD (0 если нет)

        }        }

      ]      ]

    }    }

  ],  ],

    tags: [String]             // ["spicy", "vegetarian", "comfort-food"]

  tags: [String]             // ["spicy", "vegetarian"]}

}```

```

**Индексы:**

**Indexes:** `categoryId`, `slug` (unique), `isActive`, `tags`- `categoryId`

- `slug` (unique)

---- `isActive`

- `tags`

### promotions

Special offers and deals.---



```javascript### 4. promotions

{Акции и специальные предложения.

  _id: ObjectId,

  title: String,             // "Summer Sale!"```javascript

  description: String,{

  startsAt: Date,  _id: ObjectId,

  endsAt: Date,  title: String,             // "Летняя распродажа!"

  bannerImage: String,  description: String,

  isActive: Boolean,  startsAt: Date,

  appliesTo: [ObjectId]      // product IDs  endsAt: Date,

}  bannerImage: String,       // URL баннера

```  isActive: Boolean,

  appliesTo: [ObjectId]      // массив productId, на которые действует акция

**Indexes:** `isActive`, `startsAt`, `endsAt`}

```

---

**Индексы:**

### carts- `isActive`

Shopping carts (users and guests).- `startsAt`, `endsAt`



```javascript---

{

  _id: ObjectId,### 5. carts

  userId: ObjectId | null,   // null for guestsКорзины пользователей (как авторизованных, так и гостей).

  sessionId: String,         // for guest sessions

  ```javascript

  items: [{

    {  _id: ObjectId,

      productId: ObjectId,  userId: ObjectId | null,   // null для гостей

      qty: Number,  sessionId: String,         // для гостей - ID сессии

      chosenIngredients: [String],        // ["pickles", "onions"]  items: [

      chosenOptions: {                    // {"size": "large"}    {

        [key: String]: String | [String]      productId: ObjectId,

      },      qty: Number,

      unitPriceUSD: Number,      chosenIngredients: [String],  // ["pickles", "onions"]

      totalUSD: Number         // unitPrice * qty      chosenOptions: {              // {"size": "large", "sauce": "bbq"}

    }        [key: String]: String | [String]

  ],      },

        unitPriceUSD: Number,   // цена за единицу с учетом опций

  totalUSD: Number,            // sum of all items      totalUSD: Number        // unitPriceUSD * qty

  currency: String,            // "USD"    }

  updatedAt: Date  ],

}  totalUSD: Number,           // сумма всех items[].totalUSD

```  currency: String,           // "USD"

  updatedAt: Date

**Indexes:** `userId`, `sessionId`}

```

---

**Индексы:**

### orders- `userId`

Completed orders.- `sessionId`



```javascript---

{

  _id: ObjectId,### 6. orders

  userId: ObjectId | null,Завершенные заказы с полной информацией о доставке и оплате.

  

  cartSnapshot: {              // copy of cart at checkout```javascript

    items: [/* same as cart */],{

    totalUSD: Number  _id: ObjectId,

  },  userId: ObjectId | null,   // null для гостей

    cartSnapshot: {            // копия корзины на момент заказа

  status: String,              // "pending" | "preparing" | "ready" | "delivering" | "completed" | "cancelled"    items: [/* структура как в carts.items */],

      totalUSD: Number

  payment: {  },

    method: String,            // "card" | "applepay" | "cash"  status: String,            // "pending" | "preparing" | "ready" | "delivering" | "completed" | "cancelled"

    status: String             // "pending" | "paid" | "failed"  payment: {

  },    method: String,          // "card" | "applepay" | "googlepay" | "cash"

      status: String           // "pending" | "paid" | "failed"

  delivery: {  },

    type: String,              // "pickup" | "courier"  delivery: {

    address: String,    type: String,            // "pickup" | "courier"

    eta: Date,    address: String,         // адрес (если courier)

    tracking: [    eta: Date,               // ожидаемое время

      {    tracking: [              // история статусов

        ts: Date,      {

        status: String,        ts: Date,

        note: String        status: String,

      }        note: String

    ]      }

  },    ]

    },

  totalUSD: Number,  totalUSD: Number,

  createdAt: Date  createdAt: Date

}}

``````



**Indexes:** `userId`, `status`, `createdAt`**Индексы:**

- `userId`

---- `status`

- `createdAt`

### mood_questions

Quiz questions for AI recommendations.---



```javascript### 7. mood_questions

{Вопросы для "mood quiz" (тест настроения для AI-рекомендаций).

  _id: ObjectId,

  order: Number,               // question order```javascript

  question: String,            // "How are you feeling?"{

  type: String,                // "single" | "multiple" | "scale"  _id: ObjectId,

    order: Number,             // порядок вопроса

  options: [  question: String,          // "Как вы себя чувствуете сегодня?"

    {  type: String,              // "single" | "multiple" | "scale"

      value: String,           // "happy", "tired"  options: [

      label: String,           // "Happy 😊"    {

      weight: Object           // {"energy": 1, "comfort": 0.5}      value: String,         // "happy", "tired", "hungry"

    }      label: String,         // "Счастлив 😊"

  ],      weight: Object         // {"energy": 1, "comfort": 0.5}

      }

  isActive: Boolean  ],

}  isActive: Boolean

```}

```

**Indexes:** `order`, `isActive`

**Индексы:**

---- `order`

- `isActive`

### mood_rules

Fallback rules when AI is unavailable.---



```javascript### 8. mood_rules

{Правила для fallback рекомендаций (если AI недоступен).

  _id: ObjectId,

  name: String,                // "Energy Boost"```javascript

  {

  conditions: {                // matching conditions  _id: ObjectId,

    energy: { $gt: 0.5 },  name: String,              // "Comfort Food Lover"

    comfort: { $lt: 0.3 }  conditions: {              // условия для срабатывания правила

  },    tags: [String],          // требуемые теги в ответах

      minScore: Object         // {"comfort": 3, "energy": 1}

  recommendedProducts: [ObjectId],  },

  priority: Number,  recommendedProducts: [ObjectId], // массив productId

  isActive: Boolean  priority: Number,          // приоритет правила (больше = выше)

}  isActive: Boolean

```}

```

**Indexes:** `priority`, `isActive`

**Индексы:**

---- `priority`

- `isActive`

### ai_sessions

AI recommendation history.---



```javascript### 9. ai_sessions

{История сессий AI-рекомендаций для аналитики.

  _id: ObjectId,

  userId: ObjectId | null,```javascript

  sessionId: String,{

    _id: ObjectId,

  answers: [                   // user's quiz answers  userId: ObjectId | null,

    {  sessionId: String,

      questionId: ObjectId,  answers: [                 // ответы пользователя на quiz

      value: String | [String]    {

    }      questionId: ObjectId,

  ],      selectedOptions: [String]

      }

  aiResponse: {                // AI recommendation  ],

    model: String,             // "gemini-2.0-flash-thinking-exp"  aiResponse: {              // ответ от Gemini API

    products: [ObjectId],    raw: String,             // сырой текст ответа

    reasoning: String    recommendedProducts: [ObjectId],

  },    reasoning: String        // объяснение рекомендации

    },

  fallbackUsed: Boolean,       // true if AI failed  fallbackUsed: Boolean,     // использовались ли правила вместо AI

  createdAt: Date  createdAt: Date

}}

``````



**Indexes:** `userId`, `sessionId`, `createdAt`**Индексы:**

- `userId`

---- `sessionId`

- `createdAt`

## Relationships

---

- `products.categoryId` → `categories._id`

- `carts.userId` → `users._id`## Связи между коллекциями

- `carts.items[].productId` → `products._id`

- `orders.userId` → `users._id````

- `orders.cartSnapshot.items[].productId` → `products._id`users (1) ──────────> (n) carts

- `promotions.appliesTo[]` → `products._id`users (1) ──────────> (n) orders

- `mood_rules.recommendedProducts[]` → `products._id`users (1) ──────────> (n) ai_sessions

- `ai_sessions.answers[].questionId` → `mood_questions._id`

- `ai_sessions.aiResponse.products[]` → `products._id`categories (1) ─────> (n) products


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

