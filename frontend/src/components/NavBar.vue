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
  <div class="nav-bar">
    <button class="nav-bar__item" @click="router.push({ name: 'dashboard' })"><img src="../assets/dashboard.svg" alt=""></button>
    <button class="nav-bar__item" @click="router.push({ name: 'diary' })"><img src="../assets/diary.svg" alt="" style="
    position: relative;
    left: 2px;"
></button>
    <button class="nav-bar__item"><img src="../assets/plus.svg" alt=""></button>
    <button class="nav-bar__item" @click="router.push({ name: 'shoppingList' })"><img src="../assets/shoppingList.svg" alt=""></button>
    <button class="nav-bar__item" @click="router.push({ name: 'settings' })"><img src="../assets/settings.svg" alt=""></button>
  </div>
</template>

<style>
.nav-bar {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}
.nav-bar__item {
  height: 45px;
  width: 45px;
  border-radius: 90px;
  display: flex;
  justify-content: center;
  align-items: center;
}
img {
  height: 25px;
  width: auto;
  filter: var(--nav-bar__icon-color);
}
</style>