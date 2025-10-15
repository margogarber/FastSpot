/**
 * Axios configuration with JWT interceptors
 */
import axios from 'axios'

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:3000/api/v1',
  headers: {
    'Content-Type': 'application/json'
  },
  timeout: 60000 // 60 seconds for AI recommendations (Gemini thinking mode can take 30+ seconds)
})

// Generate or get session ID for guest users
function getSessionID() {
  let sessionId = localStorage.getItem('session_id')
  if (!sessionId) {
    sessionId = 'guest_' + Math.random().toString(36).substring(2) + Date.now()
    localStorage.setItem('session_id', sessionId)
  }
  return sessionId
}

// Request interceptor - add JWT token or session ID
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('auth_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    } else {
      // For guest users, send session ID
      const sessionId = getSessionID()
      config.headers['X-Session-ID'] = sessionId
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor - handle errors
apiClient.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    if (error.response) {
      // Server responded with error status
      switch (error.response.status) {
        case 401:
          // Unauthorized - только для защищенных роутов
          // Не редиректим если уже на странице логина или если это запрос корзины
          const isLoginPage = window.location.pathname === '/admin/login'
          const isAdminRoute = window.location.pathname.startsWith('/admin')
          
          if (!isLoginPage && isAdminRoute) {
            localStorage.removeItem('auth_token')
            localStorage.removeItem('user')
            window.location.href = '/admin/login'
          }
          break
        case 403:
          console.error('Forbidden:', error.response.data)
          break
        case 404:
          console.error('Not found:', error.response.data)
          break
        case 500:
          console.error('Server error:', error.response.data)
          break
        default:
          console.error('API error:', error.response.data)
      }
    } else if (error.request) {
      // Request was made but no response received
      console.error('Network error:', error.message)
    } else {
      // Something else happened
      console.error('Error:', error.message)
    }
    return Promise.reject(error)
  }
)

export default apiClient

