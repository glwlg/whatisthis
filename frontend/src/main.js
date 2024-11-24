import { createApp } from 'vue'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'; // 引入mdi字体样式

import App from './App.vue'
import './style.css'

// 引入路由实例
import router from './router'

const vuetify = createVuetify({
    components,
    directives,
})

createApp(App).use(vuetify).use(router).mount('#app')
