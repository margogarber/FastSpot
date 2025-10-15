<template>
  <aside class="admin-sidebar">
    <nav class="sidebar-nav">
      <router-link to="/admin" class="sidebar-link" exact-active-class="active">
        📊 Dashboard
      </router-link>

      <div class="sidebar-section">
        <div class="sidebar-section-title">Menu</div>
        <router-link to="/admin/categories" class="sidebar-link">
          📂 Categories
        </router-link>
        <router-link to="/admin/products" class="sidebar-link">
          🍔 Products
        </router-link>
        <router-link to="/admin/promotions" class="sidebar-link">
          🎉 Promotions
        </router-link>
      </div>

      <div class="sidebar-section">
        <div class="sidebar-section-title">Orders</div>
        <router-link to="/admin/orders" class="sidebar-link">
          📦 All orders
        </router-link>
      </div>

      <div class="sidebar-section">
        <div class="sidebar-section-title">AI & Mood Quiz</div>
        <router-link to="/admin/mood-questions" class="sidebar-link">
          🤖 Mood Quiz (AI)
        </router-link>
      </div>

      <div class="sidebar-section">
        <router-link to="/" class="sidebar-link">
          🌐 Back to site
        </router-link>
      </div>
    </nav>
  </aside>
</template>

<style scoped>
.admin-sidebar {
  width: 260px;
  background-color: white;
  box-shadow: var(--shadow-sm);
  overflow-y: auto;
}

.sidebar-nav {
  padding: var(--spacing-lg);
}

.sidebar-section {
  margin-bottom: var(--spacing-lg);
}

.sidebar-section-title {
  font-size: var(--font-size-xs);
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-light);
  margin-bottom: var(--spacing-sm);
  padding: 0 var(--spacing-md);
}

.sidebar-link {
  display: block;
  padding: var(--spacing-sm) var(--spacing-md);
  color: var(--text);
  text-decoration: none;
  border-radius: var(--radius-md);
  font-weight: 500;
  transition: all var(--transition-fast);
  margin-bottom: var(--spacing-xs);
}

.sidebar-link:hover {
  background-color: var(--bg-cream);
  color: var(--accent-red);
}

.sidebar-link.active {
  background-color: var(--accent-red);
  color: white;
}

@media (max-width: 768px) {
  .admin-sidebar {
    width: 200px;
  }

  .sidebar-link {
    font-size: var(--font-size-sm);
  }
}
</style>

