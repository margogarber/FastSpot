/**
 * Orders Store - Управление заказами
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { ordersAPI } from '@/api'

export const useOrdersStore = defineStore('orders', () => {
  // State
  const orders = ref([])
  const order = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Actions
  async function fetchOrders(params = {}) {
    try {
      loading.value = true
      error.value = null
      const { data } = await ordersAPI.getAll(params)
      orders.value = data.data.orders
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch orders'
      console.error('Fetch orders error:', err)
    } finally {
      loading.value = false
    }
  }

  async function fetchOrderById(id) {
    try {
      loading.value = true
      error.value = null
      const { data } = await ordersAPI.getById(id)
      order.value = data.data
      return order.value
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch order'
      console.error('Fetch order error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function createOrder(orderData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await ordersAPI.create(orderData)
      order.value = data.data
      return order.value
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to create order'
      console.error('Create order error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  // Admin methods
  async function fetchAllOrdersAdmin(params = {}) {
    try {
      loading.value = true
      error.value = null
      const { data } = await ordersAPI.getAllAdmin(params)
      orders.value = data.data.orders
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch orders'
      console.error('Fetch orders error:', err)
    } finally {
      loading.value = false
    }
  }

  async function updateOrderStatus(id, status, note = '') {
    try {
      loading.value = true
      error.value = null
      const { data } = await ordersAPI.updateStatus(id, { status, note })
      
      // Обновляем в списке
      const index = orders.value.findIndex(o => o.id === id)
      if (index !== -1) {
        orders.value[index] = data.data
      }
      
      // Обновляем текущий заказ
      if (order.value?.id === id) {
        order.value = data.data
      }
      
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to update order'
      console.error('Update order error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  function getOrdersByStatus(status) {
    return orders.value.filter(o => o.status === status)
  }

  return {
    // State
    orders,
    order,
    loading,
    error,
    // Actions
    fetchOrders,
    fetchOrderById,
    createOrder,
    fetchAllOrdersAdmin,
    updateOrderStatus,
    getOrdersByStatus
  }
})

