<script lang="ts" setup>
import plusSvg from '../assets/plus.svg?raw'
</script>

<template>
  <div class="add-component__overlay" @click.self="$emit('close')">
    <div class="add-component__content">
      <div class="add-component__close-button" v-html="plusSvg" @click="$emit('close')"></div>
      <ul>
        <li>Lebensmittel</li>
        <li>Kalendereintrag</li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.add-component__overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
}

.add-component__content {
  display: flex;
  flex-direction: column;
  position: fixed;
  bottom: calc(var(--nav-bar__height) + 0.5em);
  left: 50%;
  transform: translateX(-50%);
  background-color: var(--background-color--secondary);
  padding: 1em;
  border-radius: var(--border-radius__secondary-background);
}
ul {
  display: flex;
  flex-direction: column;
  align-items: center;
  list-style-type: none;
  padding: 0;
}
li {
  padding: 0.5em;
  cursor: pointer;
}
@media (hover: hover) {
  li:hover {
    color: var(--accent-color--primary);
  }
}
.add-component__close-button, .add-component__close-button > svg {
  width: 20px;
  height: 20px;
  margin-left: auto;
  margin-right: 0;
  rotate: 45deg;
  cursor: pointer;
  fill: var(--icon-color);
}
</style>