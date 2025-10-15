<template>
  <div class="dashboard-page">
    <h1>📊 Dashboard</h1>
    
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-icon">🍔</div>
        <div class="stat-info">
          <div class="stat-value">{{ stats.productsCount || 0 }}</div>
          <div class="stat-label">Products</div>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">📦</div>
        <div class="stat-info">
          <div class="stat-value">{{ stats.ordersCount || 0 }}</div>
          <div class="stat-label">Orders</div>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">📂</div>
        <div class="stat-info">
          <div class="stat-value">{{ stats.categoriesCount || 0 }}</div>
          <div class="stat-label">Categories</div>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">🎉</div>
        <div class="stat-info">
          <div class="stat-value">{{ stats.promotionsCount || 0 }}</div>
          <div class="stat-label">Promotions</div>
        </div>
      </div>
    </div>

    <div class="quick-actions">
      <h2>Quick actions</h2>
      <div class="actions-grid">
        <router-link to="/admin/products/create" class="action-card">
          ➕ Add product
        </router-link>
        <router-link to="/admin/categories/create" class="action-card">
          📂 Create category
        </router-link>
        <router-link to="/admin/promotions/create" class="action-card">
          🎉 New promotion
        </router-link>
        <router-link to="/admin/orders" class="action-card">
          📦 View orders
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { analyticsAPI } from '@/api'

const stats = ref({
  productsCount: 0,
  ordersCount: 0,
  categoriesCount: 0,
  promotionsCount: 0
})

const loading = ref(false)

onMounted(async () => {
  await fetchStats()
})

async function fetchStats() {
  try {
    loading.value = true
    const { data } = await analyticsAPI.getStats()
    stats.value = {
      productsCount: data.data.productsCount || 0,
      ordersCount: data.data.ordersCount || 0,
      categoriesCount: data.data.categoriesCount || 0,
      promotionsCount: data.data.promotionsCount || 0
    }
  } catch (err) {
    console.error('Failed to fetch analytics:', err)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.dashboard-page h1 {
  margin-bottom: var(--spacing-2xl);
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-2xl);
}

.stat-card {
  display: flex;
  gap: var(--spacing-md);
  padding: var(--spacing-lg);
  background-color: white;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
}

.stat-icon {
  font-size: var(--font-size-4xl);
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: var(--font-size-3xl);
  font-weight: 700;
  color: var(--accent-red);
}

.stat-label {
  color: var(--text-light);
}

.quick-actions h2 {
  margin-bottom: var(--spacing-lg);
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-md);
}

.action-card {
  padding: var(--spacing-lg);
  background-color: white;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
  text-align: center;
  font-weight: 600;
  text-decoration: none;
  color: var(--text);
  transition: all var(--transition-base);
}

.action-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
  color: var(--accent-red);
}
</style>

