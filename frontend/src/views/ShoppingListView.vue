<script lang="ts" setup>
import { ref } from 'vue'
import plusSvg from '../assets/plus.svg?raw';
import ShoppingListAdd from '../components/ShoppingListAdd.vue';

let showAddShoppingList = ref(false);
function toggleShoppingListAdd() {
  console.log('toggleShoppingListAdd ' + showAddShoppingList.value);
  showAddShoppingList.value = !showAddShoppingList.value;
  console.log('toggleShoppingListAdd ' + showAddShoppingList.value);
}

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
      { name: "Steaks", quantity: 4 },
      { name: "Würstchen", quantity: 12 },
      { name: "Brot", quantity: 2 }
    ]
  }
])
</script>

<template>
  <div class="shopping-list">
    <div class="shopping-list__content">
      <div class="shopping-list__header">
        <h1>Einkaufsliste</h1>
        <button @click="toggleShoppingListAdd()" v-html="plusSvg"></button>
      </div>
      <div v-for="shoppingList in shoppingLists" :key="shoppingList.name">
        <h2>{{ shoppingList.name }}</h2>
        <p>{{ shoppingList.date }}</p>
        <ul>
          <li v-for="product in shoppingList.products" :key="product.name">
            {{ product.name }}: {{ product.quantity }}
          </li>
        </ul>
      </div>
    </div>
    <ShoppingListAdd v-if="showAddShoppingList==true" @close="toggleShoppingListAdd"/>
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
.shopping-list__header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.shopping-list__header button{
  width: 40px;
  height: 40px;
  background-color: transparent;
}
.shopping-list__header button {
  display: flex;
}
.shopping-list__header svg {
  width: 35px;
  height:35px;
}
</style>