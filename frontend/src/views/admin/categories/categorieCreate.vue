<template>
  <div class="category-create-page">
    <div class="page-header">
      <h1>➕ Create Category</h1>
      <router-link to="/admin/categories" class="btn btn-ghost">
        ← Back to list
      </router-link>
    </div>

    <form @submit.prevent="handleSubmit" class="category-form">
      <div class="form-section">
        <h3>Category Information</h3>
        
        <div class="form-group">
          <label for="name">Name *</label>
          <input 
            id="name"
            v-model="form.name" 
            type="text" 
            class="form-input"
            placeholder="e.g. Burgers"
            required
          />
        </div>

        <div class="form-group">
          <label for="slug">Slug * <small>(URL-friendly name)</small></label>
          <input 
            id="slug"
            v-model="form.slug" 
            type="text" 
            class="form-input"
            placeholder="e.g. burgers"
            required
            pattern="[a-z0-9-]+"
          />
          <small class="form-hint">Only lowercase letters, numbers, and hyphens</small>
        </div>

        <div class="form-group">
          <label for="icon">Icon (emoji) *</label>
          <input 
            id="icon"
            v-model="form.icon" 
            type="text" 
            class="form-input icon-input"
            placeholder="🍔"
            required
            maxlength="2"
          />
        </div>

        <div class="form-group">
          <label for="description">Description</label>
          <textarea 
            id="description"
            v-model="form.description" 
            class="form-input"
            rows="3"
            placeholder="Optional description..."
          ></textarea>
        </div>

        <div class="form-group">
          <label class="checkbox-label">
            <input 
              v-model="form.isActive" 
              type="checkbox"
            />
            <span>Active (visible to customers)</span>
          </label>
        </div>
      </div>

      <div class="form-actions">
        <button 
          type="submit" 
          class="btn btn-primary btn-lg"
          :disabled="loading"
        >
          {{ loading ? 'Creating...' : '✅ Create Category' }}
        </button>
        <router-link 
          to="/admin/categories" 
          class="btn btn-ghost btn-lg"
        >
          Cancel
        </router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useCategoriesStore } from '@/stores/categories'

const router = useRouter()
const categoriesStore = useCategoriesStore()

const loading = ref(false)
const form = ref({
  name: '',
  slug: '',
  icon: '',
  description: '',
  isActive: true
})

// Auto-generate slug from name
watch(() => form.value.name, (newName) => {
  if (!form.value.slug) {
    form.value.slug = newName
      .toLowerCase()
      .replace(/[^a-z0-9]+/g, '-')
      .replace(/^-|-$/g, '')
  }
})

async function handleSubmit() {
  try {
    loading.value = true
    await categoriesStore.createCategory(form.value)
    alert('Category created successfully! ✅')
    router.push('/admin/categories')
  } catch (err) {
    console.error('Create category error:', err)
    alert('Failed to create category: ' + (err.response?.data?.error || err.message))
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.category-create-page {
  background-color: white;
  border-radius: var(--radius-lg);
  padding: var(--spacing-2xl);
  box-shadow: var(--shadow-sm);
  max-width: 800px;
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
  margin-top: 0;
  margin-bottom: var(--spacing-lg);
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
  color: var(--text);
}

.form-input {
  width: 100%;
  padding: var(--spacing-sm);
  border: 2px solid var(--border);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-base);
  transition: border-color var(--transition-base);
}

.form-input:focus {
  outline: none;
  border-color: var(--accent-red);
}

.icon-input {
  font-size: var(--font-size-3xl);
  text-align: center;
  max-width: 100px;
}

.form-hint {
  display: block;
  margin-top: var(--spacing-xs);
  color: var(--text-light);
  font-size: var(--font-size-sm);
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  cursor: pointer;
  font-weight: 600;
}

.checkbox-label input[type="checkbox"] {
  width: 20px;
  height: 20px;
  cursor: pointer;
}

.form-actions {
  display: flex;
  gap: var(--spacing-md);
  padding-top: var(--spacing-xl);
  border-top: 2px solid var(--border);
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: stretch;
  }

  .form-actions {
    flex-direction: column;
  }
}
</style>