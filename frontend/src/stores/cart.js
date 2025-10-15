/**
 * Cart Store - Cart management
 */
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { cartAPI } from '@/api'

export const useCartStore = defineStore('cart', () => {
  // State
  const cart = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Getters
  const items = computed(() => cart.value?.items || [])
  const totalUSD = computed(() => cart.value?.totalUSD || 0)
  const itemsCount = computed(() => items.value.reduce((sum, item) => sum + item.qty, 0))
  const isEmpty = computed(() => items.value.length === 0)

  // Actions
  async function fetchCart() {
    try {
      loading.value = true
      error.value = null
      const { data } = await cartAPI.get()
      cart.value = data.data || data
    } catch (err) {
      // Ignore 401 errors (user not authenticated)
      if (err.response?.status === 401) {
        cart.value = { items: [], totalUSD: 0, currency: 'USD' }
        return
      }
      error.value = err.response?.data?.error || 'Failed to fetch cart'
      console.error('Fetch cart error:', err)
    } finally {
      loading.value = false
    }
  }

  async function addItem(productId, qty = 1, chosenIngredients = [], chosenOptions = {}) {
    try {
      loading.value = true
      error.value = null
      
      const { data } = await cartAPI.addItem({
        productId,
        qty,
        chosenIngredients,
        chosenOptions
      })
      
      cart.value = data.data
      return true
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to add item'
      console.error('Add item error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateItem(productId, qty, chosenIngredients = [], chosenOptions = {}) {
    try {
      loading.value = true
      error.value = null
      
      const { data } = await cartAPI.updateItem(productId, {
        qty,
        chosenIngredients,
        chosenOptions
      })
      
      cart.value = data.data
      return true
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to update item'
      console.error('Update item error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function removeItem(productId) {
    try {
      loading.value = true
      error.value = null
      
      const { data } = await cartAPI.removeItem(productId)
      cart.value = data.data
      return true
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to remove item'
      console.error('Remove item error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function clearCart() {
    try {
      loading.value = true
      error.value = null
      
      await cartAPI.clear()
      cart.value = {
        items: [],
        totalUSD: 0,
        currency: 'USD'
      }
      return true
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to clear cart'
      console.error('Clear cart error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  function getItemByProductId(productId) {
    return items.value.find(item => item.productId === productId)
  }

  return {
    // State
    cart,
    loading,
    error,
    // Getters
    items,
    totalUSD,
    itemsCount,
    isEmpty,
    // Actions
    fetchCart,
    addItem,
    updateItem,
    removeItem,
    clearCart,
    getItemByProductId
  }
})

