<template>
  <footer class="footer">
    <div class="container">
      <div class="footer-content">
        <div class="footer-section">
          <h3>🍔 FastSpot</h3>
          <p>AI-powered fast food with mood-based recommendations</p>
        </div>

        <div class="footer-section">
          <h4>Navigation</h4>
          <ul class="footer-links">
            <li><router-link to="/">Home</router-link></li>
            <li><router-link to="/menu">Menu</router-link></li>
            <li><router-link to="/promotions">Promotions</router-link></li>
            <li><router-link to="/mood">Mood Quiz</router-link></li>
          </ul>
        </div>

        <div class="footer-section">
          <h4>Contacts</h4>
          <ul class="footer-links">
            <li>📞 +1 (234) 567-8900</li>
            <li>📧 info@fastspot.com</li>
            <li>📍 123 Main St, City</li>
          </ul>
        </div>

        <div class="footer-section">
          <h4>Working hours</h4>
          <p>Mon-Sun: 10:00 - 22:00</p>
        </div>
      </div>

      <div class="footer-bottom">
        <p>&copy; {{ currentYear }} FastSpot. All rights reserved. 🍔</p>
        <p class="text-muted">Made with ❤️ and 🍔</p>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>

<style scoped>
.footer {
  background-color: var(--contrast-black);
  color: white;
  padding: var(--spacing-2xl) 0 var(--spacing-lg);
  margin-top: auto;
}

.footer-content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-xl);
  margin-bottom: var(--spacing-xl);
}

.footer-section h3,
.footer-section h4 {
  color: var(--accent-yellow);
  margin-bottom: var(--spacing-md);
}

.footer-section p {
  color: rgba(255, 255, 255, 0.8);
  margin-bottom: var(--spacing-sm);
}

.footer-links {
  list-style: none;
  padding: 0;
}

.footer-links li {
  margin-bottom: var(--spacing-sm);
}

.footer-links a {
  color: rgba(255, 255, 255, 0.8);
  text-decoration: none;
  transition: color var(--transition-fast);
}

.footer-links a:hover {
  color: var(--accent-yellow);
}

.footer-bottom {
  text-align: center;
  padding-top: var(--spacing-lg);
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.6);
}

@media (max-width: 768px) {
  .footer-content {
    grid-template-columns: 1fr;
    text-align: center;
  }
}
</style>

