<script setup lang="ts">
import { watch, onMounted } from 'vue'
import NavBar from './components/NavBar.vue'
import WeekKalender from './components/WeekKalender.vue'
import { RouterView } from 'vue-router'
import { theme } from './themes'
import { useRoute } from 'vue-router'

const route = useRoute()

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
</script>

<template>
  <WeekKalender v-if="route.name !== 'settings'" class="week-kalender"/>
  <RouterView class="view"/>
  <NavBar class="nav-bar"/>
</template>

<style>
:root {
  --nav-bar__height: 50px;
  --week-kalender__height: 50px;

  --background-color: #f6f6f6;
  --text-color: #0f0f0f;

  --nav-bar__background-color: #8fbcff;
  --nav-bar__icon-color: invert(0);
}

[data-theme='dark'] {
  --background-color: #2f2f2f;
  --text-color: #f6f6f6;
  --nav-bar__background-color: #0066ff;
  --nav-bar__icon-color: invert(1);
}

body {
  background-color: var(--background-color);
  color: var(--text-color);
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
.view > div, .view, .navBar {
  width: 100%;
}


h1 {
  margin-block-start: 0;
  margin-block-end: 0;
}
button {
  background-color: var(--background-color);
  color: var(--text-color);

  border: 1px solid var(--text-color);
  border-radius: 10px;

  cursor: pointer;
}
</style>