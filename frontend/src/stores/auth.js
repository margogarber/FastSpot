/**
 * Auth Store - Authentication management
 */
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authAPI } from '@/api'

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref(null)
  const token = ref(null)
  const sessionId = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Getters
  const isAuthenticated = computed(() => !!token.value && !!user.value)
  const isAdmin = computed(() => user.value?.role === 'admin')
  const isGuest = computed(() => !token.value || user.value?.role === 'guest')

  // Actions
  async function init() {
    // Restore from localStorage
    const savedToken = localStorage.getItem('auth_token')
    const savedUser = localStorage.getItem('user')
    const savedSessionId = localStorage.getItem('session_id')

    if (savedToken && savedUser) {
      token.value = savedToken
      user.value = JSON.parse(savedUser)
    } else if (savedSessionId) {
      sessionId.value = savedSessionId
    } else {
      // Create guest session
      await createGuestSession()
    }
  }

  async function createGuestSession() {
    try {
      loading.value = true
      error.value = null
      const { data } = await authAPI.createGuestSession()
      
      // Backend can return {sessionId, token} or error
      if (data.sessionId) {
        sessionId.value = data.sessionId
        token.value = data.token
        
        localStorage.setItem('session_id', sessionId.value)
        localStorage.setItem('auth_token', token.value)
      }
    } catch (err) {
      error.value = err.response?.data?.error || 'Failed to create session'
      console.error('Guest session error:', err)
      // Not critical if guest session failed to create
    } finally {
      loading.value = false
    }
  }

  async function login(credentials) {
    try {
      loading.value = true
      error.value = null
      const { data } = await authAPI.login(credentials)
      
      // Backend returns {success, token, user}
      user.value = data.user
      token.value = data.token
      
      localStorage.setItem('user', JSON.stringify(user.value))
      localStorage.setItem('auth_token', token.value)
      localStorage.removeItem('session_id')
      
      return true
    } catch (err) {
      error.value = err.response?.data?.error || 'Login failed'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function logout() {
    try {
      await authAPI.logout()
    } catch (err) {
      console.error('Logout error:', err)
    } finally {
      // Clear state
      user.value = null
      token.value = null
      sessionId.value = null
      
      localStorage.removeItem('user')
      localStorage.removeItem('auth_token')
      localStorage.removeItem('session_id')
      
      // Create new guest session
      await createGuestSession()
    }
  }

  async function fetchCurrentUser() {
    if (!token.value) return

    try {
      const { data } = await authAPI.me()
      user.value = data.user || data
      localStorage.setItem('user', JSON.stringify(user.value))
    } catch (err) {
      console.error('Fetch user error:', err)
      // If token is invalid, perform logout
      if (err.response?.status === 401) {
        await logout()
      }
    }
  }

  return {
    // State
    user,
    token,
    sessionId,
    loading,
    error,
    // Getters
    isAuthenticated,
    isAdmin,
    isGuest,
    // Actions
    init,
    createGuestSession,
    login,
    logout,
    fetchCurrentUser
  }
})

