<template>
  <div class="calendar-view">
    <div class="calendar-container">
      <div class="calendar-title">Wochen Kalender</div>
      <div class="calendar-header">
        <div class="nav-button" @click="prevWeek" v-html="leftArrow"></div>
        <span class="week-text">{{ formattedWeek }}</span>
        <div class="nav-button rotate" @click="nextWeek" v-html="leftArrow"></div>
      </div>
      <div>
        <button class="appointment-button">Tages Termine</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from "vue";
import { format, startOfWeek, endOfWeek, addWeeks, subWeeks } from "date-fns";
import leftArrow from '../assets/left-arrow.svg?raw';

// Startdatum der Woche initialisieren
const currentWeek = ref(new Date());

// Berechnet das Start- und Enddatum der Woche
const formattedWeek = computed(() => {
  const start = startOfWeek(currentWeek.value, { weekStartsOn: 1 });
  const end = endOfWeek(currentWeek.value, { weekStartsOn: 1 });
  return `${format(start, "dd.MM.yyyy")} - ${format(end, "dd.MM.yyyy")}`;
});

// Funktionen zum Wechseln der Wochen
const nextWeek = () => {
  currentWeek.value = addWeeks(currentWeek.value, 1);
};
const prevWeek = () => {
  currentWeek.value = subWeeks(currentWeek.value, 1);
};
</script>

<style>

select, input {
  border-radius: 6px;
}

/* Container über gesamte Breite */
.calendar-view {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  box-sizing: border-box;
  padding: 20px;
  width: 100vw;
}
.calendar-container {
  display: flex;
  flex-direction: column;
  justify-content: space-around;

  height: calc(var(--week-kalender__height) - 20px);

  background: var(--background-color--secondary);
  padding: 10px 20px;
  border-radius: var(--border-radius__secondary-background);
  text-align: center;
}

/* Kalender-Titel */
.calendar-title {
  font-size: 24px;
  font-weight: bold;
}

/* Header mit Navigation */
.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--accent-color--primary);
  padding: 4px;
  border-radius: 5px;
  color: var(--button__text-color);
}

/* Wochen-Text */
.week-text {
  font-size: 20px;
  font-weight: bold;
}

/* Navigationstasten */
.nav-button {
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: inherit;
  background: #fefeff;
  width: 50px;
  height: 50px;
  cursor: pointer;
}
.nav-button > svg {
  width: 20px;
  height: 20px;
  fill: var(--accent-color--primary);
}
.rotate > svg{
  rotate: 180deg;
}

/* Button für Termine */
.appointment-button {
  background: #f1c40f;
  padding: 10px 15px;
  font-size: 16px;
  font-weight: bold;
  border-radius: 5px;
  color: var(--button__text-color);
}
</style>
