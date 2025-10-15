<template>
  <div class="product-view-page">
    <div class="page-header">
      <h1>👁️ Product Details</h1>
      <div class="header-actions">
        <router-link 
          :to="`/admin/products/${route.params.id}/edit`" 
          class="btn btn-primary"
        >
          ✏️ Edit
        </router-link>
        <router-link to="/admin/products" class="btn btn-ghost">
          ← Back to list
        </router-link>
      </div>
    </div>

    <div v-if="loading" class="text-center py-2xl">
      <div class="spinner"></div>
      <p>Loading product...</p>
    </div>

    <div v-else-if="product" class="product-details">
      <!-- Basic Info Card -->
      <div class="details-card">
        <div class="card-header">
          <h2>📋 Basic Information</h2>
          <span 
            :class="['badge', 'badge-lg', product.isActive ? 'badge-success' : 'badge-error']"
          >
            {{ product.isActive ? '✅ Active' : '❌ Inactive' }}
          </span>
        </div>

        <div class="details-grid">
          <div class="detail-item">
            <strong>Name:</strong>
            <span>{{ product.name }}</span>
          </div>
          <div class="detail-item">
            <strong>Slug:</strong>
            <span class="badge badge-secondary">{{ product.slug }}</span>
          </div>
          <div class="detail-item full-width">
            <strong>Description:</strong>
            <p>{{ product.description }}</p>
          </div>
          <div class="detail-item">
            <strong>Category:</strong>
            <span>{{ getCategoryName(product.categoryId) }}</span>
          </div>
          <div class="detail-item">
            <strong>Product ID:</strong>
            <span class="text-muted">{{ product.id }}</span>
          </div>
        </div>
      </div>

      <!-- Image Card -->
      <div class="details-card">
        <div class="card-header">
          <h2>🖼️ Image</h2>
        </div>
        <div class="product-image">
          <img :src="product.image" :alt="product.name" />
        </div>
      </div>

      <!-- Pricing Card -->
      <div class="details-card">
        <div class="card-header">
          <h2>💰 Pricing</h2>
        </div>
        <div class="pricing-info">
          <div class="price-item">
            <strong>Current Price:</strong>
            <span class="price-current">${{ product.priceUSD.toFixed(2) }}</span>
          </div>
          <div v-if="product.oldPriceUSD > 0" class="price-item">
            <strong>Old Price:</strong>
            <span class="price-old">${{ product.oldPriceUSD.toFixed(2) }}</span>
          </div>
          <div v-if="product.oldPriceUSD > 0" class="price-item">
            <strong>Discount:</strong>
            <span class="badge badge-error">
              -{{ Math.round((1 - product.priceUSD / product.oldPriceUSD) * 100) }}%
            </span>
          </div>
        </div>
      </div>

      <!-- Tags Card -->
      <div v-if="product.tags?.length" class="details-card">
        <div class="card-header">
          <h2>🏷️ Tags</h2>
        </div>
        <div class="tags-list">
          <span 
            v-for="(tag, index) in product.tags" 
            :key="index" 
            class="badge badge-primary badge-lg"
          >
            {{ tag }}
          </span>
        </div>
      </div>

      <!-- Options Card -->
      <div v-if="product.options?.length" class="details-card full-width">
        <div class="card-header">
          <h2>⚙️ Product Options</h2>
        </div>
        <div class="options-grid">
          <div 
            v-for="(option, index) in product.options" 
            :key="index" 
            class="option-card"
          >
            <div class="option-header">
              <h3>{{ option.name }}</h3>
              <div class="option-badges">
                <span class="badge badge-secondary">{{ option.type }}</span>
                <span 
                  v-if="option.required" 
                  class="badge badge-warning"
                >
                  Required
                </span>
              </div>
            </div>
            <div class="option-values">
              <span 
                v-for="(value, vIndex) in option.values" 
                :key="vIndex" 
                class="value-badge"
              >
                {{ value }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Ingredients Card -->
      <div v-if="product.ingredients?.length" class="details-card full-width">
        <div class="card-header">
          <h2>🥬 Ingredients</h2>
        </div>
        <div class="ingredients-grid">
          <span 
            v-for="(ingredient, index) in product.ingredients" 
            :key="index" 
            class="ingredient-badge"
          >
            {{ ingredient.name }}
          </span>
        </div>
      </div>

      <!-- Metadata Card -->
      <div class="details-card full-width">
        <div class="card-header">
          <h2>📅 Metadata</h2>
        </div>
        <div class="metadata-grid">
          <div class="detail-item">
            <strong>Created:</strong>
            <span>{{ formatDate(product.createdAt) }}</span>
          </div>
          <div class="detail-item">
            <strong>Updated:</strong>
            <span>{{ formatDate(product.updatedAt) }}</span>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-2xl">
      <p class="text-error">Product not found</p>
      <router-link to="/admin/products" class="btn btn-primary">
        Back to list
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useProductsStore } from '@/stores/products'

const route = useRoute()
const productsStore = useProductsStore()

const loading = ref(true)
const product = computed(() => productsStore.product)
const categories = computed(() => productsStore.categories)

onMounted(async () => {
  await Promise.all([
    productsStore.fetchCategories(),
    loadProduct()
  ])
})

async function loadProduct() {
  try {
    loading.value = true
    await productsStore.fetchProductById(route.params.id)
  } catch (err) {
    console.error('Load product error:', err)
  } finally {
    loading.value = false
  }
}

function getCategoryName(categoryId) {
  const category = categories.value.find(c => c.id === categoryId)
  return category?.name || 'No category'
}

function formatDate(dateString) {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleString()
}
</script>

<style scoped>
.product-view-page {
  background-color: white;
  border-radius: var(--radius-lg);
  padding: var(--spacing-2xl);
  box-shadow: var(--shadow-sm);
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

.header-actions {
  display: flex;
  gap: var(--spacing-sm);
}

.product-details {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-xl);
}

.details-card {
  background-color: var(--bg-cream);
  padding: var(--spacing-lg);
  border-radius: var(--radius-md);
  border: 2px solid var(--border);
}

.details-card.full-width {
  grid-column: 1 / -1;
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
  letter-spacing: 0.5px;
}

.detail-item span,
.detail-item p {
  color: var(--text);
  font-size: var(--font-size-lg);
}

.product-image {
  border-radius: var(--radius-md);
  overflow: hidden;
  max-width: 100%;
}

.product-image img {
  width: 100%;
  height: auto;
  display: block;
}

.pricing-info {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.price-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price-current {
  font-size: var(--font-size-2xl);
  font-weight: 700;
  color: var(--accent-red);
}

.price-old {
  font-size: var(--font-size-lg);
  text-decoration: line-through;
  color: var(--text-light);
}

.tags-list {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
}

.options-grid {
  display: grid;
  gap: var(--spacing-md);
}

.option-card {
  background-color: white;
  padding: var(--spacing-md);
  border-radius: var(--radius-sm);
  border: 2px solid var(--border);
}

.option-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-sm);
}

.option-header h3 {
  margin: 0;
  font-size: var(--font-size-lg);
}

.option-badges {
  display: flex;
  gap: var(--spacing-xs);
}

.option-values {
  display: flex;
  gap: var(--spacing-xs);
  flex-wrap: wrap;
}

.value-badge {
  padding: var(--spacing-xs) var(--spacing-sm);
  background-color: var(--bg-cream);
  border: 1px solid var(--border);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-sm);
}

.ingredients-grid {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
}

.ingredient-badge {
  padding: var(--spacing-sm) var(--spacing-md);
  background-color: white;
  border: 2px solid var(--border);
  border-radius: var(--radius-md);
  font-weight: 500;
}

.metadata-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-md);
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
  .product-details {
    grid-template-columns: 1fr;
  }

  .page-header {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: stretch;
  }

  .header-actions {
    flex-direction: column;
  }

  .details-grid,
  .metadata-grid {
    grid-template-columns: 1fr;
  }
}
</style>