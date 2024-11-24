import { createRouter, createWebHistory } from 'vue-router';
import Config from '../components/Config.vue';
import WhatIsThis from '../components/WhatIsThis.vue';
import APIConfig from '../components/APIConfig.vue';

const routes = [
    {
        path: '/',
        name: 'WhatIsThis',
        component: WhatIsThis,
    },
    {
        path: '/config',
        name: 'Config',
        component: Config,
    },
    {
        path: '/apiConfig',
        name: 'APIConfig',
        component: APIConfig,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
