<script lang="ts" setup>
import croissantSvg from '../assets/croissant.svg?raw';
import pizzaSvg from '../assets/pizza.svg?raw';
import saladSvg from '../assets/salad.svg?raw';
import appleSvg from '../assets/apple.svg?raw';
import plusSvg from '../assets/plus.svg?raw';

function showMeal(meal: String) {
  console.log('Show meal:', meal)
}

function add(meal: String) {
  console.log('Add food to', meal)
}
</script>

<template>
  <div class="diary-view" style="width: calc(100% - 40px)">
    <div class="diary__container">
      <div class="diary-header">
        <h1>Tagebuch</h1>
      </div>
      <div class="diary__content">
        <div class="diary-content__item diary-content__item--breakfast">
          <div class="diary-content__icon diary-content__icon--smaller" v-html="croissantSvg" @click="showMeal('breakfast')"></div>
          <h2 @click="showMeal('breakfast')">Frühstück</h2>
          <div class="diary-content__icon--add" v-html="plusSvg" @click="add('breakfast')"></div>
        </div>
        <div class="diary-content__line"></div>
        <div class="diary-content__item diary-content__item--lunch">
          <div class="diary-content__icon diary-content__icon--bigger" v-html="pizzaSvg" @click="showMeal('lunch')"></div>
          <h2 @click="showMeal('lunch')">Mittagessen</h2>
          <div class="diary-content__icon--add" v-html="plusSvg" @click="add('lunch')"></div>
        </div>
        <div class="diary-content__line"></div>
        <div class="diary-content__item diary-content__item--dinner">
          <div class="diary-content__icon diary-content__icon--bigger" v-html="saladSvg" @click="showMeal('dinner')"></div>
          <h2 @click="showMeal('dinner')">Abendessen</h2>
          <div class="diary-content__icon--add" v-html="plusSvg" @click="add('dinner')"></div>
        </div>
        <div class="diary-content__line"></div>
        <div class="diary-content__item diary-content__item--snacks">
          <div class="diary-content__icon diary-content__icon--smaller" v-html="appleSvg" @click="showMeal('snacks')"></div>
          <h2 @click="showMeal('snacks')">Snacks</h2>
          <div class="diary-content__icon--add" v-html="plusSvg" @click="add('snacks')"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.diary__container {
  background-color: var(--background-color--secondary);
  border-radius: var(--border-radius__secondary-background);
  padding: 20px;
  width: calc(100% - 40px);
}

.diary-header {
  margin-block-end: 1.7rem;
}

.diary-content__item {
  display: flex;
  align-items: center;
}

.diary-content__icon {
  display: flex;
}
.diary-content__item svg {
  height: 50px;
  width: 50px;
  margin-right: 20px;
}
.diary-content__icon--bigger {
  scale: 1.05;
}
.diary-content__icon--smaller {
  scale: 0.93;
}

.diary-content__icon--add {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--button__background-color);
  height: 35px;
  width: 35px;
  scale: 0.9;
  border-radius: 90px;
  margin-left: auto;
}
.diary-content__icon--add > svg {
  height: 50px;
  width: 50px;
  scale: 0.75;
  fill: var(--button__text-color);
  margin: 0;
}

.diary-content__line {
  display: flex;
  justify-content: center;
  width: 100%;
  height: 2px;
  border-radius: 1px;
  background-color: var(--icon-color);
  margin: 10px 0px;
}
</style>