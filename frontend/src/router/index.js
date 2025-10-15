/**
 * Vue Router configuration
 */
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// Guest pages
const HomePage = () => import('@/views/guest/HomePage.vue')
const MenuPage = () => import('@/views/guest/MenuPage.vue')
const ProductPage = () => import('@/views/guest/ProductPage.vue')
const PromotionsPage = () => import('@/views/guest/PromotionsPage.vue')
const MoodQuizPage = () => import('@/views/guest/MoodQuizPage.vue')
const CartPage = () => import('@/views/guest/CartPage.vue')
const CheckoutPage = () => import('@/views/guest/CheckoutPage.vue')
const OrderTrackPage = () => import('@/views/guest/OrderTrackPage.vue')

// Admin pages
const AdminLogin = () => import('@/views/admin/LoginPage.vue')
const AdminDashboard = () => import('@/views/admin/DashboardPage.vue')

// Admin - Products
const ProductsList = () => import('@/views/admin/products/ProductsList.vue')
const ProductsCreate = () => import('@/views/admin/products/ProductsCreate.vue')
const ProductsEdit = () => import('@/views/admin/products/ProductsEdit.vue')
const ProductsView = () => import('@/views/admin/products/ProductsView.vue')

// Admin - Categories
const CategoriesList = () => import('@/views/admin/categories/categorieList.vue')
const CategoriesCreate = () => import('@/views/admin/categories/categorieCreate.vue')
const CategoriesEdit = () => import('@/views/admin/categories/categorieEdit.vue')
const CategoriesView = () => import('@/views/admin/categories/categorieView.vue')

// Admin - Promotions
const PromotionsList = () => import('@/views/admin/promotions/promotionList.vue')
const PromotionsCreate = () => import('@/views/admin/promotions/promotionCreate.vue')
const PromotionsEdit = () => import('@/views/admin/promotions/promotionEdit.vue')
const PromotionsView = () => import('@/views/admin/promotions/promotionView.vue')

// Admin - Orders
const OrdersList = () => import('@/views/admin/orders/orderList.vue')
const OrdersView = () => import('@/views/admin/orders/orderView.vue')

// Admin - Mood Questions (AI analyzes answers and recommends products)
const MoodQuestionsList = () => import('@/views/admin/mood-questions/mood-questionList.vue')
const MoodQuestionsCreate = () => import('@/views/admin/mood-questions/mood-questionCreate.vue')
const MoodQuestionsEdit = () => import('@/views/admin/mood-questions/mood-questionEdit.vue')
const MoodQuestionsView = () => import('@/views/admin/mood-questions/mood-questionView.vue')

const routes = [
  // Guest routes
  {
    path: '/',
    name: 'home',
    component: HomePage,
    meta: { layout: 'guest' }
  },
  {
    path: '/menu',
    name: 'menu',
    component: MenuPage,
    meta: { layout: 'guest' }
  },
  {
    path: '/products/:slug',
    name: 'product',
    component: ProductPage,
    meta: { layout: 'guest' }
  },
  {
    path: '/promotions',
    name: 'promotions',
    component: PromotionsPage,
    meta: { layout: 'guest' }
  },
  {
    path: '/mood',
    name: 'mood',
    component: MoodQuizPage,
    meta: { layout: 'guest' }
  },
  {
    path: '/cart',
    name: 'cart',
    component: CartPage,
    meta: { layout: 'guest' }
  },
  {
    path: '/checkout',
    name: 'checkout',
    component: CheckoutPage,
    meta: { layout: 'guest' }
  },
  {
    path: '/orders/:id/track',
    name: 'order-track',
    component: OrderTrackPage,
    meta: { layout: 'guest' }
  },

  // Admin routes
  {
    path: '/admin/login',
    name: 'admin-login',
    component: AdminLogin,
    meta: { layout: 'auth', guestOnly: true }
  },
  {
    path: '/admin',
    name: 'admin-dashboard',
    component: AdminDashboard,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },

  // Admin - Products
  {
    path: '/admin/products',
    name: 'admin-products',
    component: ProductsList,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/products/create',
    name: 'admin-products-create',
    component: ProductsCreate,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/products/:id',
    name: 'admin-products-view',
    component: ProductsView,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/products/:id/edit',
    name: 'admin-products-edit',
    component: ProductsEdit,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },

  // Admin - Categories
  {
    path: '/admin/categories',
    name: 'admin-categories',
    component: CategoriesList,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/categories/create',
    name: 'admin-categories-create',
    component: CategoriesCreate,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/categories/:id',
    name: 'admin-categories-view',
    component: CategoriesView,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/categories/:id/edit',
    name: 'admin-categories-edit',
    component: CategoriesEdit,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },

  // Admin - Promotions
  {
    path: '/admin/promotions',
    name: 'admin-promotions',
    component: PromotionsList,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/promotions/create',
    name: 'admin-promotions-create',
    component: PromotionsCreate,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/promotions/:id',
    name: 'admin-promotions-view',
    component: PromotionsView,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/promotions/:id/edit',
    name: 'admin-promotions-edit',
    component: PromotionsEdit,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },

  // Admin - Orders
  {
    path: '/admin/orders',
    name: 'admin-orders',
    component: OrdersList,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/orders/:id',
    name: 'admin-orders-view',
    component: OrdersView,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },

  // Admin - Mood Questions
  {
    path: '/admin/mood-questions',
    name: 'admin-mood-questions',
    component: MoodQuestionsList,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/mood-questions/create',
    name: 'admin-mood-questions-create',
    component: MoodQuestionsCreate,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/mood-questions/:id',
    name: 'admin-mood-questions-view',
    component: MoodQuestionsView,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },
  {
    path: '/admin/mood-questions/:id/edit',
    name: 'admin-mood-questions-edit',
    component: MoodQuestionsEdit,
    meta: { layout: 'admin', requiresAuth: true, requiresAdmin: true }
  },

]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// Navigation guards
router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()

  // Initialize auth if needed
  if (!authStore.token && !authStore.sessionId) {
    await authStore.init()
  }

  // Check route requirements
  if (to.meta.requiresAuth) {
    if (!authStore.isAuthenticated) {
      return next({ name: 'admin-login', query: { redirect: to.fullPath } })
    }
  }

  if (to.meta.requiresAdmin) {
    if (!authStore.isAdmin) {
      return next({ name: 'home' })
    }
  }

  if (to.meta.guestOnly) {
    if (authStore.isAdmin) {
      return next({ name: 'admin-dashboard' })
    }
  }

  next()
})

export default router

