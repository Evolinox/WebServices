import { createRouter, createWebHistory } from 'vue-router'
import DashboardView from '../views/DashboardView.vue'
import DiaryView from '../views/DiaryView.vue'
import ShoppingListView from '../views/ShoppingListView.vue'
import ShoppingListAdd from '../components/ShoppingListAdd.vue'
import SettingsView from '../views/SettingsView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: DashboardView,
    },
    {
      path: '/diary',
      name: 'diary',
      component: DiaryView,
    },
    {
      path: '/shoppingList',
      name: 'shoppingList',
      component: ShoppingListView,
    },
    {
      path: '/shoppingList/add',
      name: 'shoppingList/add',
      component: ShoppingListAdd,
    },
    {
      path: '/settings',
      name: 'settings',
      component: SettingsView,
    },
  ],
})

export default router
