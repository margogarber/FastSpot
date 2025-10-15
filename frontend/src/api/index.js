/**
 * API service methods for all entities
 */
import apiClient from './axios'

// ==================== AUTH ====================
export const authAPI = {
  // Guest session
  createGuestSession() {
    return apiClient.post('/auth/guest')
  },
  
  // Admin login
  login(credentials) {
    return apiClient.post('/auth/login', credentials)
  },
  
  // Logout
  logout() {
    return apiClient.post('/auth/logout')
  },
  
  // Get current user
  me() {
    return apiClient.get('/users/me')
  }
}

// ==================== CATEGORIES ====================
export const categoriesAPI = {
  getAll(params = {}) {
    return apiClient.get('/categories', { params })
  },
  
  getBySlug(slug) {
    return apiClient.get(`/categories/${slug}`)
  },
  
  getById(id) {
    return apiClient.get(`/admin/categories/${id}`)
  },
  
  create(data) {
    return apiClient.post('/admin/categories', data)
  },
  
  update(id, data) {
    return apiClient.put(`/admin/categories/${id}`, data)
  },
  
  delete(id) {
    return apiClient.delete(`/admin/categories/${id}`)
  }
}

// ==================== PRODUCTS ====================
export const productsAPI = {
  getAll(params = {}) {
    return apiClient.get('/products', { params })
  },
  
  getBySlug(slug) {
    return apiClient.get(`/products/${slug}`)
  },
  
  getById(id) {
    return apiClient.get(`/admin/products/${id}`)
  },
  
  create(data) {
    return apiClient.post('/admin/products', data)
  },
  
  update(id, data) {
    return apiClient.put(`/admin/products/${id}`, data)
  },
  
  delete(id) {
    return apiClient.delete(`/admin/products/${id}`)
  }
}

// ==================== PROMOTIONS ====================
export const promotionsAPI = {
  getAll(params = {}) {
    return apiClient.get('/promotions', { params })
  },
  
  getById(id) {
    return apiClient.get(`/promotions/${id}`)
  },
  
  create(data) {
    return apiClient.post('/admin/promotions', data)
  },
  
  update(id, data) {
    return apiClient.put(`/admin/promotions/${id}`, data)
  },
  
  delete(id) {
    return apiClient.delete(`/admin/promotions/${id}`)
  }
}

// ==================== CART ====================
export const cartAPI = {
  get() {
    return apiClient.get('/cart')
  },
  
  addItem(data) {
    return apiClient.post('/cart/items', data)
  },
  
  updateItem(productId, data) {
    return apiClient.put(`/cart/items/${productId}`, data)
  },
  
  removeItem(productId) {
    return apiClient.delete(`/cart/items/${productId}`)
  },
  
  clear() {
    return apiClient.delete('/cart')
  }
}

// ==================== ORDERS ====================
export const ordersAPI = {
  getAll(params = {}) {
    return apiClient.get('/orders', { params })
  },
  
  getById(id) {
    return apiClient.get(`/orders/${id}`)
  },
  
  create(data) {
    return apiClient.post('/orders', data)
  },
  
  // Admin methods
  updateStatus(id, data) {
    return apiClient.put(`/admin/orders/${id}/status`, data)
  },
  
  getAllAdmin(params = {}) {
    return apiClient.get('/admin/orders', { params })
  }
}

// ==================== MOOD QUIZ ====================
export const moodAPI = {
  getQuestions() {
    return apiClient.get('/mood/questions')
  },
  
  getRecommendations(answers) {
    return apiClient.post('/mood/recommend', { answers })
  },
  
  // Admin methods
  getAllQuestions() {
    return apiClient.get('/admin/mood/questions')
  },
  
  getQuestionById(id) {
    return apiClient.get(`/admin/mood/questions/${id}`)
  },
  
  createQuestion(data) {
    return apiClient.post('/admin/mood/questions', data)
  },
  
  updateQuestion(id, data) {
    return apiClient.put(`/admin/mood/questions/${id}`, data)
  },
  
  deleteQuestion(id) {
    return apiClient.delete(`/admin/mood/questions/${id}`)
  }
}

// ==================== ANALYTICS (Admin) ====================
export const analyticsAPI = {
  getStats(params = {}) {
    return apiClient.get('/admin/analytics', { params })
  }
}

