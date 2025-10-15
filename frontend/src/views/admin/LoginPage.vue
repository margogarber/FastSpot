<template>
  <div class="login-page">
    <div class="login-card">
      <h2>🔐 Admin login</h2>
      
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label class="form-label">Email</label>
          <input 
            v-model="form.email" 
            type="email" 
            class="form-input"
            placeholder="admin@local"
            required
          />
        </div>

        <div class="form-group">
          <label class="form-label">Password</label>
          <input 
            v-model="form.password" 
            type="password" 
            class="form-input"
            placeholder="••••••••"
            required
          />
        </div>

        <div v-if="error" class="alert alert-error">
          {{ error }}
        </div>

        <button 
          type="submit" 
          class="btn btn-primary btn-lg"
          :disabled="loading"
        >
          <span v-if="loading" class="spinner"></span>
          <span v-else>Login</span>
        </button>
      </form>

      <div class="login-footer">
        <router-link to="/">← Back to site</router-link>
      </div>

      <div class="demo-credentials">
        <small class="text-muted">
          <strong>Demo login:</strong><br />
          Email: admin@local<br />
          Password: Admin123!
        </small>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const form = ref({
  email: '',
  password: ''
})

const loading = ref(false)
const error = ref('')

async function handleLogin() {
  try {
    loading.value = true
    error.value = ''
    
    await authStore.login(form.value)
    
    // Redirect to admin dashboard
    router.push('/admin')
  } catch (err) {
    console.error('Login error:', err)
    error.value = err.response?.data?.error || err.message || 'Login error'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
}

.login-card {
  width: 100%;
  max-width: 400px;
}

.login-card h2 {
  text-align: center;
  margin-bottom: var(--spacing-2xl);
  color: var(--contrast-black);
}

.login-card .btn {
  width: 100%;
  margin-top: var(--spacing-md);
}

.login-footer {
  text-align: center;
  margin-top: var(--spacing-lg);
}

.demo-credentials {
  margin-top: var(--spacing-xl);
  padding: var(--spacing-md);
  background-color: var(--bg-cream);
  border-radius: var(--radius-md);
  text-align: center;
}
</style>

