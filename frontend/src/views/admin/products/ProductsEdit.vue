<template>
  <div class="product-edit-page">
    <div class="page-header">
      <h1>✏️ Edit Product</h1>
      <router-link to="/admin/products" class="btn btn-ghost">
        ← Back to list
      </router-link>
    </div>

    <div v-if="loadingProduct" class="text-center py-2xl">
      <div class="spinner"></div>
      <p>Loading product...</p>
    </div>

    <form v-else-if="form.name" @submit.prevent="handleSubmit" class="product-form">
      <div class="form-grid">
        <!-- Basic Info -->
        <div class="form-section">
          <h3>Basic Information</h3>
          
          <div class="form-group">
            <label for="name">Product Name *</label>
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
            />
          </div>

          <div class="form-group">
            <label for="description">Description *</label>
            <textarea 
              id="description"
              v-model="form.description" 
              class="form-input"
              rows="4"
              required
            ></textarea>
          </div>

          <div class="form-group">
            <label for="categoryId">Category *</label>
            <select 
              id="categoryId"
              v-model="form.categoryId" 
              class="form-input"
              required
            >
              <option value="">Select category</option>
              <option 
                v-for="category in categories" 
                :key="category.id" 
                :value="category.id"
              >
                {{ category.name }}
              </option>
            </select>
          </div>
        </div>

        <!-- Pricing -->
        <div class="form-section">
          <h3>Pricing</h3>
          
          <div class="form-row">
            <div class="form-group">
              <label for="priceUSD">Price (USD) *</label>
              <input 
                id="priceUSD"
                v-model.number="form.priceUSD" 
                type="number" 
                step="0.01"
                min="0"
                class="form-input"
                required
              />
            </div>

            <div class="form-group">
              <label for="oldPriceUSD">Old Price (USD)</label>
              <input 
                id="oldPriceUSD"
                v-model.number="form.oldPriceUSD" 
                type="number" 
                step="0.01"
                min="0"
                class="form-input"
              />
            </div>
          </div>
        </div>

        <!-- Media -->
        <div class="form-section">
          <h3>Media</h3>
          
          <div class="form-group">
            <label for="image">Image URL *</label>
            <input 
              id="image"
              v-model="form.image" 
              type="url" 
              class="form-input"
              required
            />
            <div v-if="form.image" class="image-preview">
              <img :src="form.image" alt="Preview" />
            </div>
          </div>
        </div>

        <!-- Tags & Status -->
        <div class="form-section">
          <h3>Tags & Status</h3>
          
          <div class="form-group">
            <label for="tags">Tags</label>
            <input 
              id="tags"
              v-model="tagsInput" 
              type="text" 
              class="form-input"
              placeholder="popular, spicy, new"
            />
            <div v-if="form.tags.length" class="tags-preview">
              <span 
                v-for="(tag, index) in form.tags" 
                :key="index" 
                class="badge badge-primary"
              >
                {{ tag }}
              </span>
            </div>
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

        <!-- Options -->
        <div class="form-section full-width">
          <div class="section-header">
            <h3>Product Options</h3>
            <button 
              type="button" 
              @click="addOption" 
              class="btn btn-sm btn-secondary"
            >
              ➕ Add Option
            </button>
          </div>

          <div v-if="form.options.length" class="options-list">
            <div 
              v-for="(option, index) in form.options" 
              :key="index" 
              class="option-item"
            >
              <div class="form-row">
                <div class="form-group">
                  <label>Option Name</label>
                  <input 
                    v-model="option.name" 
                    type="text" 
                    class="form-input"
                  />
                </div>
                <div class="form-group">
                  <label>Type</label>
                  <select v-model="option.type" class="form-input">
                    <option value="single">Single Choice</option>
                    <option value="multiple">Multiple Choice</option>
                  </select>
                </div>
                <div class="form-group">
                  <label>Required</label>
                  <input 
                    v-model="option.required" 
                    type="checkbox"
                  />
                </div>
                <button 
                  type="button" 
                  @click="removeOption(index)" 
                  class="btn btn-sm btn-ghost"
                >
                  🗑️
                </button>
              </div>

              <div class="option-values">
                <label>Values</label>
                <input 
                  v-model="option.valuesInput" 
                  @input="updateOptionValues(option)"
                  type="text" 
                  class="form-input"
                />
              </div>
            </div>
          </div>
        </div>

        <!-- Ingredients -->
        <div class="form-section full-width">
          <div class="section-header">
            <h3>Ingredients</h3>
            <button 
              type="button" 
              @click="addIngredient" 
              class="btn btn-sm btn-secondary"
            >
              ➕ Add Ingredient
            </button>
          </div>

          <div v-if="form.ingredients.length" class="ingredients-list">
            <div 
              v-for="(ingredient, index) in form.ingredients" 
              :key="index" 
              class="ingredient-item"
            >
              <input 
                v-model="ingredient.name" 
                type="text" 
                class="form-input"
              />
              <button 
                type="button" 
                @click="removeIngredient(index)" 
                class="btn btn-sm btn-ghost"
              >
                🗑️
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Submit Buttons -->
      <div class="form-actions">
        <button 
          type="submit" 
          class="btn btn-primary btn-lg"
          :disabled="loading"
        >
          {{ loading ? 'Saving...' : '💾 Save Changes' }}
        </button>
        <router-link 
          to="/admin/products" 
          class="btn btn-ghost btn-lg"
        >
          Cancel
        </router-link>
      </div>
    </form>

    <div v-else class="text-center py-2xl">
      <p class="text-error">Product not found</p>
      <router-link to="/admin/products" class="btn btn-primary">
        Back to list
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useProductsStore } from '@/stores/products'

const router = useRouter()
const route = useRoute()
const productsStore = useProductsStore()

const loading = ref(false)
const loadingProduct = ref(true)
const categories = computed(() => productsStore.categories)

const form = ref({
  name: '',
  slug: '',
  description: '',
  categoryId: '',
  priceUSD: 0,
  oldPriceUSD: 0,
  image: '',
  tags: [],
  options: [],
  ingredients: [],
  isActive: true
})

const tagsInput = ref('')

// Parse tags from input
watch(tagsInput, (newVal) => {
  form.value.tags = newVal
    .split(',')
    .map(tag => tag.trim())
    .filter(tag => tag.length > 0)
})

onMounted(async () => {
  await Promise.all([
    productsStore.fetchCategories(),
    loadProduct()
  ])
})

async function loadProduct() {
  try {
    loadingProduct.value = true
    const productId = route.params.id
    const product = await productsStore.fetchProductById(productId)
    
    // Populate form
    form.value = {
      name: product.name || '',
      slug: product.slug || '',
      description: product.description || '',
      categoryId: product.categoryId || '',
      priceUSD: product.priceUSD || 0,
      oldPriceUSD: product.oldPriceUSD || 0,
      image: product.image || '',
      tags: product.tags || [],
      options: (product.options || []).map(opt => ({
        ...opt,
        valuesInput: opt.values?.join(', ') || ''
      })),
      ingredients: product.ingredients || [],
      isActive: product.isActive !== undefined ? product.isActive : true
    }

    tagsInput.value = form.value.tags.join(', ')
  } catch (err) {
    console.error('Load product error:', err)
    alert('Failed to load product')
  } finally {
    loadingProduct.value = false
  }
}

function addOption() {
  form.value.options.push({
    name: '',
    type: 'single',
    required: false,
    values: [],
    valuesInput: ''
  })
}

function removeOption(index) {
  form.value.options.splice(index, 1)
}

function updateOptionValues(option) {
  option.values = option.valuesInput
    .split(',')
    .map(v => v.trim())
    .filter(v => v.length > 0)
}

function addIngredient() {
  form.value.ingredients.push({ name: '' })
}

function removeIngredient(index) {
  form.value.ingredients.splice(index, 1)
}

async function handleSubmit() {
  try {
    loading.value = true

    const productData = {
      ...form.value,
      options: form.value.options.map(opt => ({
        name: opt.name,
        type: opt.type,
        required: opt.required,
        values: opt.values
      })),
      ingredients: form.value.ingredients
        .filter(i => i.name.trim())
        .map(i => ({ name: i.name.trim() }))
    }

    await productsStore.updateProduct(route.params.id, productData)
    alert('Product updated successfully! ✅')
    router.push('/admin/products')
  } catch (err) {
    console.error('Update product error:', err)
    alert('Failed to update product: ' + (err.response?.data?.error || err.message))
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.product-edit-page {
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
  padding-bottom: var(--spacing-lg);
  border-bottom: 2px solid var(--border);
}

.page-header h1 {
  margin: 0;
}

.product-form {
  max-width: 1200px;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-2xl);
}

.form-section {
  background-color: var(--bg-cream);
  padding: var(--spacing-lg);
  border-radius: var(--radius-md);
}

.form-section.full-width {
  grid-column: 1 / -1;
}

.form-section h3 {
  margin-top: 0;
  margin-bottom: var(--spacing-lg);
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

.form-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-md);
}

.image-preview {
  margin-top: var(--spacing-md);
  border-radius: var(--radius-md);
  overflow: hidden;
  max-width: 300px;
}

.image-preview img {
  width: 100%;
  height: auto;
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

.checkbox-label input[type="checkbox"] {
  width: 20px;
  height: 20px;
  cursor: pointer;
}

.options-list,
.ingredients-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.option-item,
.ingredient-item {
  background-color: white;
  padding: var(--spacing-md);
  border-radius: var(--radius-sm);
  border: 2px solid var(--border);
}

.option-values {
  margin-top: var(--spacing-sm);
}

.ingredient-item {
  display: flex;
  gap: var(--spacing-sm);
  align-items: center;
}

.form-actions {
  margin-top: var(--spacing-2xl);
  display: flex;
  gap: var(--spacing-md);
  padding-top: var(--spacing-2xl);
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
  .form-grid {
    grid-template-columns: 1fr;
  }

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