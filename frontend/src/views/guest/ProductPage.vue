<template>
  <div class="product-page">
    <div class="container">
      <div v-if="loading" class="text-center py-xl">
        <div class="spinner"></div>
      </div>

      <div v-else-if="product" class="product-details">
        <!-- Product Info -->
        <div class="product-image-section">
          <img :src="product.image" :alt="product.name" class="product-image" />
        </div>

        <div class="product-info-section">
          <h1>{{ product.name }}</h1>
          <p class="product-description">{{ product.description }}</p>
          
          <div class="product-price">
            ${{ calculateTotalPrice().toFixed(2) }}
          </div>

          <!-- Ingredients -->
          <div v-if="product.ingredients && product.ingredients.length" class="customization-section">
            <h3>Ingredients:</h3>
            <div class="ingredients-list">
              <label 
                v-for="ingredient in product.ingredients" 
                :key="ingredient.key"
                class="ingredient-item"
              >
                <input 
                  type="checkbox" 
                  v-model="selectedIngredients"
                  :value="ingredient.key"
                  :checked="ingredient.defaultIncluded"
                />
                <span>{{ ingredient.label }}</span>
              </label>
            </div>
          </div>

          <!-- Options -->
          <div v-if="product.options && product.options.length" class="customization-section">
            <div v-for="option in product.options" :key="option.key" class="option-group">
              <h3>{{ option.label }}:</h3>
              
              <!-- Single choice options -->
              <div v-if="option.type === 'single'" class="radio-group">
                <label 
                  v-for="choice in option.choices" 
                  :key="choice.value"
                  class="radio-item"
                >
                  <input 
                    type="radio" 
                    :name="option.key"
                    v-model="selectedOptions[option.key]"
                    :value="choice.value"
                  />
                  <span>
                    {{ choice.label }}
                    <span v-if="choice.extraPriceUSD > 0" class="extra-price">
                      +${{ choice.extraPriceUSD.toFixed(2) }}
                    </span>
                  </span>
                </label>
              </div>

              <!-- Multiple choice options -->
              <div v-else class="checkbox-group">
                <label 
                  v-for="choice in option.choices" 
                  :key="choice.value"
                  class="checkbox-item"
                >
                  <input 
                    type="checkbox" 
                    v-model="selectedOptions[option.key]"
                    :value="choice.value"
                  />
                  <span>
                    {{ choice.label }}
                    <span v-if="choice.extraPriceUSD > 0" class="extra-price">
                      +${{ choice.extraPriceUSD.toFixed(2) }}
                    </span>
                  </span>
                </label>
              </div>
            </div>
          </div>

          <!-- Quantity -->
          <div class="quantity-section">
            <h3>Quantity:</h3>
            <div class="quantity-controls">
              <button @click="decreaseQuantity" class="btn btn-sm">-</button>
              <span class="quantity-value">{{ quantity }}</span>
              <button @click="increaseQuantity" class="btn btn-sm">+</button>
            </div>
          </div>

          <!-- Add to Cart -->
          <div class="actions">
            <button @click="addToCart" class="btn btn-lg btn-primary" :disabled="addingToCart">
              <span v-if="addingToCart">Adding...</span>
              <span v-else>🛒 Add to Cart - ${{ (calculateTotalPrice() * quantity).toFixed(2) }}</span>
            </button>
            <router-link to="/menu" class="btn btn-lg btn-secondary">
              ← Back to Menu
            </router-link>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-xl">
        <p class="text-muted">Product not found</p>
        <router-link to="/menu" class="btn btn-primary">Back to Menu</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { productsAPI, cartAPI } from '@/api'
import { useCartStore } from '@/stores/cart'

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()

const product = ref(null)
const loading = ref(false)
const addingToCart = ref(false)

const selectedIngredients = ref([])
const selectedOptions = ref({})
const quantity = ref(1)

onMounted(async () => {
  await fetchProduct()
})

async function fetchProduct() {
  try {
    loading.value = true
    const slug = route.params.slug
    const { data } = await productsAPI.getBySlug(slug)
    product.value = data.data

    // Initialize selected ingredients (default included)
    if (product.value.ingredients) {
      selectedIngredients.value = product.value.ingredients
        .filter(ing => ing.defaultIncluded)
        .map(ing => ing.key)
    }

    // Initialize selected options (first choice by default)
    if (product.value.options) {
      product.value.options.forEach(option => {
        if (option.type === 'single' && option.choices.length > 0) {
          selectedOptions.value[option.key] = option.choices[0].value
        } else if (option.type === 'multiple') {
          selectedOptions.value[option.key] = []
        }
      })
    }
  } catch (err) {
    console.error('Failed to fetch product:', err)
  } finally {
    loading.value = false
  }
}

function calculateTotalPrice() {
  if (!product.value) return 0

  let total = product.value.priceUSD

  // Add extra price from options
  if (product.value.options) {
    product.value.options.forEach(option => {
      const selectedValue = selectedOptions.value[option.key]
      
      if (option.type === 'single') {
        const choice = option.choices.find(c => c.value === selectedValue)
        if (choice) {
          total += choice.extraPriceUSD || 0
        }
      } else if (option.type === 'multiple' && Array.isArray(selectedValue)) {
        selectedValue.forEach(value => {
          const choice = option.choices.find(c => c.value === value)
          if (choice) {
            total += choice.extraPriceUSD || 0
          }
        })
      }
    })
  }

  return total
}

function increaseQuantity() {
  quantity.value++
}

function decreaseQuantity() {
  if (quantity.value > 1) {
    quantity.value--
  }
}

async function addToCart() {
  try {
    addingToCart.value = true
    
    await cartStore.addItem(
      product.value.id,
      quantity.value,
      selectedIngredients.value,
      selectedOptions.value
    )

    // Show success notification (you can implement a toast notification here)
    alert('Product added to cart!')
    
    // Optionally redirect to cart
    // router.push('/cart')
  } catch (err) {
    console.error('Failed to add to cart:', err)
    alert('Failed to add product to cart')
  } finally {
    addingToCart.value = false
  }
}
</script>

<style scoped>
.product-page {
  padding: var(--spacing-xl) 0;
  min-height: 80vh;
}

.product-details {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-2xl);
  align-items: start;
}

.product-image {
  width: 100%;
  height: auto;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-lg);
}

.product-info-section h1 {
  font-size: var(--font-size-3xl);
  color: var(--contrast-black);
  margin-bottom: var(--spacing-md);
}

.product-description {
  color: var(--text);
  font-size: var(--font-size-lg);
  line-height: 1.6;
  margin-bottom: var(--spacing-lg);
}

.product-price {
  font-size: var(--font-size-2xl);
  font-weight: 700;
  color: var(--accent-red);
  margin-bottom: var(--spacing-xl);
}

.customization-section {
  margin-bottom: var(--spacing-xl);
  padding: var(--spacing-lg);
  background-color: var(--bg-cream);
  border-radius: var(--radius-md);
}

.customization-section h3 {
  margin-bottom: var(--spacing-md);
  color: var(--contrast-black);
}

.ingredients-list,
.radio-group,
.checkbox-group {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.ingredient-item,
.radio-item,
.checkbox-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  cursor: pointer;
  padding: var(--spacing-sm);
  border-radius: var(--radius-sm);
  transition: background-color var(--transition-fast);
}

.ingredient-item:hover,
.radio-item:hover,
.checkbox-item:hover {
  background-color: white;
}

.extra-price {
  color: var(--accent-red);
  font-weight: 600;
  margin-left: var(--spacing-xs);
}

.option-group {
  margin-bottom: var(--spacing-lg);
}

.quantity-section {
  margin-bottom: var(--spacing-xl);
}

.quantity-controls {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  margin-top: var(--spacing-md);
}

.quantity-value {
  font-size: var(--font-size-xl);
  font-weight: 600;
  min-width: 40px;
  text-align: center;
}

.actions {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

@media (max-width: 1024px) {
  .product-details {
    grid-template-columns: 1fr;
  }
}
</style>
