/**
 * Categories Store - Category management
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { categoriesAPI } from '@/api'

export const useCategoriesStore = defineStore('categories', () => {
  // State
  const categories = ref([])
  const category = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Actions
  async function fetchCategories(params = {}) {
    try {
      loading.value = true
      error.value = null
      const { data } = await categoriesAPI.getAll(params)
      categories.value = data.data.categories || []
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch categories'
      console.error('Fetch categories error:', err)
    } finally {
      loading.value = false
    }
  }

  async function fetchCategoryBySlug(slug) {
    try {
      loading.value = true
      error.value = null
      const { data } = await categoriesAPI.getBySlug(slug)
      category.value = data.data
      return category.value
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch category'
      console.error('Fetch category error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function createCategory(categoryData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await categoriesAPI.create(categoryData)
      categories.value.unshift(data.data)
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to create category'
      console.error('Create category error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateCategory(id, categoryData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await categoriesAPI.update(id, categoryData)
      const index = categories.value.findIndex(c => c.id === id)
      if (index !== -1) {
        categories.value[index] = data.data
      }
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to update category'
      console.error('Update category error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteCategory(id) {
    try {
      loading.value = true
      error.value = null
      await categoriesAPI.delete(id)
      categories.value = categories.value.filter(c => c.id !== id)
      return true
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to delete category'
      console.error('Delete category error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    // State
    categories,
    category,
    loading,
    error,
    // Actions
    fetchCategories,
    fetchCategoryBySlug,
    createCategory,
    updateCategory,
    deleteCategory
  }
})
