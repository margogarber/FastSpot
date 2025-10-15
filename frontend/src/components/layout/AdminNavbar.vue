<template>
  <nav class="admin-navbar">
    <div class="navbar-content">
      <div class="navbar-left">
        <h2>🍔 FastSpot Admin</h2>
      </div>

      <div class="navbar-right">
        <span class="user-info">
          👤 {{ authStore.user?.name || 'Admin' }}
        </span>
        <button @click="handleLogout" class="btn btn-sm btn-outline">
          Logout
        </button>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = async () => {
  await authStore.logout()
  router.push('/')
}
</script>

<style scoped>
.admin-navbar {
  background-color: white;
  box-shadow: var(--shadow-sm);
  position: sticky;
  top: 0;
  z-index: 100;
}

.navbar-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--spacing-md) var(--spacing-xl);
}

.navbar-left h2 {
  margin: 0;
  color: var(--accent-red);
  font-size: var(--font-size-xl);
}

.navbar-right {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.user-info {
  font-weight: 600;
  color: var(--text);
}
</style>

