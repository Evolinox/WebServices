<script lang="ts" setup>
import { ref, watch } from 'vue'
import plusSvg from '../assets/plus.svg?raw'
import trashSvg from '../assets/trash.svg?raw'

const props = defineProps<{
  shoppingLists: Array<{ name: string; date: string; products: Array<{ name: string; quantity: string | number }> }>
  selectedListIndex: number
  addArticleInput: boolean
}>()

const addArticleInput = ref(props.addArticleInput)
watch(() => props.addArticleInput, (newValue) => {
  addArticleInput.value = newValue
})
function openAddArticle() {
  console.log('Open add article to list ' + props.shoppingLists[props.selectedListIndex].name)
  addArticleInput.value = true
}

function addArticle(div: HTMLElement) {
  addArticleInput.value = false
  const articleName = (div.querySelector('#article') as HTMLInputElement).value
  const articleQuantity = (div.querySelector('#quantity') as HTMLInputElement).value
  if (articleName === '' || articleQuantity === '') {
    console.log('Article name or quantity is empty')
    return
  }
  props.shoppingLists[props.selectedListIndex].products.push({ name: articleName, quantity: articleQuantity })
  // Update the shopping list in the backend
}

function deleteProduct(index: number) {
  console.log('Delete product')
  props.shoppingLists[props.selectedListIndex].products.splice(index, 1)
  // Update the shopping list in the backend
}
</script>

<template>
    <div class="shopping-list-show__overlay" @click.self="$emit('close')" style="width: 100%;">
        <div class="shopping-list-show__card">
          <div class="shopping-list-show__card-header">
            <h2>{{ props.shoppingLists[selectedListIndex].name }}</h2>
            <div class="shopping-list-show__close-button" v-html="plusSvg" @click="$emit('close')"></div>
          </div>
          <div class="shopping-list-show__card-content">
            <ul v-if="props.shoppingLists[selectedListIndex].products.length > 0">
              <li v-for="(product, index) in props.shoppingLists[selectedListIndex].products">
                <span>{{ product.name }}:</span> <span class="product__quantity"> {{ product.quantity }} </span>
                <div class="product__trash-icon" v-html="trashSvg" @click="deleteProduct(index)"></div>
              </li>
            </ul>
            <p v-else style="text-align: center;">Keine Artikel vorhanden</p>
          </div>
          <div class="shopping-list-show__card-footer" v-if="!addArticleInput">
            <button class="plus" v-html="plusSvg" @click="openAddArticle"></button>
          </div>
          <div class="shopping-list-show__add-article" v-if="addArticleInput">
            <input type="text" placeholder="Artikel" id="article">
            <input type="text" placeholder="Menge" id="quantity">
            <button class="plus add-article__add-button" v-html="plusSvg" @click="addArticle($el)"></button>
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
.plus {
  height: 35px;
  width: 35px;
  border-radius: 90px;
  border: none;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: transparent;
  background-color: var(--button__background-color);
  scale: 0.9;
}
.plus > svg {
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