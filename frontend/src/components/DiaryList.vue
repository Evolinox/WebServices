<script lang="ts" setup>
import { ref } from 'vue';
import trashSvg from '../assets/trash.svg?raw';

interface ProductDiary {
  ID: number;
  DailyProductsConsumedID: number;
  ProductID: number;
  ProductName: string;
  Category: string;
  WeightInGrams: number;
  Calories: number;
  ProteinsInGrams: number;
  FatsInGrams: number;
  CarbsInGrams: number;
}

const props = defineProps<{ list: ProductDiary[] }>();
const emit = defineEmits(['deleteMeal']);
const showInfo = ref(false);
const infoText = ref('');
const infoPosition = ref({ x: 0, y: 0 });

function showMacroInfo(event: MouseEvent,macro: string) {
  infoText.value = macro;
  infoPosition.value = { x: event.clientX, y: event.clientY };
  showInfo.value = true;
  setTimeout(() => {
    showInfo.value = false;
  }, 2000);
}
</script>

<template>
  <div>
    <ul>
      <li v-for="product in props.list" :key="product.ID" class="item-list__item">
        <div class="item-list__description">
          <div class="item-list__product">
            <p class="product__name-and-calories">{{ product.ProductName }} - {{ product.Calories }} kcal</p>
            <p class="product__brand-and-weight">Brand - {{ product.WeightInGrams }}g</p>
          </div>
          <div class="item-list__product-carbs">
            <span class="item-list__product-fats" @click="showMacroInfo($event, 'Fette')">{{ product.FatsInGrams.toFixed(2) }} g</span>
            <span class="item-list__product-carbs" @click="showMacroInfo($event, 'Kohlenhydrate')">{{ product.CarbsInGrams.toFixed(2) }} g</span>
            <span class="item-list__product-protein" @click="showMacroInfo($event, 'Protein')">{{ product.ProteinsInGrams.toFixed(2) }} g</span>
          </div>
        </div>
        <div class="item-list__delete-icon" v-html="trashSvg" @click="emit('deleteMeal', product.ID)"></div>
      </li>
    </ul>
    <div v-if="showInfo" class="macro-info" :style="{ top: infoPosition.y + 'px', left: infoPosition.x + 'px' }">{{ infoText }}</div>
  </div>
</template>

<style>
.item-list__item {
  display: flex;
  align-items: center;
}

.item-list__description {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  width: 85%;
}
.item-list__description p {
  margin: 0px;
}
.item-list__product {
  display: flex; flex-direction: column; justify-content: start;
}
.product__name-and-calories {
  font-size: 1.15rem;
}
.product__brand-and-weight {
  font-size: 0.95rem;
  color: var(--text-color--secondary);
}

.item-list__product-carbs {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  justify-content: space-between;
}
.item-list__product-carbs > span {
  font-size: 0.95rem;
  cursor: pointer;
}
.item-list__product-fats {
  color: red;
}
.item-list__product-carbs {
  color: orange;
}
.item-list__product-protein {
  color: green;
}

.item-list__delete-icon {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 35px;
  width: 35px;
  scale: 0.9;
  cursor: pointer;
}
.item-list__delete-icon > svg {
  height: 50px;
  width: 50px;
  scale: 0.75;
  fill: var(--icon-color);
  margin: 0px;
}

.macro-info {
  position: fixed;
  background-color: rgba(0, 0, 0, 0.8);
  color: white;
  padding: 10px;
  border-radius: 5px;
  z-index: 1000;
  font-size: 1rem;
}
</style>