<template>
  <div class="mood-question-view-page">
    <div class="page-header">
      <h1>👁️ Question Details</h1>
      <div class="header-actions">
        <router-link :to="`/admin/mood-questions/${route.params.id}/edit`" class="btn btn-primary">✏️ Edit</router-link>
        <router-link to="/admin/mood-questions" class="btn btn-ghost">← Back</router-link>
      </div>
    </div>

    <div v-if="loading" class="text-center py-2xl"><div class="spinner"></div></div>

    <div v-else-if="question" class="question-details">
      <div class="details-card">
        <div class="card-header">
          <h2>❓ Question</h2>
          <span :class="['badge', 'badge-lg', question.isActive ? 'badge-success' : 'badge-error']">
            {{ question.isActive ? '✅ Active' : '❌ Inactive' }}
          </span>
        </div>
        <div class="details-grid">
          <div class="detail-item full-width">
            <strong>Text:</strong>
            <p>{{ question.text }}</p>
          </div>
          <div class="detail-item">
            <strong>Type:</strong>
            <span class="badge badge-secondary">{{ question.type }}</span>
          </div>
          <div class="detail-item">
            <strong>Order:</strong>
            <span>{{ question.order }}</span>
          </div>
          <div v-if="question.category" class="detail-item">
            <strong>Category:</strong>
            <span>{{ question.category }}</span>
          </div>
        </div>
      </div>

      <div v-if="question.options?.length" class="details-card">
        <div class="card-header">
          <h2>🎯 Options</h2>
        </div>
        <div class="options-grid">
          <div v-for="(option, index) in question.options" :key="index" class="option-card">
            <strong>{{ option.label }}</strong>
            <div class="text-muted">Value: {{ option.value }}</div>
            <div v-if="option.tags?.length" class="tags-list">
              <span v-for="(tag, tIndex) in option.tags" :key="tIndex" class="badge badge-primary">{{ tag }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useMoodStore } from '@/stores/mood'

const route = useRoute()
const moodStore = useMoodStore()
const loading = ref(true)
const question = ref(null)

onMounted(async () => {
  try {
    question.value = await moodStore.fetchQuestionById(route.params.id)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.mood-question-view-page {
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

.question-details {
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

.tags-list {
  margin-top: var(--spacing-sm);
  display: flex;
  gap: var(--spacing-xs);
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
</style>