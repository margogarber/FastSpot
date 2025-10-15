<template>
  <div class="products-list-page">
    <div class="page-header">
      <h1>🍔 Products management</h1>
      <router-link to="/admin/products/create" class="btn btn-primary">
        ➕ Add product
      </router-link>
    </div>

    <div v-if="loading" class="text-center py-md">
      <div class="spinner"></div>
    </div>

    <div v-else-if="products.length" class="products-table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th>Image</th>
            <th>Name</th>
            <th>Category</th>
            <th>Price</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="product in products" :key="product?.id || Math.random()">
            <td>
              <img :src="product?.image || ''" :alt="product?.name || ''" class="table-image" />
            </td>
            <td>
              <strong>{{ product?.name || 'N/A' }}</strong>
              <div class="text-muted" style="font-size: 0.875rem;">
                {{ product?.slug || 'N/A' }}
              </div>
            </td>
            <td>
              {{ getCategoryName(product?.categoryId) }}
            </td>
            <td class="text-primary">
              ${{ (product?.priceUSD || 0).toFixed(2) }}
            </td>
            <td>
              <span 
                :class="['badge', product?.isActive ? 'badge-success' : 'badge-error']"
              >
                {{ product?.isActive ? 'Active' : 'Inactive' }}
              </span>
            </td>
            <td>
              <div class="action-buttons">
                <router-link 
                  :to="`/admin/products/${product.id}`"
                  class="btn btn-sm btn-ghost"
                  title="View"
                >
                  👁️
                </router-link>
                <router-link 
                  :to="`/admin/products/${product.id}/edit`"
                  class="btn btn-sm btn-ghost"
                  title="Edit"
                >
                  ✏️
                </router-link>
                <button 
                  @click="handleDelete(product)"
                  class="btn btn-sm btn-ghost"
                  title="Delete"
                >
                  🗑️
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-else class="text-center py-md">
    <p class="text-muted">Products not found</p>
      <router-link to="/admin/products/create" class="btn btn-primary mt-md">
        Add first product
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useProductsStore } from '@/stores/products'

const productsStore = useProductsStore()

const products = computed(() => productsStore.products || [])
const loading = ref(false)

onMounted(async () => {
  await fetchData()
})

async function fetchData() {
  loading.value = true
  await Promise.all([
    productsStore.fetchProducts(),
    productsStore.fetchCategories()
  ])
  loading.value = false
}

function getCategoryName(categoryId) {
  const category = productsStore.getCategoryById(categoryId)
  return category?.name || 'No category'
}

async function handleDelete(product) {
  if (!confirm(`Delete product "${product.name}"?`)) {
    return
  }

  try {
    await productsStore.deleteProduct(product.id)
    alert('Product deleted successfully!')
  } catch (err) {
    alert('Failed to delete product')
  }
}
</script>

<style scoped>
.products-list-page {
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
}

.page-header h1 {
  margin: 0;
}

.products-table-container {
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
  color: var(--text);
  border-bottom: 2px solid var(--border);
}

.data-table td {
  padding: var(--spacing-md);
  border-bottom: 1px solid var(--border);
}

.table-image {
  width: 60px;
  height: 60px;
  object-fit: cover;
  border-radius: var(--radius-sm);
}

.action-buttons {
  display: flex;
  gap: var(--spacing-xs);
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

