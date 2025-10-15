<template>
  <div class="checkout-page">
    <div class="container">
      <h1>🛍️ Checkout</h1>

      <div v-if="cartStore.isEmpty" class="empty-state">
        <p>Your cart is empty</p>
        <router-link to="/menu" class="btn btn-primary">
          Go to Menu
        </router-link>
      </div>

      <div v-else class="checkout-content">
        <!-- Order Summary -->
        <div class="order-summary-section">
          <h2>Order Summary</h2>
          
          <div class="summary-items">
            <div v-for="item in cartStore.items" :key="item.productId" class="summary-item">
              <img :src="item.image" :alt="item.name" />
              <div class="item-details">
                <h4>{{ item.name }}</h4>
                <p>Qty: {{ item.qty }} × ${{ item.unitPriceUSD.toFixed(2) }}</p>
              </div>
              <div class="item-total">
                ${{ item.totalUSD.toFixed(2) }}
              </div>
            </div>
          </div>

          <div class="summary-total">
            <div class="total-row">
              <span>Subtotal:</span>
              <span>${{ cartStore.totalUSD.toFixed(2) }}</span>
            </div>
            <div class="total-row highlight">
              <span>Total:</span>
              <span>${{ cartStore.totalUSD.toFixed(2) }}</span>
            </div>
          </div>
        </div>

        <!-- Checkout Form -->
        <div class="checkout-form-section">
          <form @submit.prevent="submitOrder" class="checkout-form">
            <!-- Contact Information -->
            <div class="form-section">
              <h3>📞 Contact Information</h3>
              
              <div class="form-group">
                <label for="name">Full Name *</label>
                <input 
                  type="text" 
                  id="name" 
                  v-model="formData.customerInfo.name"
                  required
                  placeholder="John Doe"
                />
              </div>

              <div class="form-group">
                <label for="phone">Phone Number *</label>
                <input 
                  type="tel" 
                  id="phone" 
                  v-model="formData.customerInfo.phone"
                  required
                  placeholder="+1 (555) 123-4567"
                />
              </div>

              <div class="form-group">
                <label for="email">Email (optional)</label>
                <input 
                  type="email" 
                  id="email" 
                  v-model="formData.customerInfo.email"
                  placeholder="john@example.com"
                />
              </div>
            </div>

            <!-- Delivery Type -->
            <div class="form-section">
              <h3>🚗 Delivery Type</h3>
              
              <div class="radio-group">
                <label class="radio-card" :class="{ active: formData.deliveryType === 'pickup' }">
                  <input 
                    type="radio" 
                    value="pickup" 
                    v-model="formData.deliveryType"
                  />
                  <div class="radio-content">
                    <span class="radio-icon">🏪</span>
                    <div>
                      <strong>Pickup</strong>
                      <p>Ready in 15-20 minutes</p>
                    </div>
                  </div>
                </label>

                <label class="radio-card" :class="{ active: formData.deliveryType === 'delivery' }">
                  <input 
                    type="radio" 
                    value="delivery" 
                    v-model="formData.deliveryType"
                  />
                  <div class="radio-content">
                    <span class="radio-icon">🚚</span>
                    <div>
                      <strong>Delivery</strong>
                      <p>Delivered in 30-45 minutes</p>
                    </div>
                  </div>
                </label>
              </div>
            </div>

            <!-- Delivery Address (if delivery selected) -->
            <div v-if="formData.deliveryType === 'delivery'" class="form-section">
              <h3>📍 Delivery Address</h3>
              
              <div class="form-group">
                <label for="street">Street Address *</label>
                <input 
                  type="text" 
                  id="street" 
                  v-model="formData.deliveryAddress.street"
                  :required="formData.deliveryType === 'delivery'"
                  placeholder="123 Main St, Apt 4B"
                />
              </div>

              <div class="form-row">
                <div class="form-group">
                  <label for="city">City *</label>
                  <input 
                    type="text" 
                    id="city" 
                    v-model="formData.deliveryAddress.city"
                    :required="formData.deliveryType === 'delivery'"
                    placeholder="New York"
                  />
                </div>

                <div class="form-group">
                  <label for="zipCode">ZIP Code *</label>
                  <input 
                    type="text" 
                    id="zipCode" 
                    v-model="formData.deliveryAddress.zipCode"
                    :required="formData.deliveryType === 'delivery'"
                    placeholder="10001"
                  />
                </div>
              </div>

              <div class="form-group">
                <label for="notes">Delivery Notes (optional)</label>
                <textarea 
                  id="notes" 
                  v-model="formData.deliveryAddress.notes"
                  rows="3"
                  placeholder="Any special instructions for delivery..."
                ></textarea>
              </div>
            </div>

            <!-- Payment Method -->
            <div class="form-section">
              <h3>💳 Payment Method</h3>
              
              <div class="radio-group">
                <label class="radio-card" :class="{ active: formData.paymentMethod === 'card' }">
                  <input 
                    type="radio" 
                    value="card" 
                    v-model="formData.paymentMethod"
                  />
                  <div class="radio-content">
                    <span class="radio-icon">💳</span>
                    <strong>Credit/Debit Card</strong>
                  </div>
                </label>

                <label class="radio-card" :class="{ active: formData.paymentMethod === 'applepay' }">
                  <input 
                    type="radio" 
                    value="applepay" 
                    v-model="formData.paymentMethod"
                  />
                  <div class="radio-content">
                    <span class="radio-icon">🍎</span>
                    <strong>Apple Pay</strong>
                  </div>
                </label>

                <label class="radio-card" :class="{ active: formData.paymentMethod === 'googlepay' }">
                  <input 
                    type="radio" 
                    value="googlepay" 
                    v-model="formData.paymentMethod"
                  />
                  <div class="radio-content">
                    <span class="radio-icon">🔵</span>
                    <strong>Google Pay</strong>
                  </div>
                </label>

                <label class="radio-card" :class="{ active: formData.paymentMethod === 'cash' }">
                  <input 
                    type="radio" 
                    value="cash" 
                    v-model="formData.paymentMethod"
                  />
                  <div class="radio-content">
                    <span class="radio-icon">💵</span>
                    <strong>Cash on {{ formData.deliveryType === 'pickup' ? 'Pickup' : 'Delivery' }}</strong>
                  </div>
                </label>
              </div>
            </div>

            <!-- Submit Button -->
            <div class="form-actions">
              <button type="submit" class="btn btn-primary btn-lg" :disabled="submitting">
                {{ submitting ? '⏳ Processing...' : '✅ Place Order' }}
              </button>
              <router-link to="/cart" class="btn btn-secondary">
                ← Back to Cart
              </router-link>
            </div>

            <div v-if="error" class="error-message">
              {{ error }}
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import { ordersAPI } from '@/api'

const router = useRouter()
const cartStore = useCartStore()

const submitting = ref(false)
const error = ref(null)

const formData = reactive({
  customerInfo: {
    name: '',
    email: '',
    phone: ''
  },
  deliveryType: 'pickup',
  deliveryAddress: {
    street: '',
    city: '',
    zipCode: '',
    notes: ''
  },
  paymentMethod: 'card'
})

onMounted(async () => {
  await cartStore.fetchCart()
  
  if (cartStore.isEmpty) {
    router.push('/menu')
  }
})

async function submitOrder() {
  try {
    submitting.value = true
    error.value = null

    const orderData = {
      paymentMethod: formData.paymentMethod,
      deliveryType: formData.deliveryType,
      customerInfo: {
        name: formData.customerInfo.name,
        email: formData.customerInfo.email || '',
        phone: formData.customerInfo.phone
      }
    }

    // Add delivery address if delivery type is selected
    if (formData.deliveryType === 'delivery') {
      orderData.deliveryAddress = {
        street: formData.deliveryAddress.street,
        city: formData.deliveryAddress.city,
        zipCode: formData.deliveryAddress.zipCode,
        notes: formData.deliveryAddress.notes || ''
      }
    }

    const { data } = await ordersAPI.create(orderData)
    
    // Redirect to order confirmation page
    const orderId = data.data.order.id
    router.push(`/orders/${orderId}`)
    
  } catch (err) {
    error.value = err.response?.data?.error || 'Failed to place order. Please try again.'
    console.error('Order submission error:', err)
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.checkout-page {
  min-height: 80vh;
  padding: var(--spacing-xl) 0;
  background-color: var(--bg-cream);
}

.checkout-page h1 {
  text-align: center;
  margin-bottom: var(--spacing-2xl);
  color: var(--contrast-black);
}

.empty-state {
  text-align: center;
  padding: var(--spacing-3xl);
}

.checkout-content {
  display: grid;
  grid-template-columns: 1fr 450px;
  gap: var(--spacing-2xl);
  align-items: start;
}

/* Order Summary */
.order-summary-section {
  background: white;
  padding: var(--spacing-xl);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-md);
  position: sticky;
  top: 100px;
}

.order-summary-section h2 {
  margin-bottom: var(--spacing-lg);
  color: var(--contrast-black);
}

.summary-items {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-lg);
  max-height: 400px;
  overflow-y: auto;
}

.summary-item {
  display: flex;
  gap: var(--spacing-md);
  padding: var(--spacing-sm);
  border-bottom: 1px solid var(--bg-cream);
}

.summary-item img {
  width: 60px;
  height: 60px;
  object-fit: cover;
  border-radius: var(--radius-md);
}

.item-details {
  flex: 1;
}

.item-details h4 {
  font-size: var(--font-size-md);
  margin-bottom: var(--spacing-xs);
}

.item-details p {
  font-size: var(--font-size-sm);
  color: var(--text-light);
}

.item-total {
  font-weight: 700;
  color: var(--accent-red);
}

.summary-total {
  border-top: 2px solid var(--neutral-warm);
  padding-top: var(--spacing-md);
}

.total-row {
  display: flex;
  justify-content: space-between;
  padding: var(--spacing-sm) 0;
}

.total-row.highlight {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--accent-red);
  margin-top: var(--spacing-sm);
}

/* Checkout Form */
.checkout-form-section {
  background: white;
  padding: var(--spacing-xl);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-md);
}

.form-section {
  margin-bottom: var(--spacing-2xl);
  padding-bottom: var(--spacing-xl);
  border-bottom: 1px solid var(--bg-cream);
}

.form-section:last-of-type {
  border-bottom: none;
}

.form-section h3 {
  margin-bottom: var(--spacing-lg);
  color: var(--contrast-black);
}

.form-group {
  margin-bottom: var(--spacing-lg);
}

.form-group label {
  display: block;
  margin-bottom: var(--spacing-sm);
  font-weight: 600;
  color: var(--text-brown);
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: var(--spacing-md);
  border: 2px solid var(--neutral-warm);
  border-radius: var(--radius-md);
  font-size: var(--font-size-md);
  transition: all var(--transition-base);
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: var(--accent-red);
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-md);
}

.radio-group {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-md);
}

.radio-card {
  position: relative;
  padding: var(--spacing-lg);
  border: 2px solid var(--neutral-warm);
  border-radius: var(--radius-lg);
  cursor: pointer;
  transition: all var(--transition-base);
}

.radio-card:hover {
  border-color: var(--accent-red);
  box-shadow: var(--shadow-sm);
}

.radio-card.active {
  border-color: var(--accent-red);
  background-color: var(--bg-cream);
}

.radio-card input[type="radio"] {
  position: absolute;
  opacity: 0;
}

.radio-content {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.radio-icon {
  font-size: var(--font-size-2xl);
}

.form-actions {
  display: flex;
  gap: var(--spacing-md);
  margin-top: var(--spacing-xl);
}

.form-actions .btn {
  flex: 1;
}

.error-message {
  margin-top: var(--spacing-lg);
  padding: var(--spacing-md);
  background-color: #fee;
  color: #c00;
  border-radius: var(--radius-md);
  text-align: center;
}

@media (max-width: 1024px) {
  .checkout-content {
    grid-template-columns: 1fr;
  }

  .order-summary-section {
    position: static;
  }

  .radio-group {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }

  .form-actions {
    flex-direction: column-reverse;
  }
}
</style>
