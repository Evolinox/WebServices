<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import AddComponent from './components/AddComponent.vue'
import NavBar from './components/NavBar.vue'
import WeekCalender from './components/WeekCalender.vue'
import { RouterView } from 'vue-router'
import { theme } from './themes'
import { useRoute } from 'vue-router'

// Router
const route = useRoute()

// Theme
watch(
  () => theme.value,
  (newTheme) => {
    document.documentElement.setAttribute('data-theme', newTheme)
  },
  { immediate: true },
)

onMounted(() => {
  document.documentElement.setAttribute('data-theme', theme.value)
})

// AddComponent
const showAddComponent = ref(false)
function toggleAddComponent() {
  showAddComponent.value = !showAddComponent.value
}

</script>

<template>
  <WeekCalender class="week-kalender" v-if="route.name == 'dashboard' || route.name=='diary' || route.name=='shoppingList'"/>
  <RouterView class="view"/>
  <NavBar class="nav-bar" @toggle-add-component="toggleAddComponent"/>
  <AddComponent v-if="showAddComponent" @close="toggleAddComponent" />
</template>

<style>
:root {
  --nav-bar__height: 50px;
  --week-kalender__height: 180px;
  --border-radius__secondary-background: 12px;

  
  --background-color--primary: #f8fafd;
  --background-color--secondary: #fefeff;
  --background-color--tertiary: #f1f3f5;

  --text-color--primary: #37474f;
  --text-color--secondary: #747c80;

  --accent-color--primary: #0489d7; /* Mehr */
  --accent-color--secondary: #58bdf6; /* Statistik */

  --nav__background-color: #fefeff;

  --button__text-color: #f6f6f6;
  --button__background-color: #0588d6; /* Plus */

  --icon-color: #36464e;
}

[data-theme='dark'] {
  --background-color--primary: #111926;
  --background-color--secondary: #1c2431;
  --background-color--tertiary: #323847;

  --text-color--primary: #f6f6f6;
  --text-color--secondary: #727880;

  --accent-color--primary: #0787d4; /* Mehr */
  --accent-color--secondary: #58bdf9; /* Statistik */

  --nav__background-color: #1f2b36;

  --button__text-color: #f6f6f6;
  --button__background-color: #0588d6; /* Plus */

  --icon-color: #727880;
}

body {
  background-color: var(--background-color--primary);
  color: var(--text-color--primary);
}

.week-kalender {
  height: var(--week-kalender__height);
}
.view {
  height: calc(100vh - var(--nav-bar__height) - var(--week-kalender__height));
}
.nav-bar {
  height: var(--nav-bar__height);
}
.view, .navBar {
  width: 100%;
}
.view > div, .week-kalender > div {
  width: calc(100% - 40px);
}

h1 {
  margin-block-start: 0;
  margin-block-end: 0;
}
button {
  background-color: var(--button__background-color);
  color: var(--text-color--primary);

  border: none;
  border-radius: 10px;

  cursor: pointer;
}
</style>