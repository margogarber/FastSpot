<template>
  <div class="orders-list-page">
    <div class="page-header">
      <h1>📦 Orders Management</h1>
      <button @click="refreshOrders" class="btn btn-secondary">🔄 Refresh</button>
    </div>

    <div class="filters">
      <button 
        v-for="status in statuses" 
        :key="status.value"
        @click="filterStatus = status.value"
        :class="['filter-btn', filterStatus === status.value ? 'active' : '']"
      >
        {{ status.label }} ({{ getCountByStatus(status.value) }})
      </button>
    </div>

    <div v-if="loading" class="text-center py-md">
      <div class="spinner"></div>
    </div>

    <div v-else-if="filteredOrders.length" class="orders-table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th>Order #</th>
            <th>Customer</th>
            <th>Items</th>
            <th>Total</th>
            <th>Status</th>
            <th>Date</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="order in filteredOrders" :key="order.id">
            <td><strong>{{ order.orderNumber }}</strong></td>
            <td>
              <div>{{ order.customerInfo?.name }}</div>
              <div class="text-muted">{{ order.customerInfo?.phone }}</div>
            </td>
            <td>{{ order.items?.length || 0 }} items</td>
            <td class="text-primary"><strong>${{ order.totalUSD.toFixed(2) }}</strong></td>
            <td>
              <span :class="['badge', getStatusBadgeClass(order.status)]">
                {{ order.status }}
              </span>
            </td>
            <td>{{ formatDate(order.createdAt) }}</td>
            <td>
              <router-link 
                :to="`/admin/orders/${order.id}`"
                class="btn btn-sm btn-primary"
              >
                View
              </router-link>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-else class="text-center py-md">
      <p class="text-muted">No orders found</p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useOrdersStore } from '@/stores/orders'

const ordersStore = useOrdersStore()
const orders = computed(() => ordersStore.orders)
const loading = ref(false)
const filterStatus = ref('all')

const statuses = [
  { value: 'all', label: 'All' },
  { value: 'new', label: 'New' },
  { value: 'confirmed', label: 'Confirmed' },
  { value: 'preparing', label: 'Preparing' },
  { value: 'ready', label: 'Ready' },
  { value: 'delivering', label: 'Delivering' },
  { value: 'completed', label: 'Completed' },
  { value: 'cancelled', label: 'Cancelled' }
]

const filteredOrders = computed(() => {
  if (filterStatus.value === 'all') return orders.value
  return orders.value.filter(o => o.status === filterStatus.value)
})

onMounted(async () => {
  await refreshOrders()
})

async function refreshOrders() {
  loading.value = true
  await ordersStore.fetchAllOrdersAdmin()
  loading.value = false
}

function getCountByStatus(status) {
  if (status === 'all') return orders.value.length
  return orders.value.filter(o => o.status === status).length
}

function getStatusBadgeClass(status) {
  const classMap = {
    'new': 'badge-primary',
    'confirmed': 'badge-secondary',
    'preparing': 'badge-warning',
    'ready': 'badge-success',
    'delivering': 'badge-secondary',
    'completed': 'badge-success',
    'cancelled': 'badge-error'
  }
  return classMap[status] || 'badge-secondary'
}

function formatDate(dateString) {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleString()
}
</script>

<style scoped>
.orders-list-page {
  background-color: white;
  border-radius: var(--radius-lg);
  padding: var(--spacing-2xl);
  box-shadow: var(--shadow-sm);
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-xl);
}

.page-header h1 {
  margin: 0;
}

.filters {
  display: flex;
  gap: var(--spacing-xs);
  flex-wrap: wrap;
  margin-bottom: var(--spacing-xl);
  padding: var(--spacing-md);
  background-color: var(--bg-cream);
  border-radius: var(--radius-md);
}

.filter-btn {
  padding: var(--spacing-sm) var(--spacing-md);
  border: 2px solid var(--border);
  background-color: white;
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all var(--transition-base);
  font-weight: 500;
}

.filter-btn:hover {
  border-color: var(--accent-red);
}

.filter-btn.active {
  background-color: var(--accent-red);
  color: white;
  border-color: var(--accent-red);
}

.orders-table-container {
  overflow-x: auto;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table thead {
  background-color: var(--bg-cream);
}

.data-table th {
  padding: var(--spacing-md);
  text-align: left;
  font-weight: 600;
  border-bottom: 2px solid var(--border);
}

.data-table td {
  padding: var(--spacing-md);
  border-bottom: 1px solid var(--border);
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid var(--border);
  border-top-color: var(--accent-red);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: stretch;
  }
}
</style>