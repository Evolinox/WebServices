<script lang="ts" setup>
import { ref } from 'vue';
import plusSvg from '../assets/plus.svg?raw';
const props = defineProps<{ meal: string, addArticle: boolean}>();
const addArticleBoolean = ref(props.addArticle);

function addArticle(event: Event) {
    const parentElement = (event.target as HTMLElement).parentElement;
  const article = (parentElement?.querySelector('#article') as HTMLInputElement).value;
  const quantity = (parentElement?.querySelector('#quantity') as HTMLInputElement).value;
  console.log('Add article:', article, 'Quantity:', quantity);
}
</script>

<template>
  <div class="show-meal__overlay" @click.self="$emit('close')" style="width: 100%;">
    <div class="show-meal__card">
      <div class="show-meal__card-header">
        <h2>{{ props.meal }}</h2>
        <div class="show-meal__close-button" v-html="plusSvg" @click="$emit('close')"></div>
      </div>
      <div class="show-meal__card-content">
        <!-- Hier können Sie den Inhalt der Mahlzeit hinzufügen -->
        <p>Details zur Mahlzeit: {{ props.meal }}</p>
      </div>
      <div v-if="addArticleBoolean" class="show-meal__add-article">
        <input type="text" placeholder="Artikel" id="article">
        <input type="text" placeholder="Menge" id="quantity">
        <input type="submit" @click="addArticle($event)" value="Hinzufügen" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.show-meal__overlay {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1;
}

.show-meal__card {
  display: flex;
  flex-direction: column;
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  min-width: 370px;
  background-color: var(--background-color--secondary);
  padding: 1em;
  border-radius: 8px;
}

.show-meal__card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.show-meal__card-header h2 {
  margin-left: 5px;
}

.show-meal__close-button {
  width: 20px;
  height: 20px;
  cursor: pointer;
  margin-left: 30px;
  margin-top: 0px;
  scale: 1.2;
  fill: var(--icon-color);
}

.show-meal__card-content {
  padding: 0px 35px;
  text-align: center;
}
</style>