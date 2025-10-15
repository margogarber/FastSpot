<template>
  <div class="cart-page">
    <div class="container">
      <h1>🛒 Cart</h1>

      <div v-if="loading" class="text-center py-md">
        <div class="spinner"></div>
      </div>

      <div v-else-if="cartStore.isEmpty" class="empty-cart">
        <p>Your cart is empty</p>
        <router-link to="/menu" class="btn btn-primary">
          Go to menu
        </router-link>
      </div>

      <div v-else class="cart-content">
        <div class="cart-items">
          <div v-for="item in cartStore.items" :key="item.productId" class="cart-item">
            <img :src="item.image" :alt="item.name" class="item-image" />
            
            <div class="item-info">
              <h3>{{ item.name }}</h3>
              <p class="price-per-unit">${{ item.unitPriceUSD.toFixed(2) }} each</p>
              
              <!-- Chosen Options -->
              <div v-if="Object.keys(item.chosenOptions).length" class="item-options">
                <span v-for="(value, key) in item.chosenOptions" :key="key" class="badge badge-secondary">
                  {{ key }}: {{ value }}
                </span>
              </div>

              <!-- Ingredients -->
              <div v-if="item.chosenIngredients.length" class="item-ingredients">
                <small>Ingredients: {{ item.chosenIngredients.join(', ') }}</small>
              </div>
            </div>

            <div class="item-actions">
              <div class="quantity-control">
                <button 
                  @click="decreaseQty(item)" 
                  class="btn btn-sm btn-ghost"
                  :disabled="item.qty <= 1"
                >
                  -
                </button>
                <span class="quantity">{{ item.qty }}</span>
                <button 
                  @click="increaseQty(item)" 
                  class="btn btn-sm btn-ghost"
                >
                  +
                </button>
              </div>
              
              <div class="item-price">${{ item.totalUSD.toFixed(2) }}</div>
              
              <button 
                @click="removeItem(item)" 
                class="btn btn-sm btn-outline"
              >
                🗑️ Delete
              </button>
            </div>
          </div>
        </div>

        <div class="cart-summary">
          <h3>Total</h3>
          <div class="summary-row">
            <span>Items:</span>
            <span>{{ cartStore.itemsCount }}</span>
          </div>
          <div class="summary-row total">
            <span>Sum:</span>
            <span>${{ cartStore.totalUSD.toFixed(2) }}</span>
          </div>
          
          <router-link to="/checkout" class="btn btn-primary btn-lg">
            Checkout
          </router-link>
          
          <button @click="clearCart" class="btn btn-ghost">
            Clear cart
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useCartStore } from '@/stores/cart'

const cartStore = useCartStore()
const loading = ref(false)

onMounted(async () => {
  loading.value = true
  await cartStore.fetchCart()
  loading.value = false
})

async function increaseQty(item) {
  await cartStore.updateItem(
    item.productId, 
    item.qty + 1,
    item.chosenIngredients,
    item.chosenOptions
  )
}

async function decreaseQty(item) {
  if (item.qty > 1) {
    await cartStore.updateItem(
      item.productId, 
      item.qty - 1,
      item.chosenIngredients,
      item.chosenOptions
    )
  }
}

async function removeItem(item) {
  if (confirm('Delete item from cart?')) {
    await cartStore.removeItem(item.productId)
  }
}

async function clearCart() {
  if (confirm('Clear all cart?')) {
    await cartStore.clearCart()
  }
}
</script>

<style scoped>
.cart-page {
  min-height: 80vh;
  padding: var(--spacing-xl) 0;
}

.empty-cart {
  text-align: center;
  padding: var(--spacing-2xl);
}

.cart-content {
  display: grid;
  grid-template-columns: 1fr 350px;
  gap: var(--spacing-2xl);
}

.cart-items {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.cart-item {
  display: flex;
  gap: var(--spacing-lg);
  padding: var(--spacing-lg);
  background-color: white;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
}

.item-image {
  width: 120px;
  height: 120px;
  object-fit: cover;
  border-radius: var(--radius-md);
}

.item-info {
  flex: 1;
}

.item-info h3 {
  margin-bottom: var(--spacing-sm);
}

.price-per-unit {
  color: var(--text-light);
  font-size: var(--font-size-sm);
  margin-bottom: var(--spacing-xs);
}

.item-options,
.item-ingredients {
  margin-top: var(--spacing-sm);
  display: flex;
  flex-wrap: wrap;
  gap: var(--spacing-xs);
}

.item-actions {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: var(--spacing-md);
}

.quantity-control {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  background-color: var(--bg-cream);
  padding: var(--spacing-xs);
  border-radius: var(--radius-md);
}

.quantity {
  min-width: 30px;
  text-align: center;
  font-weight: 600;
}

.item-price {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--accent-red);
}

.cart-summary {
  position: sticky;
  top: 80px;
  height: fit-content;
  padding: var(--spacing-lg);
  background-color: white;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-md);
}

.cart-summary h3 {
  margin-bottom: var(--spacing-lg);
}

.summary-row {
  display: flex;
  justify-content: space-between;
  padding: var(--spacing-sm) 0;
  border-bottom: 1px solid var(--border);
}

.summary-row.total {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--accent-red);
  border-bottom: none;
  margin-top: var(--spacing-md);
  margin-bottom: var(--spacing-lg);
}

.cart-summary .btn {
  width: 100%;
  margin-bottom: var(--spacing-sm);
}

@media (max-width: 968px) {
  .cart-content {
    grid-template-columns: 1fr;
  }
  
  .cart-item {
    flex-direction: column;
  }
  
  .item-actions {
    flex-direction: row;
    justify-content: space-between;
  }
}
</style>

