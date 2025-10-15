# FastSpot Database

Simple guide to set up and use the FastSpot MongoDB database.

## Requirements
- Node.js >= 16
- MongoDB >= 5.0
- npm or yarn

## Quick Setup
1. Install deps:
```bash
npm install
```
2. Create `.env` with:
```env
MONGODB_URI=mongodb://admin:Admin123!@localhost:27017/fastspot?authSource=admin
```
3. Start MongoDB (local or Docker) and seed data:
```bash
# local (macOS)
brew services start mongodb-community

# or Docker
docker-compose up -d mongodb

# then
node seed.js
```

## What’s Included
- 1 admin user (admin@local / Admin123!)
- Categories: Burgers, Drinks, Desserts, Snacks
- Sample products, promotions, mood questions, and AI session history

## Main Collections
- users, categories, products, carts, orders, promotions, mood_questions, mood_rules, ai_sessions

## Useful Commands
```bash
mongosh fastspot
node seed.js
db.products.find().pretty()
```

## Troubleshooting
- "Can't connect": ensure MongoDB is running and MONGODB_URI is correct.
- "Authentication failed": verify credentials and authSource.
- Check port: `lsof -i :27017`

## Security
- Do not commit `.env`
- Use strong passwords and restricted DB users in production

