import { createRouter, createWebHistory } from 'vue-router'
import DashboardView from '../views/DashboardView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/dashboard'
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: DashboardView,
      meta: {
        title: 'Dashboard'
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  document.title = to.name ? to.name.toString() : 'Ledger';
  next();
});

export default router
