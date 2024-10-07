import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Users from '../views/Users.vue'
import Settings from '../views/Settings.vue'
import Reports from '../views/Reports.vue'
import Auth from '../views/Auth.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/users',
      name: 'Users',
      component: Users
    },
    {
      path: '/settings',
      name: 'Settings',
      component: Settings
    },
    {
      path: '/reports',
      name: 'Reports',
      component: Reports
    },
    {
      path: '/auth',
      name: 'Auth',
      component: Auth
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
