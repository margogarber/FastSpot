# Backend Architecture

Go-based REST API using Gin framework.

## Structure

```
backend/
├── cmd/api/main.go           # Entry point, routes, CORS
├── internal/
│   ├── handlers/             # HTTP handlers (controllers)
│   │   └── handlers.go       # All CRUD logic
│   ├── models/               # Data structures
│   │   ├── user.go
│   │   ├── product.go
│   │   ├── category.go
│   │   ├── promotion.go
│   │   ├── order.go
│   │   ├── cart.go
│   │   └── mood.go
│   ├── repository/           # Database layer
│   │   └── repository.go     # MongoDB operations
│   ├── middleware/           # JWT auth, CORS
│   │   └── auth.go
│   ├── services/             # Business logic
│   │   └── ai/
│   │       └── gemini.go     # Gemini AI integration
│   └── config/               # Configuration
│       └── config.go
└── .env                      # Environment variables
```

## API Endpoints

### Public (Guest)
- `GET /api/v1/products` - List products (filter by category, search)
- `GET /api/v1/products/:slug` - Get product details
- `GET /api/v1/categories` - List categories
- `POST /api/v1/cart` - Add to cart
- `GET /api/v1/cart` - Get cart items
- `PUT /api/v1/cart/:id` - Update cart item
- `DELETE /api/v1/cart/:id` - Remove from cart
- `POST /api/v1/orders` - Create order
- `GET /api/v1/mood/questions` - Get active mood questions
- `POST /api/v1/mood/recommend` - Get AI recommendations

### Admin (Auth Required)
- `POST /api/v1/admin/login` - Login (returns JWT)
- `GET /api/v1/admin/analytics` - Dashboard stats
- **Products**: GET, POST, PUT, DELETE `/api/v1/admin/products`
- **Categories**: GET, POST, PUT, DELETE `/api/v1/admin/categories`
- **Promotions**: GET, POST, PUT, DELETE `/api/v1/admin/promotions`
- **Orders**: GET, PUT `/api/v1/admin/orders` (no delete, status update only)
- **Mood Questions**: GET, POST, PUT, DELETE `/api/v1/admin/mood-questions`

## Key Features

### Authentication
- JWT tokens for admin
- Session IDs for guest users (cart persistence)
- Middleware: `middleware.AuthMiddleware(config.JWTSecret)`

### AI Recommendations
- **Service**: `services/ai/gemini.go`
- **Model**: Gemini 2.5 Flash (with extended thinking)
- **Process**:
  1. User answers mood questions
  2. Backend sends answers + product catalog to Gemini
  3. AI returns 2-3 product slugs + reasoning
  4. Frontend displays recommendations

### CORS
- Allows `http://localhost:5173` (frontend)
- Headers: `Authorization`, `Content-Type`, `X-Session-ID`

### Error Handling
- Consistent JSON responses: `{"error": "message"}` or `{"success": true, "data": {...}}`
- HTTP status codes: 200, 201, 400, 401, 404, 500

## Environment Variables

```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=fastspot
JWT_SECRET=your-secret-key
GEMINI_API_KEY=your-gemini-api-key
PORT=3000
```

## Database Layer

**Repository pattern**: All DB operations in `repository/repository.go`

Each entity has:
- `Create(ctx, model)` - Insert
- `FindAll(ctx, filter)` - Query with filter
- `FindByID(ctx, id)` - Get by ObjectID
- `FindBySlug(ctx, slug)` - Get by slug (products, categories)
- `Update(ctx, id, model)` - Update
- `Delete(ctx, id)` - Soft/hard delete

## Run

```bash
cd backend
go run cmd/api/main.go
```

Server starts on `:3000`

## Notes

- All responses include `success: true/false`
- MongoDB ObjectIDs used for relations
- Admin routes protected by JWT
- Guest routes use session IDs for cart
- AI timeout: 60 seconds (Gemini thinking mode is slow)

