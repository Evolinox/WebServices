<script lang="ts" setup>
import { ref } from 'vue'
import { theme, toggleTheme } from '../themes'

const gender = ref('divers');
const weight = ref<number | null>(null);
const dailyCalories = ref<number | null>(null);

function handleGenderInput(event: Event) {
  const select = event.target as HTMLSelectElement;
  gender.value = select.value;
}

const handleWeightInput = (event: Event) => {
  const input = event.target as HTMLInputElement
  if (input?.value == "") {
    input.value = weight.value ? weight.value.toString() : "";
  } else {
    weight.value = input.value ? parseFloat(input.value) : null;
  }
}

const handleCaloriesInput = (event: Event) => {
  const input = event.target as HTMLInputElement
  if (input?.value == "") {
    input.value = dailyCalories.value ? dailyCalories.value.toString() : "";
  } else {
    dailyCalories.value = input.value? parseFloat(input.value) : null;
  }
}

function storeSettings() {
  console.log('Store settings: ', gender.value, " weight:", weight.value, "calories: ", dailyCalories.value);
  console.log('TODO: Store settings in backend');
  
}
</script>

<template>
  <div class="settings-view">
    <h1 class="settings-view__title">Einstellungen</h1>
    <div class="settings-view__content">
      <div class="settings-view__item">
        <span class="settings-view__label">Darkmode</span>
        <label class="switch">
          <input type="checkbox" @change="toggleTheme" :checked="theme === 'dark'">
          <span class="switch__slider switch__slider--round"></span>
        </label>
      </div>
    </div>
    <br>
    <div class="settings-view__content">
      <div class="settings-view__item">
        <span class="settings-view__label">Geschlecht</span>
        <select class="settings-view__select" @change="handleGenderInput($event)">
          <option value="männlich">männlich</option>
          <option value="weiblich">weiblich</option>
          <option value="divers">divers</option>
        </select>
      </div>
      <div class="settings-view__item">
        <span class="settings-view__label">Gewicht [kg]</span>
        <input type="number" class="settings-view__input" min="0" @input="handleWeightInput"/>
      </div>
      <div class="settings-view__item">
        <span class="settings-view__label">Kalorientagesziel [kcal]</span>
        <input type="number" class="settings-view__input" min="0" @input="handleCaloriesInput"/>
      </div>
      <div class="settings-view__footer">
        <button class="settings-footer__store-button" @click="storeSettings">Speichern</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
select, input {
  border-radius: 6px;
  box-shadow: none;
  border: solid var(--text-color--secondary);
}
/* Container für die Einstellungen */
.settings-view {
  height: calc(100vh - var(--nav-bar__height));
  width: 100vw;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 15px;
  box-sizing: border-box;
}

/* Stil für die Überschrift */
.settings-view__title {
  font-size: 1.2rem;
  margin-bottom: 20px;
}

.settings-view__content {
  width: calc(100% - 40px);
  padding: 20px;
  background-color: var(--background-color--secondary);
  border-radius: var(--border-radius__secondary-background);
}
.settings-view__item:last-child {
  margin-bottom: 0px;
}

/* Stil für die einzelnen Einstellungselemente */
.settings-view__item {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  margin-bottom: 20px;
}

/* Stil für die Labels */
.settings-view__label {
  font-size: 1rem;
  margin-right: 10px;
  flex: 1;
}

/* Stil für die Select-Elemente */
.settings-view__select, .settings-view__input {
  font-size: 1rem;
  padding: 5px;
  width: 50%;
  box-sizing: border-box;
  background-color: var(--background-color--tertiary);
  color: var(--text-color--primary);
}

/* Schalter-Stil */
.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
  margin-left: 10px;
}

/* Verstecke das eigentliche Kontrollkästchen */
.switch input { 
  opacity: 0;
  width: 0;
  height: 0;
}

/* Slider-Stil */
.switch__slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: .4s;
  border-radius: 34px;
}

/* Stil für den Schieberegler */
.switch__slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: .4s;
  border-radius: 50%;
}

/* Stil für den Schalter, wenn er aktiviert ist */
input:checked + .switch__slider {
  background-color: #2196F3;
}

/* Stil für den Schalter, wenn er fokussiert ist */
input:focus + .switch__slider {
  box-shadow: 0 0 1px #2196F3;
}

/* Stil für den Schieberegler, wenn der Schalter aktiviert ist */
input:checked + .switch__slider:before {
  transform: translateX(26px);
}

/* Modifikator für runde Slider */
.switch__slider--round {
  border-radius: 34px;
}
.switch__slider--round:before {
  border-radius: 50%;
}

.settings-view__footer {
  text-align: center;
}
.settings-footer__store-button {
  padding: 10px 20px;
  color: var(--button__text-color);
}
</style>