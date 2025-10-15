<template>
  <nav class="navbar">
    <div class="container">
      <div class="navbar-content">
        <!-- Logo -->
        <router-link to="/" class="logo">
          🍔 <span>FastSpot</span>
        </router-link>

        <!-- Navigation Links -->
        <div class="nav-links">
          <router-link to="/" class="nav-link">Home</router-link>
          <router-link to="/menu" class="nav-link">Menu</router-link>
          <router-link to="/promotions" class="nav-link">Promotions</router-link>
          <router-link to="/mood" class="nav-link mood-link">
            😊 Mood Quiz
          </router-link>
        </div>

        <!-- Cart & Auth -->
        <div class="nav-actions">
          <router-link to="/cart" class="cart-button">
            🛒
            <span v-if="cartStore.itemsCount > 0" class="cart-badge">
              {{ cartStore.itemsCount }}
            </span>
          </router-link>

          <router-link 
            v-if="authStore.isAdmin" 
            to="/admin" 
            class="btn btn-sm btn-primary"
          >
            Admin panel
          </router-link>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useCartStore } from '@/stores/cart'

const authStore = useAuthStore()
const cartStore = useCartStore()

onMounted(async () => {
  // Load cart for all users (guests and authenticated)
  if (!window.location.pathname.startsWith('/admin')) {
    await cartStore.fetchCart()
  }
})
</script>

<style scoped>
.navbar {
  background-color: white;
  box-shadow: var(--shadow-md);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.navbar-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--spacing-md) 0;
  gap: var(--spacing-xl);
}

.logo {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  font-size: var(--font-size-2xl);
  font-weight: 700;
  color: var(--accent-red);
  text-decoration: none;
  transition: transform var(--transition-fast);
}

.logo:hover {
  transform: scale(1.05);
}

.nav-links {
  display: flex;
  align-items: center;
  gap: var(--spacing-lg);
  flex: 1;
}

.nav-link {
  font-size: var(--font-size-base);
  font-weight: 600;
  color: var(--text);
  text-decoration: none;
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-md);
  transition: all var(--transition-fast);
}

.nav-link:hover {
  background-color: var(--bg-cream);
  color: var(--accent-red);
}

.nav-link.router-link-active {
  color: var(--accent-red);
  background-color: rgba(230, 57, 70, 0.1);
}

.mood-link {
  background-color: var(--accent-yellow);
  color: var(--text-brown);
}

.mood-link:hover {
  background-color: #FFC451;
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.cart-button {
  position: relative;
  font-size: var(--font-size-2xl);
  padding: var(--spacing-sm);
  border-radius: var(--radius-md);
  transition: all var(--transition-fast);
  text-decoration: none;
}

.cart-button:hover {
  background-color: var(--bg-cream);
  transform: scale(1.1);
}

.cart-badge {
  position: absolute;
  top: 0;
  right: 0;
  background-color: var(--accent-red);
  color: white;
  font-size: var(--font-size-xs);
  font-weight: 700;
  padding: 2px 6px;
  border-radius: var(--radius-full);
  min-width: 20px;
  text-align: center;
}

@media (max-width: 768px) {
  .navbar-content {
    flex-wrap: wrap;
  }

  .nav-links {
    width: 100%;
    justify-content: space-around;
    order: 3;
  }

  .nav-link {
    font-size: var(--font-size-sm);
    padding: var(--spacing-xs) var(--spacing-sm);
  }
}
</style>

