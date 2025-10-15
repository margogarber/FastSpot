/**
 * Mood Store - Mood Quiz management
 * AI analyzes quiz answers and recommends products on its own
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { moodAPI } from '@/api'

export const useMoodStore = defineStore('mood', () => {
  // State
  const questions = ref([])
  const question = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // ==================== QUESTIONS ====================
  async function fetchQuestions() {
    try {
      loading.value = true
      error.value = null
      const { data } = await moodAPI.getQuestions()
      questions.value = data.data.questions || []
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch questions'
      console.error('Fetch questions error:', err)
    } finally {
      loading.value = false
    }
  }

  async function fetchAllQuestions() {
    try {
      loading.value = true
      error.value = null
      const { data } = await moodAPI.getAllQuestions()
      questions.value = data.data.questions || []
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch questions'
      console.error('Fetch all questions error:', err)
    } finally {
      loading.value = false
    }
  }

  async function fetchQuestionById(id) {
    try {
      loading.value = true
      error.value = null
      const { data } = await moodAPI.getQuestionById(id)
      question.value = data.data
      return question.value
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to fetch question'
      console.error('Fetch question error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function createQuestion(questionData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await moodAPI.createQuestion(questionData)
      questions.value.unshift(data.data)
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to create question'
      console.error('Create question error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateQuestion(id, questionData) {
    try {
      loading.value = true
      error.value = null
      const { data } = await moodAPI.updateQuestion(id, questionData)
      const index = questions.value.findIndex(q => q.id === id)
      if (index !== -1) {
        questions.value[index] = data.data
      }
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to update question'
      console.error('Update question error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteQuestion(id) {
    try {
      loading.value = true
      error.value = null
      await moodAPI.deleteQuestion(id)
      questions.value = questions.value.filter(q => q.id !== id)
      return true
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to delete question'
      console.error('Delete question error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  // ==================== AI RECOMMENDATIONS ====================
  async function getRecommendations(answers) {
    try {
      loading.value = true
      error.value = null
      const { data } = await moodAPI.getRecommendations(answers)
      return data.data
    } catch (err) {
      error.value = err.response?.data?.error?.message || 'Failed to get recommendations'
      console.error('Get recommendations error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    // State
    questions,
    question,
    loading,
    error,
    // Actions - Questions
    fetchQuestions,
    fetchAllQuestions,
    fetchQuestionById,
    createQuestion,
    updateQuestion,
    deleteQuestion,
    // Actions - AI Recommendations
    getRecommendations
  }
})
