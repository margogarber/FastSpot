<template>
  <div class="order-track-page">
    <div class="container">
      <div v-if="loading" class="loading-state">
        <div class="spinner"></div>
        <p>Loading order...</p>
      </div>

      <div v-else-if="error" class="error-state">
        <p class="error-message">{{ error }}</p>
        <router-link to="/menu" class="btn btn-primary">
          Back to Menu
        </router-link>
      </div>

      <div v-else-if="order" class="order-content">
        <!-- Success Header -->
        <div class="success-header">
          <div class="success-icon">✅</div>
          <h1>Order Placed Successfully!</h1>
          <p class="order-number">Order #{{ order.orderNumber }}</p>
          <p class="thank-you">Thank you for your order, {{ order.customerInfo.name }}!</p>
        </div>

        <!-- Order Status -->
        <div class="status-card">
          <h2>📦 Order Status</h2>
          <div class="status-badge" :class="order.status">
            {{ formatStatus(order.status) }}
          </div>

          <!-- ETA -->
          <div v-if="order.delivery.eta" class="eta-info">
            <span class="eta-icon">⏰</span>
            <div>
              <strong>Estimated {{ order.delivery.type === 'pickup' ? 'Pickup' : 'Delivery' }} Time:</strong>
              <p>{{ formatTime(order.delivery.eta) }}</p>
            </div>
          </div>

          <!-- Delivery Type -->
          <div class="delivery-info">
            <span v-if="order.delivery.type === 'pickup'" class="delivery-icon">🏪</span>
            <span v-else class="delivery-icon">🚚</span>
            <div>
              <strong>{{ order.delivery.type === 'pickup' ? 'Pickup' : 'Delivery' }}</strong>
              <p v-if="order.delivery.type === 'delivery' && order.delivery.address">
                {{ order.delivery.address.street }}, {{ order.delivery.address.city }}, {{ order.delivery.address.zipCode }}
              </p>
            </div>
          </div>
        </div>

        <!-- Tracking -->
        <div class="tracking-card">
          <h2>📍 Order Tracking</h2>
          <div class="tracking-timeline">
            <div 
              v-for="(event, index) in order.delivery.tracking" 
              :key="index"
              class="tracking-event"
            >
              <div class="event-dot"></div>
              <div class="event-content">
                <strong>{{ event.status }}</strong>
                <p>{{ event.note }}</p>
                <span class="event-time">{{ formatDateTime(event.ts) }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Order Items -->
        <div class="items-card">
          <h2>🛍️ Order Items</h2>
          <div class="order-items">
            <div v-for="item in order.items" :key="item.productId" class="order-item">
              <img :src="item.image" :alt="item.name" />
              <div class="item-info">
                <h3>{{ item.name }}</h3>
                <p>Qty: {{ item.qty }} × ${{ item.unitPriceUSD.toFixed(2) }}</p>
              </div>
              <div class="item-total">
                ${{ item.totalUSD.toFixed(2) }}
              </div>
            </div>
          </div>

          <div class="order-total">
            <span>Total:</span>
            <span class="total-amount">${{ order.totalUSD.toFixed(2) }}</span>
          </div>
        </div>

        <!-- Payment Info -->
        <div class="payment-card">
          <h2>💳 Payment Information</h2>
          <div class="payment-details">
            <div class="payment-row">
              <span>Method:</span>
              <span class="payment-method">{{ formatPaymentMethod(order.payment.method) }}</span>
            </div>
            <div class="payment-row">
              <span>Status:</span>
              <span class="payment-status" :class="order.payment.status">
                {{ order.payment.status }}
              </span>
            </div>
            <div v-if="order.payment.txnId" class="payment-row">
              <span>Transaction ID:</span>
              <span class="txn-id">{{ order.payment.txnId }}</span>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="order-actions">
          <router-link to="/menu" class="btn btn-primary">
            🍔 Order More
          </router-link>
          <button @click="refreshOrder" class="btn btn-secondary">
            🔄 Refresh Status
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ordersAPI } from '@/api'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const error = ref(null)
const order = ref(null)

onMounted(async () => {
  await fetchOrder()
})

async function fetchOrder() {
  try {
    loading.value = true
    error.value = null
    
    const orderId = route.params.id
    const { data } = await ordersAPI.getById(orderId)
    
    order.value = data.data
  } catch (err) {
    error.value = err.response?.data?.error || 'Failed to load order'
    console.error('Fetch order error:', err)
  } finally {
    loading.value = false
  }
}

async function refreshOrder() {
  await fetchOrder()
}

function formatStatus(status) {
  const statusMap = {
    'new': '🆕 New',
    'confirmed': '✅ Confirmed',
    'preparing': '👨‍🍳 Preparing',
    'ready': '✨ Ready',
    'delivering': '🚚 On the way',
    'completed': '🎉 Completed',
    'cancelled': '❌ Cancelled'
  }
  return statusMap[status] || status
}

function formatPaymentMethod(method) {
  const methodMap = {
    'card': '💳 Credit/Debit Card',
    'applepay': '🍎 Apple Pay',
    'googlepay': '🔵 Google Pay',
    'cash': '💵 Cash'
  }
  return methodMap[method] || method
}

function formatTime(timestamp) {
  const date = new Date(timestamp)
  return date.toLocaleTimeString('en-US', { 
    hour: '2-digit', 
    minute: '2-digit',
    hour12: true
  })
}

function formatDateTime(timestamp) {
  const date = new Date(timestamp)
  return date.toLocaleString('en-US', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    hour12: true
  })
}
</script>

<style scoped>
.order-track-page {
  min-height: 80vh;
  padding: var(--spacing-xl) 0;
  background-color: var(--bg-cream);
}

.loading-state,
.error-state {
  text-align: center;
  padding: var(--spacing-3xl);
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid var(--bg-cream);
  border-top-color: var(--accent-red);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto var(--spacing-lg);
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.order-content {
  max-width: 800px;
  margin: 0 auto;
}

.success-header {
  text-align: center;
  padding: var(--spacing-2xl);
  background: white;
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-lg);
  margin-bottom: var(--spacing-xl);
}

.success-icon {
  font-size: 64px;
  margin-bottom: var(--spacing-md);
}

.success-header h1 {
  color: var(--accent-red);
  margin-bottom: var(--spacing-sm);
}

.order-number {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--contrast-black);
  margin-bottom: var(--spacing-sm);
}

.thank-you {
  color: var(--text-light);
}

.status-card,
.tracking-card,
.items-card,
.payment-card {
  background: white;
  padding: var(--spacing-xl);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-md);
  margin-bottom: var(--spacing-xl);
}

.status-card h2,
.tracking-card h2,
.items-card h2,
.payment-card h2 {
  margin-bottom: var(--spacing-lg);
  color: var(--contrast-black);
}

.status-badge {
  display: inline-block;
  padding: var(--spacing-sm) var(--spacing-lg);
  border-radius: var(--radius-full);
  font-weight: 700;
  font-size: var(--font-size-lg);
  margin-bottom: var(--spacing-lg);
}

.status-badge.new,
.status-badge.confirmed {
  background-color: #dbeafe;
  color: #1e40af;
}

.status-badge.preparing {
  background-color: #fef3c7;
  color: #92400e;
}

.status-badge.ready,
.status-badge.completed {
  background-color: #d1fae5;
  color: #065f46;
}

.eta-info,
.delivery-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  background-color: var(--bg-cream);
  border-radius: var(--radius-md);
  margin-bottom: var(--spacing-md);
}

.eta-icon,
.delivery-icon {
  font-size: var(--font-size-2xl);
}

.tracking-timeline {
  position: relative;
  padding-left: var(--spacing-xl);
}

.tracking-timeline::before {
  content: '';
  position: absolute;
  left: 8px;
  top: 0;
  bottom: 0;
  width: 2px;
  background-color: var(--accent-yellow);
}

.tracking-event {
  position: relative;
  padding-bottom: var(--spacing-lg);
}

.event-dot {
  position: absolute;
  left: -29px;
  top: 4px;
  width: 16px;
  height: 16px;
  background-color: var(--accent-red);
  border-radius: 50%;
  border: 3px solid white;
  box-shadow: 0 0 0 2px var(--accent-red);
}

.event-content strong {
  display: block;
  color: var(--contrast-black);
  margin-bottom: var(--spacing-xs);
}

.event-content p {
  color: var(--text-light);
  margin-bottom: var(--spacing-xs);
}

.event-time {
  font-size: var(--font-size-sm);
  color: var(--text-light);
}

.order-items {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-lg);
}

.order-item {
  display: flex;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  border-bottom: 1px solid var(--bg-cream);
}

.order-item img {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: var(--radius-md);
}

.item-info {
  flex: 1;
}

.item-info h3 {
  margin-bottom: var(--spacing-xs);
}

.item-info p {
  color: var(--text-light);
}

.item-total {
  font-weight: 700;
  color: var(--accent-red);
  font-size: var(--font-size-lg);
}

.order-total {
  display: flex;
  justify-content: space-between;
  padding-top: var(--spacing-lg);
  border-top: 2px solid var(--neutral-warm);
  font-size: var(--font-size-xl);
  font-weight: 700;
}

.total-amount {
  color: var(--accent-red);
}

.payment-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.payment-row {
  display: flex;
  justify-content: space-between;
  padding: var(--spacing-sm) 0;
  border-bottom: 1px solid var(--bg-cream);
}

.payment-status.pending {
  color: #92400e;
}

.payment-status.completed {
  color: #065f46;
  font-weight: 700;
}

.txn-id {
  font-family: monospace;
  font-size: var(--font-size-sm);
  color: var(--text-light);
}

.order-actions {
  display: flex;
  gap: var(--spacing-md);
  margin-top: var(--spacing-xl);
}

.order-actions .btn {
  flex: 1;
}

@media (max-width: 768px) {
  .order-actions {
    flex-direction: column;
  }
}
</style>
