<template>
  <div class="promotions-page">
    <div class="container">
      <h1>🎉 Current promotions</h1>

      <div v-if="loading" class="text-center py-xl">
        <div class="spinner"></div>
      </div>

      <div v-else-if="promotions.length" class="promotions-grid">
        <div 
          v-for="promo in promotions" 
          :key="promo.id"
          class="promo-card"
        >
          <img :src="promo.bannerImage" :alt="promo.title" class="promo-image" />
          <div class="promo-content">
            <h2>{{ promo.title }}</h2>
            <p class="promo-description">{{ promo.description }}</p>
            <div class="promo-dates">
              <span class="promo-date">From: {{ formatDate(promo.startsAt) }}</span>
              <span class="promo-date">To: {{ formatDate(promo.endsAt) }}</span>
            </div>
            <router-link to="/menu" class="btn btn-primary">
              View menu
            </router-link>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-xl text-muted">
        <p>No active promotions at the moment</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { promotionsAPI } from '@/api'

const promotions = ref([])
const loading = ref(false)

onMounted(async () => {
  await fetchPromotions()
})

async function fetchPromotions() {
  try {
    loading.value = true
    const { data } = await promotionsAPI.getAll({ active: true })
    promotions.value = data.data.promotions
  } catch (err) {
    console.error('Failed to fetch promotions:', err)
  } finally {
    loading.value = false
  }
}

function formatDate(dateString) {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  })
}
</script>

<style scoped>
.promotions-page {
  min-height: 80vh;
  padding: var(--spacing-xl) 0;
}

.promotions-page h1 {
  text-align: center;
  margin-bottom: var(--spacing-2xl);
  color: var(--contrast-black);
}

.promotions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: var(--spacing-xl);
}

.promo-card {
  background-color: white;
  border-radius: var(--radius-lg);
  overflow: hidden;
  box-shadow: var(--shadow-md);
  transition: transform var(--transition-base), box-shadow var(--transition-base);
}

.promo-card:hover {
  transform: translateY(-8px);
  box-shadow: var(--shadow-lg);
}

.promo-image {
  width: 100%;
  height: 250px;
  object-fit: cover;
}

.promo-content {
  padding: var(--spacing-xl);
}

.promo-content h2 {
  color: var(--accent-red);
  margin-bottom: var(--spacing-md);
  font-size: var(--font-size-2xl);
}

.promo-description {
  color: var(--text);
  margin-bottom: var(--spacing-lg);
  line-height: 1.6;
}

.promo-dates {
  display: flex;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-lg);
  font-size: var(--font-size-sm);
  color: var(--text-muted);
}

.promo-date {
  display: inline-flex;
  align-items: center;
  gap: var(--spacing-xs);
}

@media (max-width: 768px) {
  .promotions-grid {
    grid-template-columns: 1fr;
  }
}
</style>
