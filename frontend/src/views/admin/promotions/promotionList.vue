<template>
  <div class="promotions-list-page">
    <div class="page-header">
      <h1>🎉 Promotions Management</h1>
      <router-link to="/admin/promotions/create" class="btn btn-primary">
        ➕ Add Promotion
      </router-link>
    </div>

    <div v-if="loading" class="text-center py-md">
      <div class="spinner"></div>
    </div>

    <div v-else-if="promotions.length" class="promotions-table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th>Banner</th>
            <th>Title</th>
            <th>Period</th>
            <th>Products</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="promotion in promotions" :key="promotion.id">
            <td>
              <img :src="promotion.bannerImage" :alt="promotion.title" class="table-image" />
            </td>
            <td>
              <strong>{{ promotion.title }}</strong>
              <div class="text-muted" style="font-size: 0.875rem;">
                {{ promotion.description?.substring(0, 50) }}...
              </div>
            </td>
            <td>
              <div class="date-range">
                <div>{{ formatDate(promotion.startsAt) }}</div>
                <div>→</div>
                <div>{{ formatDate(promotion.endsAt) }}</div>
              </div>
            </td>
            <td>{{ promotion.appliesTo?.length || 0 }}</td>
            <td>
              <span :class="['badge', promotion.isActive ? 'badge-success' : 'badge-error']">
                {{ promotion.isActive ? 'Active' : 'Inactive' }}
              </span>
            </td>
            <td>
              <div class="action-buttons">
                <router-link 
                  :to="`/admin/promotions/${promotion.id}`"
                  class="btn btn-sm btn-ghost"
                  title="View"
                >
                  👁️
                </router-link>
                <router-link 
                  :to="`/admin/promotions/${promotion.id}/edit`"
                  class="btn btn-sm btn-ghost"
                  title="Edit"
                >
                  ✏️
                </router-link>
                <button 
                  @click="handleDelete(promotion)"
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
      <p class="text-muted">No promotions found</p>
      <router-link to="/admin/promotions/create" class="btn btn-primary mt-md">
        Add first promotion
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { usePromotionsStore } from '@/stores/promotions'

const promotionsStore = usePromotionsStore()
const promotions = computed(() => promotionsStore.promotions)
const loading = ref(false)

onMounted(async () => {
  loading.value = true
  await promotionsStore.fetchPromotions()
  loading.value = false
})

function formatDate(dateString) {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleDateString()
}

async function handleDelete(promotion) {
  if (!confirm(`Delete promotion "${promotion.title}"?`)) {
    return
  }

  try {
    await promotionsStore.deletePromotion(promotion.id)
    alert('Promotion deleted successfully!')
  } catch (err) {
    alert('Failed to delete promotion')
  }
}
</script>

<style scoped>
.promotions-list-page {
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

.promotions-table-container {
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
  width: 100px;
  height: 60px;
  object-fit: cover;
  border-radius: var(--radius-sm);
}

.date-range {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
  font-size: var(--font-size-sm);
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