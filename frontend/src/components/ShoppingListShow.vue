<script lang="ts" setup>
import { ref, watch } from 'vue'
import plusSvg from '../assets/plus.svg?raw'
import trashSvg from '../assets/trash.svg?raw'
import BASE_URL from '../baseUrl';

interface ProductGet {
  ID: number;
  Name: string;
  Quantity: string;
  ShoppingListID: number;
}

interface ShoppingListGet {
  ID: number;
  Name: string;
  Description: string;
  Date: string;
  Products: ProductGet[];
}

interface ProductSend {
  Name: string;
  Quantity: string;
  ShoppingListID: number;
}

const props = defineProps<{
  shoppingLists: ShoppingListGet[],
  selectedListIndex: number,
  addArticleInput: boolean
}>()
const listId = ref(props.shoppingLists[props.selectedListIndex].ID);
let currentList = props.shoppingLists[props.selectedListIndex];
const emit = defineEmits(['close', 'reload'])

const addArticleInput = ref(props.addArticleInput)
watch(() => props.addArticleInput, (newValue) => {
  addArticleInput.value = newValue
})
function openAddArticle() {
  addArticleInput.value = true
}

function addArticle(div: HTMLElement) {
  addArticleInput.value = false
  const articleName = (div.querySelector('#article') as HTMLInputElement).value
  const articleQuantity = (div.querySelector('#quantity') as HTMLInputElement).value
  if (articleName === '' || articleQuantity === '') {
    console.error('Article name or quantity is empty')
    return
  }
  const newProduct: ProductSend = {Name: articleName, Quantity: articleQuantity, ShoppingListID: currentList.ID}

  // Update the shopping list in the backend
  fetch (BASE_URL + '/shoppinglist/' + listId.value, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      Name: newProduct.Name,
      Quantity: newProduct.Quantity,
      ShoppingListID: newProduct.ShoppingListID
    })
  })
  .then(response => {
      if(response.status === 201) {
        emit('reload')
      }
    })
  .catch((error) => {
    console.error('Error:', error)
  })
}

function deleteProduct(index: number, productId: number) {
  // Update the shopping list in the backend
  currentList.Products.splice(index, 1)
  fetch (BASE_URL + '/shoppinglist/' + listId.value + '/products/' + productId, {
    method: 'DELETE',
  })
  .then(response => {
    if(response.status === 20) {
      emit('reload')
    }
  })
  .catch((error) => {
    console.error('Error:', error)
  })
}
</script>

<template>
    <div class="shopping-list-show__overlay" @click.self="$emit('close')" style="width: 100%;">
        <div class="shopping-list-show__card">
          <div class="shopping-list-show__card-header">
            <h2>{{ props.shoppingLists[selectedListIndex].Name }}</h2>
            <div class="shopping-list-show__close-button" v-html="plusSvg" @click="$emit('close')"></div>
          </div>
          <div class="shopping-list-show__card-content">
            <ul v-if="props.shoppingLists[selectedListIndex].Products.length > 0">
              <li v-for="(product, index) in props.shoppingLists[selectedListIndex].Products">
                <span>{{ product.Name }}:</span> <span class="product__quantity"> {{ product.Quantity }} </span>
                <div class="product__trash-icon" v-html="trashSvg" @click="deleteProduct(index, product.ID)"></div>
              </li>
            </ul>
            <p v-else style="text-align: center;">Keine Artikel vorhanden</p>
          </div>
          <div class="shopping-list-show__card-footer" v-if="!addArticleInput">
            <button class="shopping-list-show__open-add-article" v-html="plusSvg" @click="openAddArticle"></button>
          </div>
          <div class="shopping-list-show__add-article" v-if="addArticleInput">
            <input type="text" placeholder="Artikel" id="article">
            <input type="text" placeholder="Menge" id="quantity">
            <button class="add-article__add-button" v-html="plusSvg" @click="addArticle($el)"></button>
          </div>
        </div>
    </div>
</template>

<style>
.shopping-list-show__overlay {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1;
}

.shopping-list-show__card {
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

.shopping-list-show__card-header {
  align-items: center;
}
.shopping-list-show__card-header h2 {
  margin-left: 5px;
}
.shopping-list-show__close-button > svg {
  width: 20px;
  height: 20px;
  margin-right: 0;
  rotate: 45deg;
  cursor: pointer;
  margin-left: 30px;
  margin-top: 0px;
  scale: 1.2;
  fill: var(--icon-color);
}

.shopping-list-show__card-header, .shopping-list-show__item {
  display: flex;
  justify-content: space-between;
}
ul {
  padding: 0px 35px;
}
li {
  list-style-type: none;
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}
li:last-child {
  margin-bottom: 0;
}
.product__quantity {
  margin-left: auto;
  margin-right: 15px;
}
.product__trash-icon {
  cursor: pointer;
}
.product__trash-icon > svg {
  height: 20px;
  width: 20px;
  fill: var(--icon-color);
}
.product__trash-icon:hover > svg {
  fill: var(--accent-color--primary);
}

.shopping-list-show__card-footer {
  display: flex;
  justify-content: center;
}
.shopping-list-show__open-add-article, .add-article__add-button {
  height: 35px;
  width: 35px;
  border-radius: 90px;
  border: none;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--button__background-color);
  scale: 0.9;
}
.shopping-list-show__open-add-article > svg, .add-article__add-button > svg {
  height: 50px;
  width: 50px;
  scale: 0.8;
  fill: var(--button__text-color);
}

.shopping-list-show__add-article {
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-width: 370px;
}
</style>