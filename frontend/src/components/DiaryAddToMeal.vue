<script lang="ts" setup>
import { ref } from 'vue';
import plusSvg from '../assets/plus.svg?raw';
import BASE_URL from '../baseUrl';

interface Product {
  ID: number;
  Name: string;
  Brand: string;
  CaloriesPer100Grams: number;
  ProteinsInGrams: number;
  FatsInGrams: number;
  CarbsInGrams: number;
}

const props = defineProps<{ meal: string, allProducts: Product[], day: string }>();
const products = ref<Product[]>(props.allProducts);
const selectedMeal = ref(props.meal);
const emit = defineEmits(['close', 'loadDay']);

function addProduct(id: number) {
  if (!weight.value) {
    console.error('No weight given');
    return;
  }
  const addedProduct: Product = products.value[id];
  fetch(BASE_URL+'/consume/', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      Product: {
        ID: addedProduct.ID,
        Name: addedProduct.Name,
        Brand: addedProduct.Brand,
        CaloriesPer100Grams: addedProduct.CaloriesPer100Grams,
        ProteinsInGrams: addedProduct.ProteinsInGrams,
        FatsInGrams: addedProduct.FatsInGrams,
        CarbsInGrams: addedProduct.CarbsInGrams,
      },
      Date: props.day,
      Weight: weight.value,
      Category: selectedMeal.value,
    }),
  })
  .then(response => {
    if(response.status === 201) {
      emit('loadDay');
      emit('close');
    } else {
      console.error('Error:', response);
    }
  })
}

const weight = ref<number | null>(null);
const handleWeightInput = (event: Event) => {
  const input = event.target as HTMLInputElement
  input.value = input.value.replace(/[^0-9]/g, "");
  if (input.value.length > 5) {
    input.value = input.value.slice(0, 5);
  }
  weight.value = parseInt(input.value);
}
</script>

<template>
  <div class="add-meal__overlay" @click.self="$emit('close')" style="width: 100%;">
    <div class="add-meal__card">
      <div class="add-meal__card-header">
        <h2>Füge Produkt hinzu</h2>
        <div class="add-meal__close-button" v-html="plusSvg" @click="$emit('close')"></div>
      </div>
      <div class="add-meal__card-content">
        <select name="meal-select" v-model="selectedMeal">
          <option value="Frühstück">Frühstück</option>
          <option value="Mittagessen">Mittagessen</option>
          <option value="Abendessen">Abendessen</option>
          <option value="Snack">Snack</option>
        </select>
        <div class="add-meal__product-list">
          <div v-for="(product, index) in products" class="add-meal__product">
            <div class="product__description">
              <p>{{ product.Name }} - {{ product.Brand }}</p>
              <p style="color: var(--text-color--secondary);">{{ product.CaloriesPer100Grams }} kcal / 100 g</p>
            </div>
            <div class="product-list__icon--add" v-html="plusSvg" @click="addProduct(index)"></div>
          </div>
        </div>
        <input type="text" placeholder="Menge in Gramm" inputmode="numeric" pattern="[0-9]*" @input="handleWeightInput"/>
      </div>
    </div>
  </div>
</template>

<style>
.add-meal__overlay {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1;
}

.add-meal__card {
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

.add-meal__card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.add-meal__card-header h2 {
  margin-left: 5px;
}

.add-meal__close-button {
  width: 20px;
  height: 20px;
  cursor: pointer;
  margin-left: 30px;
  margin-top: 0px;
  scale: 1.2;
  fill: var(--icon-color);
}

.add-meal__card-content {
  padding: 0px 35px;
  text-align: center;
}

.add-meal__product {
  display: flex;
  margin: 10px 0;
}
.add-meal__product p {
  margin: 3px 0px;
}
.product__description {
  display: flex;
  flex-direction: column;
  justify-content: start;
  text-align: start;
}
.product-list__icon--add {
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
.product-list__icon--add > svg {
  height: 50px;
  width: 50px;
  scale: 0.75;
  fill: var(--button__text-color);
  margin: 0;
}




.add-meal__card-footer {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}
.add-meal__add-article--button {
  height: 35px;
  width: 35px;
  border-radius: 90px;
  border: none;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--button__background-color);
  scale: 0.9;
  fill: var(--button__text-color);
}
</style>