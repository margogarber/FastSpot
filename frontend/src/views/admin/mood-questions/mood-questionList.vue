<template>
  <div class="mood-questions-list-page">
    <div class="page-header">
      <h1>❓ Mood Questions</h1>
      <router-link to="/admin/mood-questions/create" class="btn btn-primary">➕ Add Question</router-link>
    </div>

    <div v-if="loading" class="text-center py-md"><div class="spinner"></div></div>

    <div v-else-if="questions.length" class="questions-table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th>#</th>
            <th>Question</th>
            <th>Type</th>
            <th>Options</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="question in questions" :key="question.id">
            <td>{{ question.order }}</td>
            <td><strong>{{ question.text }}</strong></td>
            <td><span class="badge badge-secondary">{{ question.type }}</span></td>
            <td>{{ question.options?.length || 0 }}</td>
            <td><span :class="['badge', question.isActive ? 'badge-success' : 'badge-error']">{{ question.isActive ? 'Active' : 'Inactive' }}</span></td>
            <td>
              <div class="action-buttons">
                <router-link :to="`/admin/mood-questions/${question.id}`" class="btn btn-sm btn-ghost">👁️</router-link>
                <router-link :to="`/admin/mood-questions/${question.id}/edit`" class="btn btn-sm btn-ghost">✏️</router-link>
                <button @click="handleDelete(question)" class="btn btn-sm btn-ghost">🗑️</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-else class="text-center py-md">
      <p class="text-muted">No questions found</p>
      <router-link to="/admin/mood-questions/create" class="btn btn-primary mt-md">Add first question</router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useMoodStore } from '@/stores/mood'

const moodStore = useMoodStore()
const questions = computed(() => moodStore.questions)
const loading = ref(false)

onMounted(async () => {
  loading.value = true
  await moodStore.fetchAllQuestions()
  loading.value = false
})

async function handleDelete(question) {
  if (!confirm(`Delete question "${question.text}"?`)) return
  try {
    await moodStore.deleteQuestion(question.id)
    alert('Question deleted!')
  } catch (err) {
    alert('Failed to delete')
  }
}
</script>

<style scoped>
.mood-questions-list-page {
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

.questions-table-container {
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
  border-bottom: 2px solid var(--border);
}

.data-table td {
  padding: var(--spacing-md);
  border-bottom: 1px solid var(--border);
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
</style>