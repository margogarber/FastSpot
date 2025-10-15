<template>
  <router-link 
    :to="`/products/${product.slug}`" 
    class="product-card"
  >
    <div class="product-image">
      <img :src="product.image" :alt="product.name" />
      <span v-if="product.tags?.includes('popular')" class="badge badge-primary popular-badge">
        🔥 POPULAR
      </span>
    </div>

    <div class="product-info">
      <h3 class="product-name">{{ product.name }}</h3>
      <p class="product-description">{{ truncateText(product.description, 80) }}</p>
      
      <div class="product-footer">
        <span class="product-price">${{ product.priceUSD.toFixed(2) }}</span>
        <button @click.prevent="handleAddToCart" class="btn btn-sm btn-primary">
          Add to Cart
        </button>
      </div>
    </div>
  </router-link>
</template>

<script setup>
import { useCartStore } from '@/stores/cart'
import { useRouter } from 'vue-router'

const props = defineProps({
  product: {
    type: Object,
    required: true
  }
})

const cartStore = useCartStore()
const router = useRouter()

function truncateText(text, maxLength) {
  if (!text) return ''
  return text.length > maxLength ? text.substring(0, maxLength) + '...' : text
}

async function handleAddToCart() {
  try {
    await cartStore.addItem(props.product.id, 1)
    // You can add a toast notification here
    alert('Product added to cart!')
  } catch (err) {
    alert('Failed to add to cart')
  }
}
</script>

<style scoped>
.product-card {
  display: block;
  background-color: white;
  border-radius: var(--radius-lg);
  overflow: hidden;
  box-shadow: var(--shadow-md);
  transition: all var(--transition-base);
  text-decoration: none;
  color: inherit;
}

.product-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-lg);
}

.product-image {
  position: relative;
  width: 100%;
  height: 200px;
  overflow: hidden;
}

.product-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform var(--transition-base);
}

.product-card:hover .product-image img {
  transform: scale(1.05);
}

.popular-badge {
  position: absolute;
  top: var(--spacing-sm);
  right: var(--spacing-sm);
}

.product-info {
  padding: var(--spacing-lg);
}

.product-name {
  font-size: var(--font-size-lg);
  font-weight: 700;
  margin-bottom: var(--spacing-sm);
  color: var(--contrast-black);
}

.product-description {
  color: var(--text-light);
  font-size: var(--font-size-sm);
  margin-bottom: var(--spacing-md);
  line-height: 1.5;
}

.product-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.product-price {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--accent-red);
}
</style>

