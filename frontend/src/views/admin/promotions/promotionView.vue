<template>
  <div class="promotion-view-page">
    <div class="page-header">
      <h1>👁️ Promotion Details</h1>
      <div class="header-actions">
        <router-link :to="`/admin/promotions/${route.params.id}/edit`" class="btn btn-primary">✏️ Edit</router-link>
        <router-link to="/admin/promotions" class="btn btn-ghost">← Back</router-link>
      </div>
    </div>

    <div v-if="loading" class="text-center py-2xl">
      <div class="spinner"></div>
    </div>

    <div v-else-if="promotion" class="promotion-details">
      <div class="details-card">
        <div class="card-header">
          <h2>📋 Information</h2>
          <span :class="['badge', 'badge-lg', promotion.isActive ? 'badge-success' : 'badge-error']">
            {{ promotion.isActive ? '✅ Active' : '❌ Inactive' }}
          </span>
        </div>
        <div class="promotion-banner">
          <img :src="promotion.bannerImage" :alt="promotion.title" />
        </div>
        <div class="details-grid">
          <div class="detail-item full-width">
            <strong>Title:</strong>
            <span>{{ promotion.title }}</span>
          </div>
          <div class="detail-item full-width">
            <strong>Description:</strong>
            <p>{{ promotion.description }}</p>
          </div>
          <div class="detail-item">
            <strong>Starts:</strong>
            <span>{{ formatDate(promotion.startsAt) }}</span>
          </div>
          <div class="detail-item">
            <strong>Ends:</strong>
            <span>{{ formatDate(promotion.endsAt) }}</span>
          </div>
        </div>
      </div>

      <div v-if="promotion.appliesTo?.length" class="details-card">
        <div class="card-header">
          <h2>🎯 Applies To</h2>
        </div>
        <div class="applies-to-list">
          <span v-for="(slug, index) in promotion.appliesTo" :key="index" class="badge badge-primary badge-lg">
            {{ slug }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { usePromotionsStore } from '@/stores/promotions'

const route = useRoute()
const promotionsStore = usePromotionsStore()
const loading = ref(true)
const promotion = ref(null)

onMounted(async () => {
  try {
    promotion.value = await promotionsStore.fetchPromotionById(route.params.id)
  } finally {
    loading.value = false
  }
})

function formatDate(dateString) {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleString()
}
</script>

<style scoped>
.promotion-view-page {
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

.promotion-details {
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

.promotion-banner {
  margin-bottom: var(--spacing-lg);
  border-radius: var(--radius-md);
  overflow: hidden;
}

.promotion-banner img {
  width: 100%;
  display: block;
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
}

.detail-item span, .detail-item p {
  color: var(--text);
  font-size: var(--font-size-lg);
}

.applies-to-list {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
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