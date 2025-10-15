<template>
  <div class="category-view-page">
    <div class="page-header">
      <h1>👁️ Category Details</h1>
      <div class="header-actions">
        <router-link 
          :to="`/admin/categories/${route.params.id}/edit`" 
          class="btn btn-primary"
        >
          ✏️ Edit
        </router-link>
        <router-link to="/admin/categories" class="btn btn-ghost">
          ← Back to list
        </router-link>
      </div>
    </div>

    <div v-if="loading" class="text-center py-2xl">
      <div class="spinner"></div>
      <p>Loading category...</p>
    </div>

    <div v-else-if="category" class="category-details">
      <div class="details-card">
        <div class="card-header">
          <h2>📋 Category Information</h2>
          <span 
            :class="['badge', 'badge-lg', category.isActive ? 'badge-success' : 'badge-error']"
          >
            {{ category.isActive ? '✅ Active' : '❌ Inactive' }}
          </span>
        </div>

        <div class="category-icon">
          {{ category.icon }}
        </div>

        <div class="details-grid">
          <div class="detail-item">
            <strong>Name:</strong>
            <span>{{ category.name }}</span>
          </div>
          <div class="detail-item">
            <strong>Slug:</strong>
            <span class="badge badge-secondary">{{ category.slug }}</span>
          </div>
          <div v-if="category.description" class="detail-item full-width">
            <strong>Description:</strong>
            <p>{{ category.description }}</p>
          </div>
          <div class="detail-item">
            <strong>Category ID:</strong>
            <span class="text-muted">{{ category.id }}</span>
          </div>
        </div>
      </div>

      <div class="details-card">
        <div class="card-header">
          <h2>📅 Metadata</h2>
        </div>
        <div class="details-grid">
          <div class="detail-item">
            <strong>Created:</strong>
            <span>{{ formatDate(category.createdAt) }}</span>
          </div>
          <div class="detail-item">
            <strong>Updated:</strong>
            <span>{{ formatDate(category.updatedAt) }}</span>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-2xl">
      <p class="text-error">Category not found</p>
      <router-link to="/admin/categories" class="btn btn-primary">
        Back to list
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { categoriesAPI } from '@/api'

const route = useRoute()
const loading = ref(true)
const category = ref(null)

onMounted(async () => {
  await loadCategory()
})

async function loadCategory() {
  try {
    loading.value = true
    const categoryId = route.params.id
    const { data } = await categoriesAPI.getById ? 
      await categoriesAPI.getById(categoryId) : 
      await categoriesAPI.getBySlug(categoryId)
    category.value = data.data
  } catch (err) {
    console.error('Load category error:', err)
  } finally {
    loading.value = false
  }
}

function formatDate(dateString) {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleString()
}
</script>

<style scoped>
.category-view-page {
  background-color: white;
  border-radius: var(--radius-lg);
  padding: var(--spacing-2xl);
  box-shadow: var(--shadow-sm);
  max-width: 900px;
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

.category-details {
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

.category-icon {
  font-size: 80px;
  text-align: center;
  margin: var(--spacing-lg) 0;
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

  .header-actions {
    flex-direction: column;
  }

  .details-grid {
    grid-template-columns: 1fr;
  }
}
</style>