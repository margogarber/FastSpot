<template>
  <div class="guest-layout">
    <GuestNavbar />
    <main class="main-content">
      <slot />
    </main>
    <GuestFooter />
  </div>
</template>

<script setup>
import GuestNavbar from './GuestNavbar.vue'
import GuestFooter from './GuestFooter.vue'
</script>

<style scoped>
.guest-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  padding: var(--spacing-xl) 0;
  background-color: var(--bg-cream);
}
</style>

