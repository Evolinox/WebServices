<script lang="ts" setup>
import { onMounted, ref, watch, computed } from 'vue';
import { format } from "date-fns";
import croissantSvg from '../assets/croissant.svg?raw';
import pizzaSvg from '../assets/pizza.svg?raw';
import saladSvg from '../assets/salad.svg?raw';
import appleSvg from '../assets/apple.svg?raw';
import plusSvg from '../assets/plus.svg?raw';
import currentDay from '../day';
import BASE_URL from '../baseUrl';
import DiaryList from '../components/DiaryList.vue';
import AddMeal from '../components/AddToMeal.vue';

interface ProductDiary {
  ID: number;
  DailyProductsConsumedID: number;
  ProductID: number;
  ProductName: string;
  Brand: string;
  Category: string;
  WeightInGrams: number;
  Calories: number;
  ProteinsInGrams: number;
  FatsInGrams: number;
  CarbsInGrams: number;
}

interface Product {
  ID: number;
  Name: string;
  Brand: string;
  CaloriesPer100Grams: number;
  ProteinsInGrams: number;
  FatsInGrams: number;
  CarbsInGrams: number;
}

const openAddToMealBoolean = ref(false);
const meal = ref('');
const breakfastProducts = ref<ProductDiary[]>([]);
const lunchProducts = ref<ProductDiary[]>([]);
const dinnerProducts = ref<ProductDiary[]>([]);
const snackProducts = ref<ProductDiary[]>([]);
const allProducts = ref<Product[]>([]);

const currentDayBackend = computed(() => {
  return format(currentDay.value, "yyyy-MM-dd");
});

onMounted(() => {
  loadDay();
  loadProducts();
})

watch(currentDay, () => {
  loadDay()
})

function openAddToMeal(newMeal: string) {
  meal.value = newMeal;
  openAddToMealBoolean.value = true;
}
function closeAddToMeal() {
  openAddToMealBoolean.value = false;
}

function loadDay() {
  fetch(BASE_URL + '/diary/date/' + currentDayBackend.value)
  .then(response => response.json())
  .then(data => {
    breakfastProducts.value = data.products.filter((product: ProductDiary) => product.Category === 'Frühstück');
    lunchProducts.value = data.products.filter((product: ProductDiary) => product.Category === 'Mittagessen');
    dinnerProducts.value = data.products.filter((product: ProductDiary) => product.Category === 'Abendessen');
    snackProducts.value = data.products.filter((product: ProductDiary) => product.Category === 'Snack');
  })
  .catch(error => {
    console.error('Error:', error);
  })
}

function loadProducts() {
  fetch(BASE_URL + '/products/')
  .then(response => response.json())
  .then(data => {
    allProducts.value = data;
  })
  .catch(error => {
    console.error('Error:', error);
  })
}

function deleteMeal(productID: number) {
  fetch(BASE_URL + '/diary/date/'+ currentDayBackend.value + '/' + productID.toString(), {
    method: 'DELETE',
    headers: {
      'accept': 'application/json',
    },
  })
  .then(response => {
    if(response.status === 200) {
      loadDay();
    } else {
      console.error('Error:', response);
    }
  })
}
</script>

<template>
  <div class="diary-view" style="width: calc(100% - 40px); height: calc(100vh - var(--nav-bar__height) - var(--day-calender__height));">
    <div class="diary__container">
      <div class="diary-header">
        <h1>Tagebuch</h1>
      </div>
      <div class="diary__content">
        <div class="diary-content__item">
          <div class="diary-content__icon diary-content__icon--smaller" v-html="croissantSvg"></div>
          <h2>Frühstück</h2>
          <div class="diary-content__icon--add" v-html="plusSvg" @click="openAddToMeal('Frühstück')"></div>
        </div>
        <DiaryList v-if="breakfastProducts.length > 0" :list="breakfastProducts" @deleteMeal="deleteMeal"/>
        <div class="diary-content__line"></div>
        <div class="diary-content__item">
          <div class="diary-content__icon diary-content__icon--bigger" v-html="pizzaSvg"></div>
          <h2>Mittagessen</h2>
          <div class="diary-content__icon--add" v-html="plusSvg" @click="openAddToMeal('Mittagessen')"></div>
        </div>
        <DiaryList v-if="lunchProducts.length > 0" :list="lunchProducts" @deleteMeal="deleteMeal"/>
        <div class="diary-content__line"></div>
        <div class="diary-content__item">
          <div class="diary-content__icon diary-content__icon--bigger" v-html="saladSvg"></div>
          <h2>Abendessen</h2>
          <div class="diary-content__icon--add" v-html="plusSvg" @click="openAddToMeal('Abendessen')"></div>
        </div>
        <DiaryList v-if="dinnerProducts.length > 0" :list="dinnerProducts" @deleteMeal="deleteMeal"/>
        <div class="diary-content__line"></div>
        <div class="diary-content__item">
          <div class="diary-content__icon diary-content__icon--smaller" v-html="appleSvg"></div>
          <h2>Snack</h2>
          <div class="diary-content__icon--add" v-html="plusSvg" @click="openAddToMeal('Snack')"></div>
        </div>
        <DiaryList v-if="snackProducts.length > 0" :list="snackProducts" @deleteMeal="deleteMeal"/>
      </div>
    </div>
    <AddMeal v-if="openAddToMealBoolean" :meal="meal" :all-products="allProducts" :day="currentDayBackend" @close="closeAddToMeal" @loadDay="loadDay"/>
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
  cursor: pointer;
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