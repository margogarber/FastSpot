/**
 * Promotions Store - Promotion management
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { promotionsAPI } from '@/api'

export const usePromotionsStore = defineStore('promotions', () => {
  // State
  const promotions = ref([])
  const promotion = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Actions
  async function fetchPromotions(params = {}) {
    try {
      loading.value = true
      error.value = null
      const { data } = await promotionsAPI.getAll(params)
      promotions.value = data.data.promotions || []
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch promotions'
      console.error('Fetch promotions error:', err)
    } finally {
      loading.value = false
    }
  }

  async function fetchPromotionById(id) {
    try {
      loading.value = true
      error.value = null
      const { data } = await promotionsAPI.getById(id)
      promotion.value = data.data
      return promotion.value
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch promotion'
      console.error('Fetch promotion error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function createPromotion(promotionData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await promotionsAPI.create(promotionData)
      promotions.value.unshift(data.data)
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to create promotion'
      console.error('Create promotion error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updatePromotion(id, promotionData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await promotionsAPI.update(id, promotionData)
      const index = promotions.value.findIndex(p => p.id === id)
      if (index !== -1) {
        promotions.value[index] = data.data
      }
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to update promotion'
      console.error('Update promotion error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deletePromotion(id) {
    try {
      loading.value = true
      error.value = null
      await promotionsAPI.delete(id)
      promotions.value = promotions.value.filter(p => p.id !== id)
      return true
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to delete promotion'
      console.error('Delete promotion error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    // State
    promotions,
    promotion,
    loading,
    error,
    // Actions
    fetchPromotions,
    fetchPromotionById,
    createPromotion,
    updatePromotion,
    deletePromotion
  }
})
