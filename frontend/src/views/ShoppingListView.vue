<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import plusSvg from '../assets/plus.svg?raw';
import trashSvg from '../assets/trash.svg?raw';
import ShoppingListNew from '../components/ShoppingListNew.vue';
import ShoppingListShow from '../components/ShoppingListShow.vue';

const showAddShoppingList = ref(false);
const selectedListIndex = ref<number>(0);
function openShoppingListShow(index: number) {
  showAddShoppingList.value = !showAddShoppingList.value;
  console.log('toggle shopping list show ' + showAddShoppingList.value + ' index: ' + index + ' input?: ' + openWithInput.value)
  selectedListIndex.value = index;
}
function closeShoppingListShow() {
  showAddShoppingList.value = false;
  openWithInput.value = false;
}

const showNewShoppingList = ref(false);
function toggleShoppingListNew() {
  showNewShoppingList.value = !showNewShoppingList.value;
}

const openWithInput = ref(false);
function openAddArticleWithInput(index: number) {
  openWithInput.value = true;
  console.log('open with input?: ' + openWithInput.value)
  openShoppingListShow(index)
}

function deleteList(index: number) {
  console.log('Delete list number: ' + index)
  shoppingLists.value.splice(index, 1)
  // TODO: Update the shopping list in the backend
}

onMounted(() => {
  // TODO: Fetch the shopping lists from the backend
})
// Simulierte JSON-Daten
const shoppingLists = ref([
  {
    name: "Wocheneinkauf",
    date: "12.02.2025",
    products: [
      { name: "Äpfel", quantity: "500g" },
      { name: "Bananen", quantity: "6 Stück" },
      { name: "Milch", quantity: "2 Liter" }
    ]
  },
  {
    name: "Partyvorbereitung",
    date: "15.02.2025",
    products: [
      { name: "Chips", quantity: "3 Tüten" },
      { name: "Cola", quantity: "4 Flaschen" },
      { name: "Bier", quantity: "10 Kästen" }
    ]
  },
  {
    name: "Grillabend",
    date: "20.02.2025",
    products: [
      { name: "Steaks", quantity: "4 Stück"},
      { name: "Würstchen", quantity: "12 Stück" },
      { name: "Brot", quantity: "2 Leib" }
    ]
  }
])
</script>

<template>
  <div class="shopping-list">
    <div class="shopping-list__content">
      <div class="shopping-list__header">
        <h1>Einkaufsliste</h1>
        <button @click="toggleShoppingListNew()" v-html="plusSvg"></button>
      </div>
      <div class="shopping-list__lists">
        <div v-if="shoppingLists.length > 0" class="shopping-list__list" v-for="(shoppingList, index) in shoppingLists" :key="shoppingList.name">
          <div class="shopping-list__date-name" @click="openShoppingListShow(index)">
            <p><strong>{{ shoppingList.date }}</strong> - {{ shoppingList.name }}</p>
          </div>
          <div class="shopping-list__add-to-list--icon" v-html="plusSvg" @click="openAddArticleWithInput(index)"></div>
          <div class="shopping-list__delete-list--icon" v-html="trashSvg" @click="deleteList(index)"></div>
        </div>
        <p v-else style="text-align: center;">Keine Einkaufslisten vorhanden</p>
      </div>
    </div>
    <ShoppingListNew v-if="showNewShoppingList==true" @close="toggleShoppingListNew" :shoppingLists="shoppingLists"/>
    <ShoppingListShow v-if="showAddShoppingList==true" @close="closeShoppingListShow" :shoppingLists="shoppingLists" :selectedListIndex="selectedListIndex" :addArticleInput="openWithInput"/>
  </div>
</template>

<style>
.shopping-list {
  display: flex;
  flex-direction: column;
  width: calc(100% - 40px);
  padding: 20px;
  border-radius: 4px;
}
.shopping-list__content {
  background-color: var(--background-color--secondary);
  padding: 20px;
  border-radius: var(--border-radius__secondary-background);
}
.shopping-list__content:last-child p {
  margin-bottom: 0;
}

.shopping-list__header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-block-end: 0.7rem;
}
.shopping-list__header button{
  width: 40px;
  height: 40px;
  background-color: transparent;
  display: flex;
}
.shopping-list__header svg {
  width: 35px;
  height:35px;
  fill: var(--text-color--primary);
}

.shopping-list__list, .shopping-list__date-name {
  display: flex;
  align-items: center;
}
.shopping-list__date-name {
  width: -webkit-fill-available;
}
.shopping-list__list p {
  margin: 0px;
}

.shopping-list__add-to-list--icon {
  margin-right: 10px;
}
.shopping-list__add-to-list--icon,  .shopping-list__delete-list--icon{
  height: 35px;
  width: 35px;
  display: flex;
  align-items: center;
}
.shopping-list__add-to-list--icon > svg, .shopping-list__delete-list--icon > svg {
  height: 50px;
  width: 50px;
  scale: 0.8;
  fill: var(--icon-color);
}

.shopping-list__date-name, .shopping-list__add-to-list--icon, .shopping-list__delete-list--icon {
  cursor: pointer;
}
</style>