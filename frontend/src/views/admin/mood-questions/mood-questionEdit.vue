<template>
  <div class="mood-question-edit-page">
    <div class="page-header">
      <h1>✏️ Edit Mood Question</h1>
      <router-link to="/admin/mood-questions" class="btn btn-ghost">← Back</router-link>
    </div>

    <div v-if="loadingQuestion" class="text-center py-2xl"><div class="spinner"></div></div>

    <form v-else-if="form.text" @submit.prevent="handleSubmit" class="question-form">
      <div class="form-section">
        <h3>Question Details</h3>
        <div class="form-group">
          <label>Question Text *</label>
          <input v-model="form.text" type="text" class="form-input" required />
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>Type *</label>
            <select v-model="form.type" class="form-input" required>
              <option value="single">Single Choice</option>
              <option value="multiple">Multiple Choice</option>
            </select>
          </div>
          <div class="form-group">
            <label>Order *</label>
            <input v-model.number="form.order" type="number" class="form-input" required />
          </div>
          <div class="form-group">
            <label>Category</label>
            <select v-model="form.category" class="form-input">
              <option value="">None</option>
              <option value="mood">Mood</option>
              <option value="preference">Preference</option>
              <option value="occasion">Occasion</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <label class="checkbox-label">
            <input v-model="form.isActive" type="checkbox" />
            <span>Active</span>
          </label>
        </div>
      </div>

      <div class="form-section">
        <div class="section-header">
          <h3>Options</h3>
          <button type="button" @click="addOption" class="btn btn-sm btn-secondary">➕ Add Option</button>
        </div>
        <div v-if="form.options.length" class="options-list">
          <div v-for="(option, index) in form.options" :key="index" class="option-item">
            <div class="form-row">
              <div class="form-group">
                <label>Value</label>
                <input v-model="option.value" type="text" class="form-input" />
              </div>
              <div class="form-group">
                <label>Label</label>
                <input v-model="option.label" type="text" class="form-input" />
              </div>
              <button type="button" @click="removeOption(index)" class="btn btn-sm btn-ghost">🗑️</button>
            </div>
            <div class="form-group">
              <label>Tags</label>
              <input v-model="option.tagsInput" @input="updateOptionTags(option)" type="text" class="form-input" />
            </div>
          </div>
        </div>
      </div>

      <div class="form-actions">
        <button type="submit" class="btn btn-primary btn-lg" :disabled="loading">
          {{ loading ? 'Saving...' : '💾 Save' }}
        </button>
        <router-link to="/admin/mood-questions" class="btn btn-ghost btn-lg">Cancel</router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useMoodStore } from '@/stores/mood'

const router = useRouter()
const route = useRoute()
const moodStore = useMoodStore()
const loading = ref(false)
const loadingQuestion = ref(true)

const form = ref({
  text: '',
  type: 'single',
  order: 1,
  category: '',
  isActive: true,
  options: []
})

onMounted(async () => {
  await loadQuestion()
})

async function loadQuestion() {
  try {
    const question = await moodStore.fetchQuestionById(route.params.id)
    form.value = {
      text: question.text || '',
      type: question.type || 'single',
      order: question.order || 1,
      category: question.category || '',
      isActive: question.isActive !== undefined ? question.isActive : true,
      options: (question.options || []).map(opt => ({
        ...opt,
        tagsInput: opt.tags?.join(', ') || ''
      }))
    }
  } finally {
    loadingQuestion.value = false
  }
}

function addOption() {
  form.value.options.push({ value: '', label: '', tags: [], tagsInput: '' })
}

function removeOption(index) {
  form.value.options.splice(index, 1)
}

function updateOptionTags(option) {
  option.tags = option.tagsInput.split(',').map(t => t.trim()).filter(t => t.length > 0)
}

async function handleSubmit() {
  try {
    loading.value = true
    const data = {
      ...form.value,
      options: form.value.options.map(opt => ({
        value: opt.value,
        label: opt.label,
        tags: opt.tags
      }))
    }
    await moodStore.updateQuestion(route.params.id, data)
    alert('Question updated! ✅')
    router.push('/admin/mood-questions')
  } catch (err) {
    alert('Failed: ' + (err.response?.data?.error || err.message))
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* Same styles as Create page */
.mood-question-edit-page {
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

.form-section {
  background-color: var(--bg-cream);
  padding: var(--spacing-lg);
  border-radius: var(--radius-md);
  margin-bottom: var(--spacing-xl);
}

.form-section h3 {
  margin: 0 0 var(--spacing-lg) 0;
  color: var(--accent-red);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.section-header h3 {
  margin: 0;
}

.form-group {
  margin-bottom: var(--spacing-md);
}

.form-group:last-child {
  margin-bottom: 0;
}

.form-group label {
  display: block;
  margin-bottom: var(--spacing-xs);
  font-weight: 600;
}

.form-input {
  width: 100%;
  padding: var(--spacing-sm);
  border: 2px solid var(--border);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-base);
}

.form-input:focus {
  outline: none;
  border-color: var(--accent-red);
}

.form-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-md);
  align-items: end;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  cursor: pointer;
  font-weight: 600;
}

.checkbox-label input {
  width: 20px;
  height: 20px;
}

.options-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.option-item {
  background-color: white;
  padding: var(--spacing-md);
  border-radius: var(--radius-sm);
  border: 2px solid var(--border);
}

.form-actions {
  display: flex;
  gap: var(--spacing-md);
  padding-top: var(--spacing-xl);
  border-top: 2px solid var(--border);
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