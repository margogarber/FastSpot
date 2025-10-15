# FastSpot

**FastSpot** is a modern fast-food ordering system with AI-powered mood-based recommendations.

## What It Does

- **Guest Ordering**: Browse menu, filter by categories, add items to cart, place orders
- **Mood Quiz**: AI analyzes your mood and recommends food items using Gemini API
- **Admin Panel**: Full CRUD for products, categories, promotions, orders, and mood questions
- **Real-time Analytics**: Dashboard with sales statistics and charts

## Tech Stack

- **Frontend**: Vue.js 3 + Pinia + Vue Router
- **Backend**: Go (Gin framework)
- **Database**: MongoDB
- **AI**: Google Gemini 2.5 Flash

## Quick Start

### Prerequisites
- Go 1.21+
- Node.js 18+
- MongoDB running on `localhost:27017`

### 1. Setup Database

```bash
# Seed database with sample data
cd db
node seed.js
```

### 2. Start Backend

```bash
cd backend
go run cmd/api/main.go
```

Backend runs on `http://localhost:3000`

### 3. Start Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend runs on `http://localhost:5173`

### 4. Access

- **Guest**: `http://localhost:5173` - browse menu, take mood quiz, order
- **Admin**: `http://localhost:5173/admin/login` - manage everything
  - Email: `admin@fastspot.com`
  - Password: `admin123`

## Stop

```bash
# Kill backend (from project root)
lsof -ti:3000 | xargs kill -9

# Kill frontend (Ctrl+C in terminal)
```

## Environment

Create `backend/.env`:

```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=fastspot
JWT_SECRET=your-secret-key-change-in-production
GEMINI_API_KEY=your-gemini-api-key
```

Get Gemini API key: https://aistudio.google.com/app/apikey

---

For detailed architecture see `BACKEND.md`, `FRONTEND.md`, `DATABASE.md`
