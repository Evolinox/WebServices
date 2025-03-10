<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { format } from 'date-fns';
import plusSvg from '../assets/plus.svg?raw';
import trashSvg from '../assets/trash.svg?raw';
import ShoppingListNew from '../components/ShoppingListNew.vue';
import ShoppingListShow from '../components/ShoppingListShow.vue';
import BASE_URL from '../baseUrl';

interface Product {
  ID: number;
  Name: string;
  Quantity: string;
  ShoppingListID: number;
}

interface ShoppingList {
  ID: number;
  Name: string;
  Description: string;
  Date: string;
  Products: Product[];
}

const showAddShoppingList = ref(false);
const selectedListIndex = ref<number>(0);
const shoppingLists = ref<ShoppingList[]>([]);

onMounted(() => {
  loadLists();
});

function openShoppingListShow(index: number) {
  showAddShoppingList.value = !showAddShoppingList.value;
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
  openShoppingListShow(index)
}

function deleteList(index: number, shoppingList: ShoppingList) {
  shoppingLists.value.splice(index, 1)
  fetch (BASE_URL + '/shoppinglist/' + shoppingList.ID, {
    method: 'DELETE',
  })
  .then(response => {
    if(response.status != 200) {
      console.error('Error:', response);
    }
  })
}

function loadLists() {
  fetch(BASE_URL + '/shoppinglist/')
    .then(response => response.json())
    .then(data => {
      shoppingLists.value = data.map((list: ShoppingList) => ({
        ID: list.ID,
        Name: list.Name,
        Description: list.Description,
        Date: list.Date,
        Products: Array.isArray(list.Products) ? list.Products.map((product: Product) => ({
          ID: product.ID,
          Name: product.Name,
          Quantity: product.Quantity,
          ShoppingListID: product.ShoppingListID
        })) : []
      }));
    })
    .catch(error => {
      console.error('Error fetching shopping lists:', error);
    });
}
</script>

<template>
  <div class="shopping-list">
    <div class="shopping-list__content">
      <div class="shopping-list__header">
        <h1>Einkaufsliste</h1>
        <button @click="toggleShoppingListNew()" v-html="plusSvg"></button>
      </div>
      <div class="shopping-list__lists">
        <div v-if="shoppingLists.length > 0" class="shopping-list__list" v-for="(shoppingList, index) in shoppingLists" :key="shoppingList.Name">
          <div class="shopping-list__date-name" @click="openShoppingListShow(index)">
            <p><strong>{{ format(new Date(shoppingList.Date), 'dd.MM.yyyy') }}</strong> - {{ shoppingList.Name }}</p>
          </div>
          <div class="shopping-list__add-to-list--icon" v-html="plusSvg" @click="openAddArticleWithInput(index)"></div>
          <div class="shopping-list__delete-list--icon" v-html="trashSvg" @click="deleteList(index, shoppingList)"></div>
        </div>
        <p v-else style="text-align: center;">Keine Einkaufslisten vorhanden</p>
      </div>
    </div>
    <ShoppingListNew v-if="showNewShoppingList==true" @close="toggleShoppingListNew" @reload="loadLists" />
    <ShoppingListShow v-if="showAddShoppingList==true" @close="closeShoppingListShow" @reload="loadLists" :shoppingLists="shoppingLists" :selectedListIndex="selectedListIndex" :addArticleInput="openWithInput" />
  </div>
</template>

<style>
.shopping-list {
  display: flex;
  flex-direction: column;
  width: calc(100% - 40px);
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