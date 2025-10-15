# FastSpot Deployment Guide

## Deploy to Railway (Recommended)

Railway is a platform that makes it easy to deploy full-stack applications with databases.

### Prerequisites

1. A Railway account (sign up at [railway.app](https://railway.app))
2. GitHub account with FastSpot repository

### Step-by-Step Deployment

#### 1. Create a New Project on Railway

1. Go to [railway.app](https://railway.app) and sign in
2. Click **"New Project"**
3. Select **"Deploy from GitHub repo"**
4. Choose your **FastSpot** repository
5. Railway will automatically detect your configuration

#### 2. Add MongoDB Database

1. In your Railway project, click **"+ New"**
2. Select **"Database"**
3. Choose **"MongoDB"**
4. Railway will automatically create a MongoDB instance and set the `MONGODB_URI` environment variable

#### 3. Configure Environment Variables

Click on your FastSpot service, then go to **"Variables"** tab and add:

```bash
# Required Variables
PORT=3000
GIN_MODE=release
JWT_SECRET=your-super-secret-jwt-key-generate-a-strong-one
JWT_EXPIRATION=24h
GEMINI_API_KEY=your-gemini-api-key-here
PAYMENTS_PROVIDER=stub

# MONGODB_URI is automatically set by Railway when you add MongoDB
# ALLOWED_ORIGINS will be set after first deployment
```

**Important:** Generate a strong JWT_SECRET. You can use:
```bash
openssl rand -base64 32
```

#### 4. Update CORS After First Deploy

After your first deployment:

1. Copy your Railway deployment URL (e.g., `https://fastspot-production.up.railway.app`)
2. Go back to **Variables** and update:
   ```bash
   ALLOWED_ORIGINS=https://fastspot-production.up.railway.app
   ```
3. Railway will automatically redeploy

#### 5. Access Your Application

Your FastSpot application will be available at:
```
https://your-project-name.up.railway.app
```

### What Gets Deployed

- ✅ **Backend API** (Go server)
- ✅ **Frontend** (Vue.js SPA)
- ✅ **Database** (MongoDB)

Railway will:
1. Build your Go backend
2. Build your Vue.js frontend
3. Serve the frontend through the Go server
4. Connect to MongoDB automatically

### Monitoring

- View logs in Railway dashboard
- Monitor resource usage
- Set up custom domains (optional)

---

## Alternative: Deploy Backend and Frontend Separately

### Backend Deployment Options

1. **Railway** (recommended)
2. **Render** - `https://render.com`
3. **Fly.io** - `https://fly.io`
4. **DigitalOcean App Platform** - `https://www.digitalocean.com/products/app-platform`

### Frontend Deployment Options

1. **Vercel** - `https://vercel.com`
2. **Netlify** - `https://netlify.com`
3. **GitHub Pages** (already configured)

If deploying separately, update `VITE_API_URL` in frontend to point to your backend URL.

---

## Environment Variables Reference

### Backend (.env)

| Variable | Description | Example |
|----------|-------------|---------|
| `PORT` | Server port | `3000` |
| `GIN_MODE` | Gin framework mode | `release` |
| `MONGODB_URI` | MongoDB connection string | Auto-set by Railway |
| `JWT_SECRET` | Secret key for JWT tokens | Generate with `openssl rand -base64 32` |
| `JWT_EXPIRATION` | JWT token expiration | `24h` |
| `GEMINI_API_KEY` | Google Gemini AI API key | Get from Google AI Studio |
| `PAYMENTS_PROVIDER` | Payment provider | `stub` (for testing) |
| `ALLOWED_ORIGINS` | CORS allowed origins | Your deployment URL |

### Frontend (.env.production)

| Variable | Description | Value |
|----------|-------------|-------|
| `VITE_API_URL` | Backend API URL | `/api/v1` (same domain) |

---

## Troubleshooting

### Build Fails

- Check Railway logs for specific errors
- Ensure all dependencies are in `go.mod` and `package.json`
- Verify environment variables are set correctly

### Database Connection Issues

- Ensure MongoDB service is running in Railway
- Check `MONGODB_URI` is set correctly
- Verify network connectivity between services

### CORS Errors

- Update `ALLOWED_ORIGINS` to include your deployment URL
- Ensure the URL includes the protocol (`https://`)
- No trailing slash in the origin URL

### Frontend Not Loading

- Check that `npm run build` completes successfully
- Verify static files are in `frontend/dist`
- Check browser console for errors

---

## Local Development

To run locally:

1. **Start MongoDB:**
   ```bash
   docker-compose up -d mongodb
   ```

2. **Start Backend:**
   ```bash
   cd backend
   go run cmd/api/main.go
   ```

3. **Start Frontend:**
   ```bash
   cd frontend
   npm run dev
   ```

---

## Need Help?

- Railway Documentation: `https://docs.railway.app`
- FastSpot Issues: Create an issue on GitHub
