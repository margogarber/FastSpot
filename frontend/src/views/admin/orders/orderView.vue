<template>
  <div class="order-view-page">
    <div class="page-header">
      <h1>📦 Order Details</h1>
      <router-link to="/admin/orders" class="btn btn-ghost">← Back</router-link>
    </div>

    <div v-if="loading" class="text-center py-2xl">
      <div class="spinner"></div>
    </div>

    <div v-else-if="order" class="order-details">
      <!-- Order Info -->
      <div class="details-card">
        <div class="card-header">
          <h2>Order Information</h2>
          <span :class="['badge', 'badge-lg', getStatusBadgeClass(order.status)]">
            {{ order.status }}
          </span>
        </div>
        <div class="details-grid">
          <div class="detail-item">
            <strong>Order Number:</strong>
            <span>{{ order.orderNumber }}</span>
          </div>
          <div class="detail-item">
            <strong>Date:</strong>
            <span>{{ formatDate(order.createdAt) }}</span>
          </div>
          <div class="detail-item">
            <strong>Total:</strong>
            <span class="text-primary"><strong>${{ order.totalUSD.toFixed(2) }}</strong></span>
          </div>
          <div class="detail-item">
            <strong>Payment Method:</strong>
            <span>{{ order.payment?.method }}</span>
          </div>
        </div>
      </div>

      <!-- Customer Info -->
      <div class="details-card">
        <div class="card-header">
          <h2>👤 Customer</h2>
        </div>
        <div class="details-grid">
          <div class="detail-item">
            <strong>Name:</strong>
            <span>{{ order.customerInfo?.name }}</span>
          </div>
          <div class="detail-item">
            <strong>Phone:</strong>
            <span>{{ order.customerInfo?.phone }}</span>
          </div>
          <div v-if="order.customerInfo?.email" class="detail-item">
            <strong>Email:</strong>
            <span>{{ order.customerInfo?.email }}</span>
          </div>
        </div>
      </div>

      <!-- Delivery Info -->
      <div class="details-card">
        <div class="card-header">
          <h2>🚚 Delivery</h2>
        </div>
        <div class="details-grid">
          <div class="detail-item">
            <strong>Type:</strong>
            <span>{{ order.delivery?.type }}</span>
          </div>
          <div v-if="order.delivery?.type === 'delivery'" class="detail-item full-width">
            <strong>Address:</strong>
            <p>
              {{ order.delivery?.address?.street }}, 
              {{ order.delivery?.address?.city }}, 
              {{ order.delivery?.address?.zipCode }}
            </p>
            <p v-if="order.delivery?.address?.notes">
              <em>Notes: {{ order.delivery?.address?.notes }}</em>
            </p>
          </div>
          <div v-if="order.delivery?.eta" class="detail-item">
            <strong>ETA:</strong>
            <span>{{ formatDate(order.delivery?.eta) }}</span>
          </div>
        </div>
      </div>

      <!-- Items -->
      <div class="details-card">
        <div class="card-header">
          <h2>🍔 Items</h2>
        </div>
        <div class="items-list">
          <div v-for="item in order.items" :key="item.productId" class="item-row">
            <img :src="item.image" :alt="item.name" class="item-image" />
            <div class="item-info">
              <strong>{{ item.name }}</strong>
              <div class="text-muted">Qty: {{ item.qty }}</div>
            </div>
            <div class="item-price">
              ${{ item.totalUSD.toFixed(2) }}
            </div>
          </div>
        </div>
      </div>

      <!-- Status Update -->
      <div class="details-card">
        <div class="card-header">
          <h2>🔄 Update Status</h2>
        </div>
        <div class="status-update-form">
          <div class="form-group">
            <label>New Status:</label>
            <select v-model="newStatus" class="form-input">
              <option value="new">New</option>
              <option value="confirmed">Confirmed</option>
              <option value="preparing">Preparing</option>
              <option value="ready">Ready</option>
              <option value="delivering">Delivering</option>
              <option value="completed">Completed</option>
              <option value="cancelled">Cancelled</option>
            </select>
          </div>
          <div class="form-group">
            <label>Note (optional):</label>
            <input v-model="statusNote" type="text" class="form-input" placeholder="Optional note" />
          </div>
          <button @click="updateStatus" class="btn btn-primary" :disabled="updating">
            {{ updating ? 'Updating...' : '✅ Update Status' }}
          </button>
        </div>
      </div>

      <!-- Tracking -->
      <div v-if="order.delivery?.tracking?.length" class="details-card">
        <div class="card-header">
          <h2>📍 Tracking</h2>
        </div>
        <div class="tracking-timeline">
          <div v-for="(event, index) in order.delivery.tracking" :key="index" class="timeline-event">
            <div class="timeline-dot"></div>
            <div class="timeline-content">
              <strong>{{ event.status }}</strong>
              <div class="text-muted">{{ formatDate(event.ts) }}</div>
              <p v-if="event.note">{{ event.note }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useOrdersStore } from '@/stores/orders'

const route = useRoute()
const ordersStore = useOrdersStore()
const loading = ref(true)
const updating = ref(false)
const order = ref(null)
const newStatus = ref('')
const statusNote = ref('')

onMounted(async () => {
  try {
    order.value = await ordersStore.fetchOrderById(route.params.id)
    newStatus.value = order.value.status
  } finally {
    loading.value = false
  }
})

async function updateStatus() {
  if (!newStatus.value) return
  
  try {
    updating.value = true
    order.value = await ordersStore.updateOrderStatus(
      route.params.id,
      newStatus.value,
      statusNote.value
    )
    alert('Status updated! ✅')
    statusNote.value = ''
  } catch (err) {
    alert('Failed to update status')
  } finally {
    updating.value = false
  }
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
.order-view-page {
  background-color: white;
  border-radius: var(--radius-lg);
  padding: var(--spacing-2xl);
  box-shadow: var(--shadow-sm);
  max-width: 1000px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-2xl);
  padding-bottom: var(--spacing-lg);
  border-bottom: 2px solid var(--border);
}

.page-header h1 {
  margin: 0;
}

.order-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xl);
}

.details-card {
  background-color: var(--bg-cream);
  padding: var(--spacing-lg);
  border-radius: var(--radius-md);
  border: 2px solid var(--border);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
  padding-bottom: var(--spacing-md);
  border-bottom: 2px solid var(--border);
}

.card-header h2 {
  margin: 0;
  color: var(--accent-red);
}

.details-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-md);
}

.detail-item {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.detail-item.full-width {
  grid-column: 1 / -1;
}

.detail-item strong {
  color: var(--text-light);
  font-size: var(--font-size-sm);
  text-transform: uppercase;
}

.detail-item span, .detail-item p {
  color: var(--text);
  font-size: var(--font-size-lg);
}

.items-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.item-row {
  display: flex;
  gap: var(--spacing-md);
  align-items: center;
  padding: var(--spacing-md);
  background-color: white;
  border-radius: var(--radius-sm);
  border: 2px solid var(--border);
}

.item-image {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: var(--radius-sm);
}

.item-info {
  flex: 1;
}

.item-price {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--accent-red);
}

.status-update-form {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.form-group label {
  font-weight: 600;
}

.form-input {
  padding: var(--spacing-sm);
  border: 2px solid var(--border);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-base);
}

.form-input:focus {
  outline: none;
  border-color: var(--accent-red);
}

.tracking-timeline {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.timeline-event {
  display: flex;
  gap: var(--spacing-md);
  position: relative;
}

.timeline-dot {
  width: 16px;
  height: 16px;
  background-color: var(--accent-red);
  border-radius: 50%;
  margin-top: 4px;
}

.timeline-content {
  flex: 1;
  padding-bottom: var(--spacing-md);
  border-left: 2px solid var(--border);
  padding-left: var(--spacing-md);
  margin-left: 7px;
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

  .details-grid {
    grid-template-columns: 1fr;
  }

  .item-row {
    flex-direction: column;
    text-align: center;
  }
}
</style>