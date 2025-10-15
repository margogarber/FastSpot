# FastSpot Deployment Guide

## Frontend Deployment - GitHub Pages

The FastSpot frontend is automatically deployed to GitHub Pages via GitHub Actions.

### Live URL

Your frontend is available at:
```
https://margogarber.github.io/FastSpot/
```

### How It Works

1. **Automatic Deployment**: Every push to the `main` branch triggers a GitHub Actions workflow
2. **Build Process**: The workflow builds the Vue.js frontend using Vite
3. **Deployment**: The built files are deployed to GitHub Pages

### Configuration

The deployment is configured via:
- `.github/workflows/deploy.yml` - GitHub Actions workflow
- `frontend/vite.config.js` - Base path set to `/FastSpot/`

### Viewing Deployment Status

1. Go to your GitHub repository: `https://github.com/margogarber/FastSpot`
2. Click the **"Actions"** tab
3. View the latest workflow run

### Enabling GitHub Pages (First Time Setup)

If GitHub Pages isn't enabled yet:

1. Go to your repository on GitHub
2. Click **Settings** → **Pages** (left sidebar)
3. Under **Source**, select **GitHub Actions**
4. Save and wait for deployment to complete

---

## Backend Deployment (Local Development)

The backend API runs locally for development. To deploy the backend to production, consider:

### Deployment Options for Backend:

1. **Railway** - `https://railway.app` (requires payment method, $5/month with $5 free credit)
2. **Render** - `https://render.com` (free tier available)
3. **Fly.io** - `https://fly.io` (free tier available)
4. **DigitalOcean App Platform** - `https://www.digitalocean.com/products/app-platform`
5. **Heroku** - `https://heroku.com` (paid plans only)

### Connecting Frontend to Production Backend

Once you deploy your backend, update the frontend API URL:

1. Create `frontend/.env.production`:
   ```bash
   VITE_API_URL=https://your-backend-url.com/api/v1
   ```

2. Push changes to trigger a new deployment

---

## Local Development

### Prerequisites

- Go 1.22+
- Node.js 20+
- MongoDB

### Start MongoDB

```bash
docker-compose up -d mongodb
```

### Start Backend

```bash
cd backend
cp .env.example .env
# Edit .env with your configuration
go run cmd/api/main.go
```

Backend will run on: `http://localhost:3000`

### Start Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend will run on: `http://localhost:5173`

---

## Environment Variables

### Backend (.env)

| Variable | Description | Example |
|----------|-------------|---------|
| `PORT` | Server port | `3000` |
| `GIN_MODE` | Gin framework mode | `debug` or `release` |
| `MONGODB_URI` | MongoDB connection string | `mongodb://localhost:27017/fastspot` |
| `JWT_SECRET` | Secret key for JWT tokens | Generate with `openssl rand -base64 32` |
| `JWT_EXPIRATION` | JWT token expiration | `24h` |
| `GEMINI_API_KEY` | Google Gemini AI API key | Get from [Google AI Studio](https://aistudio.google.com/app/apikey) |
| `PAYMENTS_PROVIDER` | Payment provider | `stub` (for testing) |
| `ALLOWED_ORIGINS` | CORS allowed origins | `http://localhost:5173,http://localhost:3000` |

### Frontend (.env)

| Variable | Description | Example |
|----------|-------------|---------|
| `VITE_API_URL` | Backend API URL | `http://localhost:3000/api/v1` |

---

## Troubleshooting

### Frontend Not Loading on GitHub Pages

1. Check the **Actions** tab for deployment errors
2. Verify GitHub Pages is enabled in repository settings
3. Ensure the workflow completed successfully
4. Clear browser cache and try again

### API Errors in Frontend

- The frontend on GitHub Pages will try to connect to `localhost:3000` by default
- You need to deploy the backend separately and update `VITE_API_URL`
- Or run the backend locally while testing the GitHub Pages frontend

### CORS Errors

- Update `ALLOWED_ORIGINS` in backend `.env` to include your GitHub Pages URL:
  ```bash
  ALLOWED_ORIGINS=https://margogarber.github.io
  ```

---

## Project Structure

```
FastSpot/
├── backend/          # Go API server
├── frontend/         # Vue.js application
├── db/              # Database schema and seeds
└── .github/         # GitHub Actions workflows
    └── workflows/
        └── deploy.yml  # Frontend deployment workflow
```

---

## Need Help?

- **GitHub Pages Docs**: https://docs.github.com/en/pages
- **Vite Deployment**: https://vitejs.dev/guide/static-deploy.html
- **FastSpot Issues**: Create an issue on GitHub

