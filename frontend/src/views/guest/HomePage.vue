<template>
  <div class="home-page">
    <!-- Hero Section -->
    <section class="hero">
      <div class="container">
        <div class="hero-content">
          <h1>🍔 Welcome to FastSpot!</h1>
          <p class="hero-subtitle">
            AI-powered fast food with personal recommendations
          </p>
          <div class="hero-actions">
            <router-link to="/menu" class="btn btn-lg btn-primary">
              View menu
            </router-link>
            <router-link to="/mood" class="btn btn-lg btn-secondary">
              😊 Take Mood Quiz
            </router-link>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Products -->
    <section class="featured-section">
      <div class="container">
        <h2 class="section-title">🔥 Popular dishes</h2>
        
        <div v-if="loading" class="text-center">
          <div class="spinner"></div>
        </div>

        <div v-else-if="featuredProducts.length" class="products-grid">
          <ProductCard 
            v-for="product in featuredProducts" 
            :key="product.id"
            :product="product"
          />
        </div>

        <div v-else class="text-center text-muted">
          Products not found
        </div>
      </div>
    </section>

    <!-- Promotions Section -->
    <section class="promotions-section bg-cream">
      <div class="container">
        <h2 class="section-title">🎉 Current promotions</h2>
        
        <div v-if="promotions.length" class="promotions-grid">
          <div 
            v-for="promo in promotions" 
            :key="promo.id"
            class="promo-card"
          >
            <img :src="promo.bannerImage" :alt="promo.title" />
            <div class="promo-content">
              <h3>{{ promo.title }}</h3>
              <p>{{ promo.description }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { productsAPI, promotionsAPI } from '@/api'
import ProductCard from '@/components/ui/ProductCard.vue'

const featuredProducts = ref([])
const promotions = ref([])
const loading = ref(false)

onMounted(async () => {
  await Promise.all([
    fetchFeaturedProducts(),
    fetchPromotions()
  ])
})

async function fetchFeaturedProducts() {
  try {
    loading.value = true
    const { data } = await productsAPI.getAll({ 
      limit: 6,
      tags: 'popular'
    })
    featuredProducts.value = data.data.products
  } catch (err) {
    console.error('Failed to fetch products:', err)
  } finally {
    loading.value = false
  }
}

async function fetchPromotions() {
  try {
    const { data } = await promotionsAPI.getAll({ active: true })
    promotions.value = data.data.promotions.slice(0, 3)
  } catch (err) {
    console.error('Failed to fetch promotions:', err)
  }
}
</script>

<style scoped>
.home-page {
  min-height: 100vh;
}

.hero {
  background: linear-gradient(135deg, var(--accent-red) 0%, var(--accent-yellow) 100%);
  color: white;
  padding: var(--spacing-2xl) 0;
  text-align: center;
}

.hero-content h1 {
  font-size: var(--font-size-4xl);
  margin-bottom: var(--spacing-md);
  color: white;
}

.hero-subtitle {
  font-size: var(--font-size-xl);
  margin-bottom: var(--spacing-2xl);
  opacity: 0.95;
}

.hero-actions {
  display: flex;
  gap: var(--spacing-md);
  justify-content: center;
  flex-wrap: wrap;
}

.featured-section,
.promotions-section {
  padding: var(--spacing-2xl) 0;
}

.section-title {
  text-align: center;
  margin-bottom: var(--spacing-xl);
  color: var(--contrast-black);
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: var(--spacing-lg);
}

.promotions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-lg);
}

.promo-card {
  background-color: white;
  border-radius: var(--radius-lg);
  overflow: hidden;
  box-shadow: var(--shadow-md);
  transition: transform var(--transition-base);
}

.promo-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-lg);
}

.promo-card img {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.promo-content {
  padding: var(--spacing-lg);
}

.promo-content h3 {
  margin-bottom: var(--spacing-sm);
  color: var(--accent-red);
}

@media (max-width: 768px) {
  .hero-content h1 {
    font-size: var(--font-size-3xl);
  }

  .hero-subtitle {
    font-size: var(--font-size-lg);
  }

  .products-grid {
    grid-template-columns: 1fr;
  }
}
</style>

