<template>
  <div class="categories-list-page">
    <div class="page-header">
      <h1>📂 Categories Management</h1>
      <router-link to="/admin/categories/create" class="btn btn-primary">
        ➕ Add Category
      </router-link>
    </div>

    <div v-if="loading" class="text-center py-md">
      <div class="spinner"></div>
    </div>

    <div v-else-if="categories.length" class="categories-table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th>Icon</th>
            <th>Name</th>
            <th>Slug</th>
            <th>Products</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="category in categories" :key="category.id">
            <td class="icon-cell">{{ category.icon }}</td>
            <td><strong>{{ category.name }}</strong></td>
            <td><span class="badge badge-secondary">{{ category.slug }}</span></td>
            <td>{{ category.productsCount || 0 }}</td>
            <td>
              <span :class="['badge', category.isActive ? 'badge-success' : 'badge-error']">
                {{ category.isActive ? 'Active' : 'Inactive' }}
              </span>
            </td>
            <td>
              <div class="action-buttons">
                <router-link 
                  :to="`/admin/categories/${category.id}`"
                  class="btn btn-sm btn-ghost"
                  title="View"
                >
                  👁️
                </router-link>
                <router-link 
                  :to="`/admin/categories/${category.id}/edit`"
                  class="btn btn-sm btn-ghost"
                  title="Edit"
                >
                  ✏️
                </router-link>
                <button 
                  @click="handleDelete(category)"
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
      <p class="text-muted">No categories found</p>
      <router-link to="/admin/categories/create" class="btn btn-primary mt-md">
        Add first category
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useCategoriesStore } from '@/stores/categories'

const categoriesStore = useCategoriesStore()
const categories = computed(() => categoriesStore.categories)
const loading = ref(false)

onMounted(async () => {
  loading.value = true
  await categoriesStore.fetchCategories()
  loading.value = false
})

async function handleDelete(category) {
  if (!confirm(`Delete category "${category.name}"?`)) {
    return
  }

  try {
    await categoriesStore.deleteCategory(category.id)
    alert('Category deleted successfully!')
  } catch (err) {
    alert('Failed to delete category')
  }
}
</script>

<style scoped>
.categories-list-page {
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

.categories-table-container {
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

.icon-cell {
  font-size: var(--font-size-2xl);
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