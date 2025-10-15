<template>
  <component :is="layout">
    <router-view />
  </component>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// Layouts
import GuestLayout from '@/components/layout/GuestLayout.vue'
import AdminLayout from '@/components/layout/AdminLayout.vue'
import AuthLayout from '@/components/layout/AuthLayout.vue'

const route = useRoute()
const authStore = useAuthStore()

const layout = computed(() => {
  const layoutName = route.meta.layout || 'guest'
  
  switch (layoutName) {
    case 'admin':
      return AdminLayout
    case 'auth':
      return AuthLayout
    case 'guest':
    default:
      return GuestLayout
  }
})

onMounted(async () => {
  // Инициализируем auth при старте приложения
  await authStore.init()
})
</script>

<style>
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>
