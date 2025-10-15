<template>
  <div class="menu-page">
    <div class="container">
      <h1>🍔 Menu</h1>

      <!-- Filters -->
      <div class="filters">
        <div class="filter-group">
          <label>Category:</label>
          <select v-model="selectedCategory" class="form-select" @change="handleFilterChange">
            <option value="">All categories</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.slug">
              {{ cat.name }}
            </option>
          </select>
        </div>

        <div class="filter-group">
          <label>Sorting:</label>
          <select v-model="sortBy" class="form-select" @change="handleFilterChange">
            <option value="">Default</option>
            <option value="price-asc">Price: ascending</option>
            <option value="price-desc">Price: descending</option>
            <option value="name">By name</option>
          </select>
        </div>

        <div class="filter-group">
          <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="Search..." 
            class="form-input"
            @input="handleSearch"
          />
        </div>
      </div>

      <!-- Products Grid -->
      <div v-if="loading" class="text-center py-md">
        <div class="spinner"></div>
      </div>

      <div v-else-if="products.length" class="products-grid">
        <ProductCard 
          v-for="product in products" 
          :key="product.id"
          :product="product"
        />
      </div>

      <div v-else class="text-center py-md">
        <p class="text-muted">Products not found</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useProductsStore } from '@/stores/products'
import ProductCard from '@/components/ui/ProductCard.vue'

const productsStore = useProductsStore()

const selectedCategory = ref('')
const sortBy = ref('')
const searchQuery = ref('')
const loading = ref(false)

const products = ref([])
const categories = ref([])

onMounted(async () => {
  await fetchCategories()
  await fetchProducts()
})

async function fetchCategories() {
  await productsStore.fetchCategories({ active: true })
  categories.value = productsStore.categories
}

async function fetchProducts() {
  loading.value = true
  const params = {}
  
  if (selectedCategory.value) {
    params.category = selectedCategory.value
  }
  
  if (searchQuery.value) {
    params.search = searchQuery.value
  }

  await productsStore.fetchProducts(params)
  products.value = productsStore.products
  
  // Client-side sorting
  if (sortBy.value) {
    sortProducts()
  }
  
  loading.value = false
}

function sortProducts() {
  switch (sortBy.value) {
    case 'price-asc':
      products.value.sort((a, b) => a.priceUSD - b.priceUSD)
      break
    case 'price-desc':
      products.value.sort((a, b) => b.priceUSD - a.priceUSD)
      break
    case 'name':
      products.value.sort((a, b) => a.name.localeCompare(b.name))
      break
  }
}

function handleFilterChange() {
  fetchProducts()
}

let searchTimeout
function handleSearch() {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    fetchProducts()
  }, 300)
}
</script>

<style scoped>
.menu-page {
  min-height: 80vh;
  padding: var(--spacing-xl) 0;
}

.filters {
  display: flex;
  gap: var(--spacing-md);
  align-items: flex-end;
  margin-bottom: var(--spacing-2xl);
  padding: var(--spacing-lg);
  background-color: white;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
}

.filter-group {
  flex: 1;
  min-width: 0;
}

.filter-group:last-child {
  flex: 1.5;
}

.filter-group label {
  display: block;
  font-weight: 600;
  margin-bottom: var(--spacing-xs);
  color: var(--text);
}

@media (max-width: 768px) {
  .filters {
    flex-direction: column;
    align-items: stretch;
  }
  
  .filter-group:last-child {
    flex: 1;
  }
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: var(--spacing-lg);
}
</style>

