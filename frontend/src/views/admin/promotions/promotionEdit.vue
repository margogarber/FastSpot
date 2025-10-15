<template>
  <div class="promotion-edit-page">
    <div class="page-header">
      <h1>✏️ Edit Promotion</h1>
      <router-link to="/admin/promotions" class="btn btn-ghost">← Back</router-link>
    </div>

    <div v-if="loadingPromotion" class="text-center py-2xl">
      <div class="spinner"></div>
    </div>

    <form v-else-if="form.title" @submit.prevent="handleSubmit" class="promotion-form">
      <div class="form-section">
        <h3>Basic Information</h3>
        <div class="form-group">
          <label>Title *</label>
          <input v-model="form.title" type="text" class="form-input" required />
        </div>
        <div class="form-group">
          <label>Description</label>
          <textarea v-model="form.description" class="form-input" rows="3"></textarea>
        </div>
        <div class="form-group">
          <label>Banner Image URL *</label>
          <input v-model="form.bannerImage" type="url" class="form-input" required />
          <div v-if="form.bannerImage" class="image-preview">
            <img :src="form.bannerImage" alt="Preview" />
          </div>
        </div>
      </div>

      <div class="form-section">
        <h3>Period</h3>
        <div class="form-row">
          <div class="form-group">
            <label>Starts At *</label>
            <input v-model="form.startsAt" type="datetime-local" class="form-input" required />
          </div>
          <div class="form-group">
            <label>Ends At *</label>
            <input v-model="form.endsAt" type="datetime-local" class="form-input" required />
          </div>
        </div>
      </div>

      <div class="form-section">
        <h3>Applies To Products</h3>
        <div class="form-group">
          <input v-model="appliesToInput" type="text" class="form-input" />
          <div v-if="form.appliesTo.length" class="tags-preview">
            <span v-for="(slug, index) in form.appliesTo" :key="index" class="badge badge-primary">{{ slug }}</span>
          </div>
        </div>
      </div>

      <div class="form-section">
        <label class="checkbox-label">
          <input v-model="form.isActive" type="checkbox" />
          <span>Active</span>
        </label>
      </div>

      <div class="form-actions">
        <button type="submit" class="btn btn-primary btn-lg" :disabled="loading">
          {{ loading ? 'Saving...' : '💾 Save' }}
        </button>
        <router-link to="/admin/promotions" class="btn btn-ghost btn-lg">Cancel</router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { usePromotionsStore } from '@/stores/promotions'

const router = useRouter()
const route = useRoute()
const promotionsStore = usePromotionsStore()
const loading = ref(false)
const loadingPromotion = ref(true)

const form = ref({
  title: '',
  description: '',
  startsAt: '',
  endsAt: '',
  bannerImage: '',
  appliesTo: [],
  isActive: true
})

const appliesToInput = ref('')

watch(appliesToInput, (val) => {
  form.value.appliesTo = val.split(',').map(s => s.trim()).filter(s => s.length > 0)
})

onMounted(async () => {
  await loadPromotion()
})

async function loadPromotion() {
  try {
    const promo = await promotionsStore.fetchPromotionById(route.params.id)
    form.value = {
      title: promo.title || '',
      description: promo.description || '',
      startsAt: promo.startsAt ? new Date(promo.startsAt).toISOString().slice(0, 16) : '',
      endsAt: promo.endsAt ? new Date(promo.endsAt).toISOString().slice(0, 16) : '',
      bannerImage: promo.bannerImage || '',
      appliesTo: promo.appliesTo || [],
      isActive: promo.isActive !== undefined ? promo.isActive : true
    }
    appliesToInput.value = form.value.appliesTo.join(', ')
  } catch (err) {
    alert('Failed to load')
  } finally {
    loadingPromotion.value = false
  }
}

async function handleSubmit() {
  try {
    loading.value = true
    
    // Validate dates
    const start = new Date(form.value.startsAt)
    const end = new Date(form.value.endsAt)
    
    if (start >= end) {
      alert('❌ Start date must be before end date!')
      loading.value = false
      return
    }
    
    const data = {
      ...form.value,
      startsAt: start.toISOString(),
      endsAt: end.toISOString()
    }
    await promotionsStore.updatePromotion(route.params.id, data)
    alert('Promotion updated! ✅')
    router.push('/admin/promotions')
  } catch (err) {
    alert('Failed: ' + (err.response?.data?.error || err.message))
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.promotion-edit-page {
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
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-md);
}

.image-preview {
  margin-top: var(--spacing-md);
  border-radius: var(--radius-md);
  overflow: hidden;
  max-width: 400px;
}

.image-preview img {
  width: 100%;
  display: block;
}

.tags-preview {
  margin-top: var(--spacing-sm);
  display: flex;
  gap: var(--spacing-xs);
  flex-wrap: wrap;
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

@media (max-width: 768px) {
  .page-header, .form-row, .form-actions {
    flex-direction: column;
  }
}
</style>