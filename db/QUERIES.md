# 📝 Useful Database Queries# 📝 MongoDB Queries - Полезные примеры



Handy MongoDB queries for working with FastSpot data.Коллекция практических MongoDB запросов для FastSpot.



## Finding Products---



```javascript## 🔍 Поиск и фильтрация

// All active products

db.products.find({ isActive: true })### Продукты



// Search by category```javascript

db.products.find({ // Найти все активные бургеры

  categoryId: ObjectId("..."),db.products.find({

  isActive: true   categoryId: ObjectId("..."),

})  isActive: true

})

// Search by tags (vegetarian)

db.products.find({ // Поиск по тегам (вегетарианское)

  tags: "vegetarian",db.products.find({

  isActive: true   tags: "vegetarian",

})  isActive: true

})

// Price filter

db.products.find({// Поиск по нескольким тегам (OR)

  priceUSD: { $gte: 5, $lte: 10 }db.products.find({

}).sort({ priceUSD: 1 })  tags: { $in: ["comfort-food", "healthy"] },

  isActive: true

// Text search})

db.products.find({

  name: { $regex: /burger/i }// Поиск по нескольким тегам (AND)

})db.products.find({

```  tags: { $all: ["comfort-food", "popular"] },

  isActive: true

## Orders})



```javascript// Фильтр по цене

// User's ordersdb.products.find({

db.orders.find({   priceUSD: { $gte: 5, $lte: 10 },

  userId: ObjectId("...")   isActive: true

}).sort({ createdAt: -1 })}).sort({ priceUSD: 1 })



// Active orders (not completed)// Текстовый поиск по названию (регулярное выражение)

db.orders.find({db.products.find({

  status: { $in: ["pending", "preparing", "delivering"] }  name: { $regex: /burger/i },

})  isActive: true

})

// Today's orders

db.orders.find({// Продукты с определенным ингредиентом

  createdAt: { $gte: new Date().setHours(0,0,0,0) }db.products.find({

})  "ingredients.key": "pickles",

  isActive: true

// Paid orders})

db.orders.find({ 

  "payment.status": "paid" // Продукты с опцией "размер"

})db.products.find({

```  "options.key": "size",

  isActive: true

## Shopping Carts})

```

```javascript

// User's cart### Категории

db.carts.findOne({ userId: ObjectId("...") })

```javascript

// Guest cart// Найти категорию по slug

db.carts.findOne({ sessionId: "guest_abc123" })db.categories.findOne({

  slug: "burgers",

// Carts with items  isActive: true

db.carts.find({})

  items: { $ne: [] },

  totalUSD: { $gt: 0 }// Все активные категории

})db.categories.find({

```  isActive: true

}).sort({ name: 1 })

## Active Promotions```



```javascript### Заказы

// Current promotions

const now = new Date()```javascript

db.promotions.find({// Заказы пользователя

  isActive: true,db.orders.find({

  startsAt: { $lte: now },  userId: ObjectId("...")

  endsAt: { $gte: now }}).sort({ createdAt: -1 })

})

```// Активные заказы (не завершены)

db.orders.find({

## Analytics  status: { $in: ["pending", "preparing", "delivering"] }

}).sort({ createdAt: -1 })

```javascript

// Products by category// Заказы за сегодня

db.products.aggregate([const today = new Date()

  { $match: { isActive: true } },today.setHours(0, 0, 0, 0)

  { $group: {

    _id: "$categoryId",db.orders.find({

    count: { $sum: 1 },  createdAt: { $gte: today }

    avgPrice: { $avg: "$priceUSD" }})

  }}

])// Заказы за период

db.orders.find({

// Top ordered products  createdAt: {

db.orders.aggregate([    $gte: new Date("2024-01-01"),

  { $match: { status: "completed" } },    $lte: new Date("2024-01-31")

  { $unwind: "$cartSnapshot.items" },  }

  { $group: {})

    _id: "$cartSnapshot.items.productId",

    totalOrdered: { $sum: "$cartSnapshot.items.qty" }// Заказы на доставку

  }},db.orders.find({

  { $sort: { totalOrdered: -1 } },  "delivery.type": "courier"

  { $limit: 10 }})

])

// Оплаченные заказы

// Total revenuedb.orders.find({

db.orders.aggregate([  "payment.status": "paid"

  { $match: { })

    status: "completed",```

    "payment.status": "paid" 

  }},### Корзины

  { $group: {

    _id: null,```javascript

    totalRevenue: { $sum: "$totalUSD" }// Корзина пользователя

  }}db.carts.findOne({

])  userId: ObjectId("...")

```})


// Корзина гостя
db.carts.findOne({
  sessionId: "guest_abc123"
})

// Непустые корзины
db.carts.find({
  items: { $ne: [] },
  totalUSD: { $gt: 0 }
})

// Корзины с определенным продуктом
db.carts.find({
  "items.productId": ObjectId("...")
})
```

### Акции

```javascript
// Активные акции на данный момент
const now = new Date()

db.promotions.find({
  isActive: true,
  startsAt: { $lte: now },
  endsAt: { $gte: now }
})

// Акции для конкретного продукта
db.promotions.find({
  appliesTo: ObjectId("..."),
  isActive: true
})
```

---

## 📊 Агрегация

### Статистика продуктов

```javascript
// Количество продуктов по категориям
db.products.aggregate([
  { $match: { isActive: true } },
  { $group: {
    _id: "$categoryId",
    count: { $sum: 1 },
    avgPrice: { $avg: "$priceUSD" },
    minPrice: { $min: "$priceUSD" },
    maxPrice: { $max: "$priceUSD" }
  }},
  { $lookup: {
    from: "categories",
    localField: "_id",
    foreignField: "_id",
    as: "category"
  }},
  { $unwind: "$category" },
  { $project: {
    categoryName: "$category.name",
    count: 1,
    avgPrice: { $round: ["$avgPrice", 2] },
    minPrice: 1,
    maxPrice: 1
  }}
])
```

### Топ продукты в заказах

```javascript
db.orders.aggregate([
  { $match: { status: "completed" } },
  { $unwind: "$cartSnapshot.items" },
  { $group: {
    _id: "$cartSnapshot.items.productId",
    totalQuantity: { $sum: "$cartSnapshot.items.qty" },
    totalRevenue: { $sum: "$cartSnapshot.items.totalUSD" },
    orderCount: { $sum: 1 }
  }},
  { $lookup: {
    from: "products",
    localField: "_id",
    foreignField: "_id",
    as: "product"
  }},
  { $unwind: "$product" },
  { $project: {
    productName: "$product.name",
    productImage: "$product.image",
    totalQuantity: 1,
    totalRevenue: { $round: ["$totalRevenue", 2] },
    orderCount: 1
  }},
  { $sort: { totalQuantity: -1 } },
  { $limit: 10 }
])
```

### Выручка по дням

```javascript
db.orders.aggregate([
  { $match: { 
    status: "completed",
    createdAt: { $gte: new Date("2024-01-01") }
  }},
  { $group: {
    _id: {
      year: { $year: "$createdAt" },
      month: { $month: "$createdAt" },
      day: { $dayOfMonth: "$createdAt" }
    },
    totalRevenue: { $sum: "$totalUSD" },
    orderCount: { $sum: 1 },
    avgOrderValue: { $avg: "$totalUSD" }
  }},
  { $sort: { "_id.year": 1, "_id.month": 1, "_id.day": 1 } },
  { $project: {
    date: {
      $dateFromParts: {
        year: "$_id.year",
        month: "$_id.month",
        day: "$_id.day"
      }
    },
    totalRevenue: { $round: ["$totalRevenue", 2] },
    orderCount: 1,
    avgOrderValue: { $round: ["$avgOrderValue", 2] }
  }}
])
```

### Статистика AI-рекомендаций

```javascript
db.ai_sessions.aggregate([
  { $group: {
    _id: null,
    totalSessions: { $sum: 1 },
    aiUsed: {
      $sum: { $cond: [{ $eq: ["$fallbackUsed", false] }, 1, 0] }
    },
    fallbackUsed: {
      $sum: { $cond: [{ $eq: ["$fallbackUsed", true] }, 1, 0] }
    }
  }},
  { $project: {
    _id: 0,
    totalSessions: 1,
    aiUsed: 1,
    fallbackUsed: 1,
    aiSuccessRate: {
      $round: [
        { $divide: ["$aiUsed", "$totalSessions"] },
        2
      ]
    }
  }}
])
```

### Популярные теги в заказах

```javascript
db.orders.aggregate([
  { $match: { status: "completed" } },
  { $unwind: "$cartSnapshot.items" },
  { $lookup: {
    from: "products",
    localField: "cartSnapshot.items.productId",
    foreignField: "_id",
    as: "product"
  }},
  { $unwind: "$product" },
  { $unwind: "$product.tags" },
  { $group: {
    _id: "$product.tags",
    count: { $sum: "$cartSnapshot.items.qty" }
  }},
  { $sort: { count: -1 } },
  { $limit: 10 }
])
```

---

## ✏️ Создание и обновление

### Создать корзину

```javascript
db.carts.insertOne({
  userId: ObjectId("..."),
  sessionId: null,
  items: [],
  totalUSD: 0,
  currency: "USD",
  updatedAt: new Date()
})
```

### Добавить товар в корзину

```javascript
// Найти продукт и рассчитать цену
const product = db.products.findOne({ _id: ObjectId("...") })
const extraPrice = 2.50 // от опций
const unitPrice = product.priceUSD + extraPrice
const qty = 1

// Добавить в корзину
db.carts.updateOne(
  { userId: ObjectId("...") },
  {
    $push: {
      items: {
        productId: product._id,
        qty: qty,
        chosenIngredients: ["pickles", "onions"],
        chosenOptions: { size: "large", sauce: "bbq" },
        unitPriceUSD: unitPrice,
        totalUSD: unitPrice * qty
      }
    },
    $inc: { totalUSD: unitPrice * qty },
    $set: { updatedAt: new Date() }
  }
)
```

### Обновить количество в корзине

```javascript
db.carts.updateOne(
  { 
    userId: ObjectId("..."),
    "items.productId": ObjectId("...")
  },
  {
    $set: {
      "items.$.qty": 3,
      "items.$.totalUSD": 3 * 7.99,
      updatedAt: new Date()
    }
  }
)

// Пересчитать общую сумму
const cart = db.carts.findOne({ userId: ObjectId("...") })
const total = cart.items.reduce((sum, item) => sum + item.totalUSD, 0)

db.carts.updateOne(
  { _id: cart._id },
  { $set: { totalUSD: total } }
)
```

### Удалить товар из корзины

```javascript
// Получить стоимость удаляемого товара
const cart = db.carts.findOne({ userId: ObjectId("...") })
const item = cart.items.find(i => i.productId.equals(ObjectId("...")))

// Удалить и пересчитать
db.carts.updateOne(
  { userId: ObjectId("...") },
  {
    $pull: { items: { productId: ObjectId("...") } },
    $inc: { totalUSD: -item.totalUSD },
    $set: { updatedAt: new Date() }
  }
)
```

### Создать заказ из корзины

```javascript
// 1. Получить корзину
const cart = db.carts.findOne({ userId: ObjectId("...") })

// 2. Создать заказ
db.orders.insertOne({
  userId: cart.userId,
  cartSnapshot: {
    items: cart.items,
    totalUSD: cart.totalUSD
  },
  status: "pending",
  payment: {
    method: "card",
    status: "pending"
  },
  delivery: {
    type: "courier",
    address: "123 Main St",
    eta: new Date(Date.now() + 45 * 60000), // +45 минут
    tracking: [
      {
        ts: new Date(),
        status: "pending",
        note: "Заказ принят"
      }
    ]
  },
  totalUSD: cart.totalUSD,
  createdAt: new Date()
})

// 3. Очистить корзину
db.carts.updateOne(
  { userId: cart.userId },
  {
    $set: {
      items: [],
      totalUSD: 0,
      updatedAt: new Date()
    }
  }
)
```

### Обновить статус заказа

```javascript
db.orders.updateOne(
  { _id: ObjectId("...") },
  {
    $set: { status: "preparing" },
    $push: {
      "delivery.tracking": {
        ts: new Date(),
        status: "preparing",
        note: "Ваш заказ готовится на кухне"
      }
    }
  }
)
```

### Создать AI-сессию

```javascript
db.ai_sessions.insertOne({
  userId: ObjectId("..."),
  sessionId: "session_abc123",
  answers: [
    {
      questionId: ObjectId("..."),
      selectedOptions: ["tired", "comfort"]
    }
  ],
  aiResponse: {
    raw: "На основе ваших ответов...",
    recommendedProducts: [
      ObjectId("..."),
      ObjectId("...")
    ],
    reasoning: "Вы выбрали комфортную еду..."
  },
  fallbackUsed: false,
  createdAt: new Date()
})
```

---

## 🗑️ Удаление

### Удалить старые гостевые корзины

```javascript
// Корзины старше 7 дней
const weekAgo = new Date()
weekAgo.setDate(weekAgo.getDate() - 7)

db.carts.deleteMany({
  userId: null,
  sessionId: { $exists: true },
  updatedAt: { $lt: weekAgo }
})
```

### Удалить завершенные старые заказы (архивация)

```javascript
// Заказы старше 90 дней
const threeMonthsAgo = new Date()
threeMonthsAgo.setDate(threeMonthsAgo.getDate() - 90)

db.orders.deleteMany({
  status: "completed",
  createdAt: { $lt: threeMonthsAgo }
})
```

### Деактивировать продукт (мягкое удаление)

```javascript
db.products.updateOne(
  { _id: ObjectId("...") },
  { $set: { isActive: false } }
)
```

---

## 📈 Аналитика и отчеты

### Средний чек

```javascript
db.orders.aggregate([
  { $match: { status: "completed" } },
  { $group: {
    _id: null,
    avgOrderValue: { $avg: "$totalUSD" },
    minOrderValue: { $min: "$totalUSD" },
    maxOrderValue: { $max: "$totalUSD" },
    totalOrders: { $sum: 1 },
    totalRevenue: { $sum: "$totalUSD" }
  }}
])
```

### Конверсия корзины в заказ

```javascript
// Общее количество корзин
const totalCarts = db.carts.countDocuments({ items: { $ne: [] } })

// Количество заказов за период
const totalOrders = db.orders.countDocuments({
  createdAt: { $gte: new Date("2024-01-01") }
})

const conversionRate = (totalOrders / totalCarts * 100).toFixed(2)
print(`Conversion rate: ${conversionRate}%`)
```

### Популярное время заказов

```javascript
db.orders.aggregate([
  { $match: { status: "completed" } },
  { $group: {
    _id: { $hour: "$createdAt" },
    orderCount: { $sum: 1 },
    avgRevenue: { $avg: "$totalUSD" }
  }},
  { $sort: { _id: 1 } },
  { $project: {
    hour: "$_id",
    orderCount: 1,
    avgRevenue: { $round: ["$avgRevenue", 2] }
  }}
])
```

### Среднее время доставки

```javascript
db.orders.aggregate([
  { $match: { 
    status: "completed",
    "delivery.type": "courier"
  }},
  { $project: {
    deliveryTime: {
      $divide: [
        { $subtract: [
          { $last: "$delivery.tracking.ts" },
          "$createdAt"
        ]},
        1000 * 60 // в минутах
      ]
    }
  }},
  { $group: {
    _id: null,
    avgDeliveryTime: { $avg: "$deliveryTime" },
    minDeliveryTime: { $min: "$deliveryTime" },
    maxDeliveryTime: { $max: "$deliveryTime" }
  }}
])
```

---

## 🔧 Полезные операции

### Пересчитать все цены в корзинах

```javascript
db.carts.find().forEach(cart => {
  let newTotal = 0
  
  cart.items.forEach(item => {
    newTotal += item.totalUSD
  })
  
  db.carts.updateOne(
    { _id: cart._id },
    { $set: { totalUSD: newTotal } }
  )
})
```

### Добавить новый тег всем бургерам

```javascript
db.products.updateMany(
  { categoryId: ObjectId("...") }, // burgers category
  { $addToSet: { tags: "savory" } }
)
```

### Изменить цену продукта

```javascript
db.products.updateOne(
  { _id: ObjectId("...") },
  { $set: { priceUSD: 6.99 } }
)
```

### Backup коллекции в JSON

```bash
# В терминале
mongoexport \
  --db=fastspot \
  --collection=products \
  --out=products_backup.json \
  --pretty
```

### Restore коллекции из JSON

```bash
# В терминале
mongoimport \
  --db=fastspot \
  --collection=products \
  --file=products_backup.json \
  --drop
```

---

## 🔍 Полезные проверки

### Найти продукты без изображений

```javascript
db.products.find({
  $or: [
    { image: { $exists: false } },
    { image: "" },
    { image: null }
  ]
})
```

### Найти заказы без tracking

```javascript
db.orders.find({
  "delivery.tracking": { $size: 0 }
})
```

### Найти корзины с недействительными продуктами

```javascript
db.carts.find().forEach(cart => {
  cart.items.forEach(item => {
    const product = db.products.findOne({ _id: item.productId })
    if (!product || !product.isActive) {
      print(`Invalid product in cart ${cart._id}: ${item.productId}`)
    }
  })
})
```

---

**Эти запросы помогут эффективно работать с MongoDB в FastSpot! 🚀**

