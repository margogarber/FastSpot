<template>
  <div class="category-edit-page">
    <div class="page-header">
      <h1>✏️ Edit Category</h1>
      <router-link to="/admin/categories" class="btn btn-ghost">
        ← Back to list
      </router-link>
    </div>

    <div v-if="loadingCategory" class="text-center py-2xl">
      <div class="spinner"></div>
      <p>Loading category...</p>
    </div>

    <form v-else-if="form.name" @submit.prevent="handleSubmit" class="category-form">
      <div class="form-section">
        <h3>Category Information</h3>
        
        <div class="form-group">
          <label for="name">Name *</label>
          <input 
            id="name"
            v-model="form.name" 
            type="text" 
            class="form-input"
            required
          />
        </div>

        <div class="form-group">
          <label for="slug">Slug *</label>
          <input 
            id="slug"
            v-model="form.slug" 
            type="text" 
            class="form-input"
            required
            pattern="[a-z0-9-]+"
          />
        </div>

        <div class="form-group">
          <label for="icon">Icon (emoji) *</label>
          <input 
            id="icon"
            v-model="form.icon" 
            type="text" 
            class="form-input icon-input"
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
          {{ loading ? 'Saving...' : '💾 Save Changes' }}
        </button>
        <router-link 
          to="/admin/categories" 
          class="btn btn-ghost btn-lg"
        >
          Cancel
        </router-link>
      </div>
    </form>

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
import { useRouter, useRoute } from 'vue-router'
import { useCategoriesStore } from '@/stores/categories'
import { categoriesAPI } from '@/api'

const router = useRouter()
const route = useRoute()
const categoriesStore = useCategoriesStore()

const loading = ref(false)
const loadingCategory = ref(true)

const form = ref({
  name: '',
  slug: '',
  icon: '',
  description: '',
  isActive: true
})

onMounted(async () => {
  await loadCategory()
})

async function loadCategory() {
  try {
    loadingCategory.value = true
    const categoryId = route.params.id
    const { data } = await categoriesAPI.getById ? 
      await categoriesAPI.getById(categoryId) : 
      await categoriesAPI.getBySlug(categoryId)
    
    const category = data.data
    form.value = {
      name: category.name || '',
      slug: category.slug || '',
      icon: category.icon || '',
      description: category.description || '',
      isActive: category.isActive !== undefined ? category.isActive : true
    }
  } catch (err) {
    console.error('Load category error:', err)
    alert('Failed to load category')
  } finally {
    loadingCategory.value = false
  }
}

async function handleSubmit() {
  try {
    loading.value = true
    await categoriesStore.updateCategory(route.params.id, form.value)
    alert('Category updated successfully! ✅')
    router.push('/admin/categories')
  } catch (err) {
    console.error('Update category error:', err)
    alert('Failed to update category: ' + (err.response?.data?.error || err.message))
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.category-edit-page {
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

  .form-actions {
    flex-direction: column;
  }
}
</style>