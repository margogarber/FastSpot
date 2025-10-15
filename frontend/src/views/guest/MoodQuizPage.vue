<template>
  <div class="mood-quiz-page">
    <!-- Hero Section -->
    <div class="quiz-hero">
      <h1>🎯 Find Your Perfect Meal</h1>
      <p>Not sure what to order? Take our quick mood quiz and let AI recommend the perfect meal for you!</p>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="loading-container">
      <div class="spinner"></div>
      <p>Loading questions...</p>
    </div>

    <!-- Quiz Questions -->
    <div v-else-if="!showResults && questions.length > 0" class="quiz-container">
      <div class="quiz-progress">
        <div class="progress-bar">
          <div 
            class="progress-fill" 
            :style="{ width: `${progress}%` }"
          ></div>
        </div>
        <p class="progress-text">Question {{ currentQuestionIndex + 1 }} of {{ questions.length }}</p>
      </div>

      <div class="question-card">
        <h2 class="question-text">{{ currentQuestion.text }}</h2>
        
        <div class="options-grid">
          <button
            v-for="option in currentQuestion.options"
            :key="option.value"
            class="option-btn"
            :class="{ 
              'selected': isOptionSelected(option.value),
              'single': currentQuestion.type === 'single'
            }"
            @click="selectOption(option.value)"
          >
            <span class="option-label">{{ option.label }}</span>
            <span v-if="isOptionSelected(option.value)" class="check-icon">✓</span>
          </button>
        </div>

        <div class="quiz-actions">
          <button 
            v-if="currentQuestionIndex > 0"
            @click="previousQuestion"
            class="btn btn-secondary"
          >
            ← Previous
          </button>
          <button 
            v-if="currentQuestionIndex < questions.length - 1"
            @click="nextQuestion"
            :disabled="!hasAnswer"
            class="btn btn-primary"
          >
            Next →
          </button>
          <button 
            v-else
            @click="submitQuiz"
            :disabled="!hasAnswer || submitting"
            class="btn btn-primary btn-lg"
          >
            {{ submitting ? '🤖 AI is thinking...' : '✨ Get My Recommendations' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Results Section -->
    <div v-else-if="showResults" class="results-container">
      <div class="results-hero">
        <h2>🎉 Your Perfect Meal Awaits!</h2>
        <p class="ai-reasoning">{{ recommendations.reasoning }}</p>
      </div>

      <div class="recommendations-grid">
        <div 
          v-for="product in recommendedProducts"
          :key="product.slug"
          class="recommendation-card"
        >
          <div class="product-image">
            <img :src="product.image" :alt="product.name" />
            <span v-if="product.tags?.includes('popular')" class="badge badge-primary">
              🔥 POPULAR
            </span>
          </div>
          <div class="product-info">
            <h3>{{ product.name }}</h3>
            <p class="product-description">{{ product.description }}</p>
            <div class="product-footer">
              <span class="price">${{ product.priceUSD.toFixed(2) }}</span>
              <button @click="addToCart(product)" class="btn btn-primary">
                Add to Cart
              </button>
            </div>
          </div>
        </div>
      </div>

      <div class="results-actions">
        <button @click="resetQuiz" class="btn btn-secondary">
          🔄 Take Quiz Again
        </button>
        <router-link to="/menu" class="btn btn-outline">
          Browse Full Menu
        </router-link>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="error-container">
      <p class="error-message">{{ error }}</p>
      <button @click="loadQuestions" class="btn btn-primary">
        Try Again
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { moodAPI, productsAPI } from '@/api'
import { useCartStore } from '@/stores/cart'

const router = useRouter()
const cartStore = useCartStore()

// State
const loading = ref(false)
const submitting = ref(false)
const error = ref(null)
const questions = ref([])
const currentQuestionIndex = ref(0)
const answers = ref([])
const showResults = ref(false)
const recommendations = ref(null)
const recommendedProducts = ref([])

// Computed
const currentQuestion = computed(() => questions.value[currentQuestionIndex.value])
const progress = computed(() => ((currentQuestionIndex.value + 1) / questions.value.length) * 100)

const currentAnswer = computed(() => {
  return answers.value.find(a => a.questionId === currentQuestion.value?.id)
})

const hasAnswer = computed(() => {
  return currentAnswer.value && currentAnswer.value.selectedOptions.length > 0
})

// Methods
async function loadQuestions() {
  try {
    loading.value = true
    error.value = null
    const { data } = await moodAPI.getQuestions()
    questions.value = data.data.questions || []
    
    if (questions.value.length === 0) {
      error.value = 'No questions available. Please try again later.'
      return
    }

    // Initialize answers array
    answers.value = questions.value.map(q => ({
      questionId: q.id,
      selectedOptions: []
    }))
  } catch (err) {
    error.value = err.response?.data?.error || 'Failed to load questions'
    console.error('Load questions error:', err)
  } finally {
    loading.value = false
  }
}

function isOptionSelected(optionValue) {
  return currentAnswer.value?.selectedOptions.includes(optionValue)
}

function selectOption(optionValue) {
  if (!currentAnswer.value) return

  const isSingleChoice = currentQuestion.value.type === 'single'
  
  if (isSingleChoice) {
    // Single choice: replace selection
    currentAnswer.value.selectedOptions = [optionValue]
  } else {
    // Multiple choice: toggle selection
    const index = currentAnswer.value.selectedOptions.indexOf(optionValue)
    if (index > -1) {
      currentAnswer.value.selectedOptions.splice(index, 1)
    } else {
      currentAnswer.value.selectedOptions.push(optionValue)
    }
  }
}

function nextQuestion() {
  if (currentQuestionIndex.value < questions.value.length - 1) {
    currentQuestionIndex.value++
  }
}

function previousQuestion() {
  if (currentQuestionIndex.value > 0) {
    currentQuestionIndex.value--
  }
}

async function submitQuiz() {
  try {
    submitting.value = true
    error.value = null

    const { data } = await moodAPI.getRecommendations(answers.value)
    recommendations.value = data.data
    
    // Fetch recommended products
    const productSlugs = data.data.recommendations
    await loadRecommendedProducts(productSlugs)
    
    showResults.value = true
    
    // Scroll to top
    window.scrollTo({ top: 0, behavior: 'smooth' })
  } catch (err) {
    error.value = err.response?.data?.error || 'Failed to get recommendations'
    console.error('Submit quiz error:', err)
  } finally {
    submitting.value = false
  }
}

async function loadRecommendedProducts(slugs) {
  try {
    const promises = slugs.map(slug => productsAPI.getBySlug(slug))
    const results = await Promise.all(promises)
    recommendedProducts.value = results.map(r => r.data.data)
  } catch (err) {
    console.error('Load recommended products error:', err)
  }
}

async function addToCart(product) {
  try {
    await cartStore.addItem(product.id, 1)
    alert(`${product.name} added to cart!`)
  } catch (err) {
    alert('Failed to add to cart')
  }
}

function resetQuiz() {
  currentQuestionIndex.value = 0
  answers.value = questions.value.map(q => ({
    questionId: q.id,
    selectedOptions: []
  }))
  showResults.value = false
  recommendations.value = null
  recommendedProducts.value = []
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

onMounted(() => {
  loadQuestions()
})
</script>

<style scoped>
.mood-quiz-page {
  min-height: 100vh;
  padding-bottom: var(--spacing-xl);
}

.quiz-hero {
  background: linear-gradient(135deg, var(--accent-red), var(--accent-yellow));
  color: white;
  text-align: center;
  padding: var(--spacing-3xl) var(--spacing-xl);
  margin-bottom: var(--spacing-2xl);
}

.quiz-hero h1 {
  font-size: var(--font-size-3xl);
  margin-bottom: var(--spacing-md);
  text-shadow: 2px 2px 4px rgba(0,0,0,0.2);
}

.quiz-hero p {
  font-size: var(--font-size-lg);
  max-width: 600px;
  margin: 0 auto;
  opacity: 0.95;
}

.loading-container,
.error-container {
  text-align: center;
  padding: var(--spacing-3xl) var(--spacing-xl);
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid var(--bg-cream);
  border-top-color: var(--accent-red);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto var(--spacing-lg);
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.quiz-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 0 var(--spacing-lg);
}

.quiz-progress {
  margin-bottom: var(--spacing-2xl);
}

.progress-bar {
  height: 8px;
  background-color: var(--bg-cream);
  border-radius: 10px;
  overflow: hidden;
  margin-bottom: var(--spacing-sm);
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--accent-red), var(--accent-yellow));
  transition: width 0.3s ease;
}

.progress-text {
  text-align: center;
  color: var(--text-light);
  font-size: var(--font-size-sm);
}

.question-card {
  background: white;
  padding: var(--spacing-2xl);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-lg);
}

.question-text {
  font-size: var(--font-size-2xl);
  color: var(--contrast-black);
  margin-bottom: var(--spacing-2xl);
  text-align: center;
}

.options-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-2xl);
}

.option-btn {
  position: relative;
  padding: var(--spacing-xl);
  border: 2px solid var(--neutral-warm);
  border-radius: var(--radius-lg);
  background: white;
  cursor: pointer;
  transition: all var(--transition-base);
  text-align: left;
}

.option-btn:hover {
  border-color: var(--accent-red);
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.option-btn.selected {
  border-color: var(--accent-red);
  background-color: var(--bg-cream);
}

.option-label {
  display: block;
  font-size: var(--font-size-md);
  color: var(--contrast-black);
  font-weight: 500;
}

.check-icon {
  position: absolute;
  top: var(--spacing-sm);
  right: var(--spacing-sm);
  color: var(--accent-red);
  font-size: var(--font-size-xl);
  font-weight: bold;
}

.quiz-actions {
  display: flex;
  gap: var(--spacing-md);
  justify-content: space-between;
}

.quiz-actions .btn {
  flex: 1;
}

.results-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 var(--spacing-lg);
}

.results-hero {
  text-align: center;
  margin-bottom: var(--spacing-2xl);
}

.results-hero h2 {
  font-size: var(--font-size-3xl);
  color: var(--contrast-black);
  margin-bottom: var(--spacing-lg);
}

.ai-reasoning {
  font-size: var(--font-size-lg);
  color: var(--text-light);
  max-width: 700px;
  margin: 0 auto;
  line-height: 1.6;
  padding: var(--spacing-lg);
  background: var(--bg-cream);
  border-radius: var(--radius-lg);
  border-left: 4px solid var(--accent-yellow);
}

.recommendations-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-xl);
  margin-bottom: var(--spacing-2xl);
}

.recommendation-card {
  background: white;
  border-radius: var(--radius-xl);
  overflow: hidden;
  box-shadow: var(--shadow-lg);
  transition: all var(--transition-base);
}

.recommendation-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-xl);
}

.product-image {
  position: relative;
  height: 250px;
  overflow: hidden;
}

.product-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.product-info {
  padding: var(--spacing-xl);
}

.product-info h3 {
  font-size: var(--font-size-xl);
  color: var(--contrast-black);
  margin-bottom: var(--spacing-sm);
}

.product-description {
  color: var(--text-light);
  margin-bottom: var(--spacing-lg);
  line-height: 1.5;
}

.product-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price {
  font-size: var(--font-size-2xl);
  font-weight: 700;
  color: var(--accent-red);
}

.results-actions {
  display: flex;
  gap: var(--spacing-lg);
  justify-content: center;
  margin-top: var(--spacing-2xl);
}

.error-message {
  color: var(--error-color, #dc2626);
  font-size: var(--font-size-lg);
  margin-bottom: var(--spacing-lg);
}

@media (max-width: 768px) {
  .quiz-hero h1 {
    font-size: var(--font-size-2xl);
  }

  .options-grid {
    grid-template-columns: 1fr;
  }

  .quiz-actions {
    flex-direction: column;
  }

  .results-actions {
    flex-direction: column;
  }

  .recommendations-grid {
    grid-template-columns: 1fr;
  }
}
</style>
