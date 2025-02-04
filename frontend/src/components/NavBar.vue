<script lang="ts" setup>
import { watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

watch(
  () => route.path,
  (newPath) => {
    console.log('Route changed to: ', newPath)
  }
)
</script>

<template>
  <div class="navBar">
    <button @click="router.push({ name: 'dashboard' })">Übersicht</button>
    <button @click="router.push({ name: 'diary' })">Tagebuch</button>
    <button>+</button>
    <button @click="router.push({ name: 'shoppingList' })">Einkaufsliste</button>
    <button @click="router.push({ name: 'settings' })">Einstellungen</button>
  </div>
</template>

<style>
.navBar {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  height: 30px;
}
</style>