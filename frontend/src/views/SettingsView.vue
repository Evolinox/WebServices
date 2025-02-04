<script lang="ts" setup>
import { theme, toggleTheme } from '../themes'
</script>

<template>
  <div class="settingsView">
    <h1>Einstellungen</h1>
    <button class="themeSwitch" @click="toggleTheme">
      {{ theme === 'light' ? 'Switch to Dark Mode' : 'Switch to Light Mode' }}
    </button>
  </div>
</template>

<style>
.settingsView {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.themeSwitch {
  height: 20px;
  border: 1px solid black;
  border-radius: 10px;
}
</style>