<script lang="ts" setup>
import plusSvg from '../assets/plus.svg?raw'
import { format } from 'date-fns';
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

const props = defineProps<{ shoppingLists: ShoppingList[]; lastId: number }>()

const emit = defineEmits(['close']);

function submitNewList(event: Event) {
  event.preventDefault();
  const name = (document.getElementById('name') as HTMLInputElement).value;
  const dateInput = (document.getElementById('date') as HTMLInputElement).value;
  if (name === '' || dateInput === '') {
    console.log('Name or date is empty');
    return;
  }
  const date = format(new Date(dateInput), 'dd.MM.yyyy');
  console.log('Name: ' + name + ' Date: ' + date);
  props.shoppingLists.push({ ID: props.lastId + 1, Name: name, Description: '', Date: date, Products: [] });
  //TODO: Add the new shopping list to the backend
  fetch(BASE_URL + '/shoppinglist/', {
    method: 'POST',
  })
  emit('close');
}
</script>

<template>
    <div class="shopping-list-new__overlay" @click.self="$emit('close')" style="width: 100%;">
        <div class="shopping-list-new__card">
          <div class="shopping-list-new__header">
            <h2>Neue Liste</h2>
            <div class="shopping-list-new__close-button" v-html="plusSvg" @click="$emit('close')"></div>
          </div>
          <div class="shopping-list-new__content">
            <form class="shopping-list-new__form" @submit="submitNewList($event)">
                <div class="shopping-list-new__item">
                  <span class="shopping-list-new__label">Name: </span>
                  <input type="text" class="shopping-list-new__input" id="name"/>
                </div>
                <div class="shopping-list-new__item">
                  <span class="shopping-list-new__label">Datum: </span>
                  <input type="date" class="shopping-list-new__input" id="date"/>
                </div>
                <input class="shopping-list-new__submit-button" type="submit" name="Erstellen" value="Erstellen"/>
            </form>
          </div>
        </div>
    </div>
</template>

<style>
.shopping-list-new__overlay {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1;
}

.shopping-list-new__card {
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
h2 {
  margin: 0;
}
.shopping-list-new__close-button > svg {
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

.shopping-list-new__header, .shopping-list-new__item {
  display: flex;
  justify-content: space-between;
}

.shopping-list-new__form {
  text-align: center;
  padding: 0px 35px;
  margin-block-start: 1em;
}
.shopping-list-new__item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 12px;
}
.shopping-list-new__item:last-child {
  margin-bottom: 0;
}
.shopping-list-new__label {
  display: flex;
  align-items: center;
}
.shopping-list-new__input {
  min-width: 215px;
}
.shopping-list-new__submit-button {
  background-color: var(--button__background-color);
  color: var(--button__text-color);
  border: none;
  padding: 5px 10px;
  font-size: 1.1rem;
}
</style>