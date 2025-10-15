/**
 * Products Store - Product management
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { productsAPI, categoriesAPI } from '@/api'

export const useProductsStore = defineStore('products', () => {
  // State
  const products = ref([])
  const product = ref(null)
  const categories = ref([])
  const loading = ref(false)
  const error = ref(null)
  const pagination = ref({
    total: 0,
    limit: 20,
    offset: 0,
    hasMore: false
  })

  // Actions
  async function fetchProducts(params = {}) {
    try {
      loading.value = true
      error.value = null
      const { data } = await productsAPI.getAll(params)
      products.value = data.data.products
      pagination.value = data.data.pagination || pagination.value
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch products'
      console.error('Fetch products error:', err)
    } finally {
      loading.value = false
    }
  }

  async function fetchProductBySlug(slug) {
    try {
      loading.value = true
      error.value = null
      const { data } = await productsAPI.getBySlug(slug)
      product.value = data.data
      return product.value
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch product'
      console.error('Fetch product error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function fetchProductById(id) {
    try {
      loading.value = true
      error.value = null
      const { data } = await productsAPI.getById(id)
      product.value = data.data
      return product.value
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch product'
      console.error('Fetch product error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function createProduct(productData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await productsAPI.create(productData)
      // Add to list if needed
      products.value.unshift(data.data)
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to create product'
      console.error('Create product error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateProduct(id, productData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await productsAPI.update(id, productData)
      // Update in list
      const index = products.value.findIndex(p => p.id === id)
      if (index !== -1) {
        products.value[index] = data.data
      }
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to update product'
      console.error('Update product error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteProduct(id) {
    try {
      loading.value = true
      error.value = null
      await productsAPI.delete(id)
      // Remove from list
      products.value = products.value.filter(p => p.id !== id)
      return true
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to delete product'
      console.error('Delete product error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function fetchCategories(params = {}) {
    try {
      const { data } = await categoriesAPI.getAll(params)
      categories.value = data.data.categories || []
    } catch (err) {
      console.error('Fetch categories error:', err)
    }
  }

  function getCategoryById(id) {
    return categories.value.find(c => c.id === id)
  }

  return {
    // State
    products,
    product,
    categories,
    loading,
    error,
    pagination,
    // Actions
    fetchProducts,
    fetchProductBySlug,
    fetchProductById,
    createProduct,
    updateProduct,
    deleteProduct,
    fetchCategories,
    getCategoryById
  }
})

