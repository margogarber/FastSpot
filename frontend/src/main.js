import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import pinia from './stores'

// Import global styles
import './assets/theme.css'

const app = createApp(App)

app.use(pinia)
app.use(router)

app.mount('#app')
