<template>
  <div class="auth-layout">
    <div class="auth-container">
      <div class="auth-logo">
        <router-link to="/">🍔 FastSpot</router-link>
      </div>
      <div class="auth-content">
        <slot />
      </div>
    </div>
  </div>
</template>

<style scoped>
.auth-layout {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, var(--accent-red) 0%, var(--accent-yellow) 100%);
  padding: var(--spacing-lg);
}

.auth-container {
  width: 100%;
  max-width: 450px;
}

.auth-logo {
  text-align: center;
  margin-bottom: var(--spacing-2xl);
}

.auth-logo a {
  font-size: var(--font-size-4xl);
  font-weight: 700;
  color: white;
  text-decoration: none;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
  transition: transform var(--transition-fast);
  display: inline-block;
}

.auth-logo a:hover {
  transform: scale(1.05);
}

.auth-content {
  background-color: white;
  border-radius: var(--radius-xl);
  padding: var(--spacing-2xl);
  box-shadow: var(--shadow-xl);
}
</style>

