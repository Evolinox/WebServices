<script setup lang="ts">
import { watch, onMounted } from 'vue'
import NavBar from './components/NavBar.vue'
import { RouterView } from 'vue-router'
import { theme } from './themes'

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
  <RouterView class="view"/>
  <NavBar class="navBar"/>
</template>

<style>
:root {
  --background-color: #f6f6f6;
  --text-color: #0f0f0f;

  --navBar__height: 50px;
}

[data-theme='dark'] {
  --background-color: #2f2f2f;
  --text-color: #f6f6f6;
}

body {
  background-color: var(--background-color);
  color: var(--text-color);
}

.view {
  height: calc(100vh - var(--navBar__height));
  overflow-y: auto;
}
.navBar {
  height: var(--navBar__height);
}
.view > div, .view, .navBar {
  width: 100%;
  display: flex;
}

button {
  background-color: var(--background-color);
  color: var(--text-color);

  border: 1px solid var(--text-color);
  border-radius: 10px;

  cursor: pointer;
}
</style>