<template>
  <div class="admin-layout">
    <AdminNavbar />
    <div class="admin-container">
      <AdminSidebar />
      <main class="admin-main">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup>
import AdminNavbar from './AdminNavbar.vue'
import AdminSidebar from './AdminSidebar.vue'
</script>

<style scoped>
.admin-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #F5F5F7;
}

.admin-container {
  display: flex;
  flex: 1;
}

.admin-main {
  flex: 1;
  padding: var(--spacing-xl);
  overflow-y: auto;
}

@media (max-width: 768px) {
  .admin-main {
    padding: var(--spacing-md);
  }
}
</style>

