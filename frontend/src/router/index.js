import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    {
        path: '/',
        name: 'WhatIsThis',
        component: () => import('../components/WhatIsThis.vue')
    },
    {
        path: '/config',
        name: 'Config',
        component: () => import('../components/Config.vue')
    },
    {
        path: '/apiConfig',
        name: 'APIConfig',
        component: () => import('../components/APIConfig.vue')
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
