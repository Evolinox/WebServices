<script lang="ts" setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import dashboardSvg from '../assets/dashboard.svg?raw'
import diarySvg from '../assets/diary.svg?raw'
import plusSvg from '../assets/plus.svg?raw'
import shoppingListSvg from '../assets/shoppingList.svg?raw'
import settingsSvg from '../assets/settings.svg?raw'

const route = useRoute()
const router = useRouter()

const isActive = (path: string) => computed(() => route.name?.toString().includes(path))
</script>

<template>
  <div class="nav-bar">
    <button :class="['dashboard', 'nav-bar__item', { 'nav-bar__item--active': isActive('dashboard').value }]" v-html="dashboardSvg" @click="router.push({ name: 'dashboard' })"></button>
    <button :class="['nav-bar__item', { 'nav-bar__item--active': isActive('diary').value }, 'rightShift']" v-html="diarySvg" @click="router.push({ name: 'diary' })"></button>
    <button :class="['plus', 'nav-bar__item']" v-html="plusSvg" @click="$emit('toggle-add-component')"></button>
    <button :class="['nav-bar__item', { 'nav-bar__item--active': isActive('shoppingList').value }]" v-html="shoppingListSvg" @click="router.push({ name: 'shoppingList' })"></button>
    <button :class="['nav-bar__item', { 'nav-bar__item--active': isActive('settings').value }]" v-html="settingsSvg" @click="router.push({ name: 'settings' })"></button>
  </div>
</template>

<style>
.nav-bar {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
}
.nav-bar__item {
  height: 45px;
  width: 45px;
  border-radius: 90px;
  border: none;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: transparent;
}
.nav-bar__item > svg {
  height: 50px;
  width: 50px;
  fill: var(--icon-color);
}
.nav-bar__item--active > svg {
  fill: var(--accent-color--primary);
}
.rightShift > svg {
  position: relative;
  left: 2px;
}
.plus {
  background-color: var(--button__background-color);
  border-radius: 90px;
}
.dashboard > svg {
  scale: 0.9;
}
.plus > svg {
  fill: var(--button__text-color);
}
</style>