import { createApp } from 'vue'
import App from './App.vue'
import Router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

const app = createApp(App)
app
    .use(ElementPlus)
    .use(Router)

app.mount('#app')
