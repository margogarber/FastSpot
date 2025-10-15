# Frontend Architecture

Vue.js 3 SPA with Pinia state management.

## Structure

```
frontend/src/
├── main.js                   # App entry point
├── App.vue                   # Root component
├── router/
│   └── index.js              # Routes + auth guards
├── stores/                   # Pinia stores (state management)
│   ├── auth.js               # Authentication
│   ├── products.js           # Products + admin CRUD
│   ├── categories.js         # Categories
│   ├── promotions.js         # Promotions
│   ├── cart.js               # Cart management
│   ├── orders.js             # Orders
│   └── mood.js               # Mood quiz
├── api/
│   ├── axios.js              # Axios config + interceptors
│   └── index.js              # API methods
├── views/
│   ├── guest/                # Public pages
│   │   ├── HomePage.vue
│   │   ├── MenuPage.vue
│   │   ├── ProductPage.vue
│   │   ├── CartPage.vue
│   │   └── MoodQuizPage.vue
│   └── admin/                # Admin panel
│       ├── DashboardPage.vue
│       ├── products/         # Products CRUD
│       ├── categories/       # Categories CRUD
│       ├── promotions/       # Promotions CRUD
│       ├── orders/           # Orders management
│       └── mood-questions/   # Mood questions CRUD
└── components/
    ├── layout/
    │   ├── GuestLayout.vue   # Guest wrapper (header, footer)
    │   ├── AdminLayout.vue   # Admin wrapper (sidebar)
    │   └── AdminSidebar.vue
    └── ui/                   # Reusable UI components
```

## Routes

### Guest Routes
- `/` - Home page
- `/menu` - Product catalog (with category filters)
- `/menu/:slug` - Product details
- `/cart` - Shopping cart
- `/mood-quiz` - AI mood quiz

### Admin Routes (Protected)
- `/admin/login` - Login page
- `/admin/dashboard` - Analytics dashboard
- `/admin/products` - Products list
- `/admin/products/create` - Create product
- `/admin/products/:id/edit` - Edit product
- `/admin/products/:id` - View product
- Same structure for: categories, promotions, orders, mood-questions

## State Management (Pinia)

### Auth Store (`stores/auth.js`)
- `login(email, password)` - Admin login
- `logout()` - Clear token
- `isAuthenticated` - Computed state

### Products Store (`stores/products.js`)
- `fetchAllProducts(category, search)` - Filter products
- `fetchProductBySlug(slug)` - Get details
- **Admin methods**: `createProduct()`, `updateProduct()`, `deleteProduct()`

### Cart Store (`stores/cart.js`)
- `addToCart(product, quantity)`
- `updateQuantity(id, quantity)`
- `removeFromCart(id)`
- `clearCart()`
- `totalPrice` - Computed total

### Mood Store (`stores/mood.js`)
- `fetchAllQuestions()` - Get active questions
- `getRecommendations(answers)` - Send to AI
- **Admin methods**: `createQuestion()`, `updateQuestion()`, `deleteQuestion()`

## API Layer

### Axios Config (`api/axios.js`)
- **Base URL**: `http://localhost:3000/api/v1`
- **Timeout**: 60 seconds (for AI requests)
- **Interceptors**:
  - Request: Adds `Authorization: Bearer <token>` or `X-Session-ID` header
  - Response: Handles 401 (redirect to login), network errors

### API Methods (`api/index.js`)

All API calls wrapped in named functions:

```javascript
// Products
export const productsAPI = {
  getAll: (category, search) => apiClient.get('/products', { params }),
  getBySlug: (slug) => apiClient.get(`/products/${slug}`),
  // Admin
  create: (data) => apiClient.post('/admin/products', data),
  update: (id, data) => apiClient.put(`/admin/products/${id}`, data),
  delete: (id) => apiClient.delete(`/admin/products/${id}`)
}

// Similar for: categoriesAPI, promotionsAPI, ordersAPI, moodAPI
```

## Key Features

### Route Guards (`router/index.js`)

```javascript
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    // Check JWT token
    if (!authStore.isAuthenticated) {
      next('/admin/login')
    } else {
      next()
    }
  } else {
    next()
  }
})
```

### Session Persistence
- **Admin**: JWT in `localStorage.auth_token`
- **Guest**: Session ID in `localStorage.session_id` (for cart)

### Mood Quiz Flow
1. Fetch questions: `moodStore.fetchAllQuestions()`
2. User answers (checkboxes/radio)
3. Submit: `moodStore.getRecommendations(answers)`
4. Backend calls Gemini AI
5. Display 2-3 recommended products

### Admin CRUD Pattern

Example: Products Create

```vue
<script setup>
import { ref } from 'vue'
import { useProductsStore } from '@/stores/products'
import { useRouter } from 'vue-router'

const store = useProductsStore()
const router = useRouter()
const form = ref({ name: '', slug: '', price: 0, ... })

async function handleSubmit() {
  await store.createProduct(form.value)
  router.push('/admin/products')
}
</script>
```

## Styling

- Custom CSS (no framework)
- Responsive design
- Color scheme: `--primary`, `--success`, `--danger`, `--bg-color`

## Run

```bash
cd frontend
npm install
npm run dev
```

Runs on `http://localhost:5173`

## Build

```bash
npm run build
# Output: dist/
```

## Notes

- All admin routes require authentication
- Cart persists via session ID (guest users)
- Mood quiz accessible to guests (no auth)
- Products filtered by category slug (not ID)
- AI recommendations take 10-30 seconds (Gemini thinking mode)

