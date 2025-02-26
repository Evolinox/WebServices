<script lang="ts" setup>

</script>

<template>
    <div class="shopping-list-add__overlay" @click.self="$emit('close')">
        <div class="shopping-list-add__card">
          <div class="shopping-list-add__header">
            <h1>Neue Einkaufsliste</h1>
            <img class="shopping-list-add__close-button" src="../assets/plus.svg" @click="$emit('close')"/>
          </div>
          <div class="shopping-list-add__content">
            <div class="shopping-list-add__item">
              <span class="shopping-list-add__label">Name</span>
              <input type="text" class="shopping-list-add__input"/>
            </div>
            <div class="shopping-list-add__item">
              <span class="shopping-list-add__label">Datum</span>
              <input type="date" class="shopping-list-add__input"/>
            </div>
            <div class="shopping-list-add__item">
              <span class="shopping-list-add__label">Produkte</span>
              <ul>
                <li>
                  <input type="text" class="shopping-list-add__input"/>
                  <input type="number" class="shopping-list-add__input"/>
                </li>
              </ul>
            </div>

          </div>
        </div>
    </div>
</template>

<style>
.shopping-list-add__overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1;
}

.shopping-list-add__card {
  display: flex;
  flex-direction: column;
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: var(--background-color--secondary);
  padding: 1em;
  border-radius: 8px;
}
.shopping-list-add__close-button {
  width: 20px;
  height: 20px;
  margin-left: auto;
  margin-right: 0;
  rotate: 45deg;
  cursor: pointer;
}
</style>