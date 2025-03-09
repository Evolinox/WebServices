<script lang="ts" setup>
import { computed, ref } from "vue";
import { format, addDays, subDays } from "date-fns";
import leftArrow from '../assets/left-arrow.svg?raw';
import currentDay from "../day";
import DayAppointments from "./DayAppointments.vue";

// Berechnet das Datum des aktuellen Tages
const formattedDay = computed(() => {
  return format(currentDay.value, "dd.MM.yyyy");
});

// Funktionen zum Wechseln der Tage
const nextDay = () => {
  currentDay.value = addDays(currentDay.value, 1);
};
const prevDay = () => {
  currentDay.value = subDays(currentDay.value, 1);
};

const openDayAppointments = ref(false);
const addAppointmentInput = ref(false);

// Beispielhafte Terminliste
const appointments = ref([
  { date: "01.05.2025", description: "Tag der Arbeit" },
]);
</script>

<template>
  <div class="calendar-view" :class="$attrs.class">
    <div class="calendar-container">
      <div class="calendar-title">Tages Kalender</div>
      <div class="calendar-header">
        <div class="nav-button" @click="prevDay" v-html="leftArrow"></div>
        <span class="day-text">{{ formattedDay }}</span>
        <div class="nav-button rotate" @click="nextDay" v-html="leftArrow"></div>
      </div>
      <div class="appointments-button-container">
        <button class="appointments-button" @click="openDayAppointments = true">Tages Termine</button>
      </div>
    </div>
  </div>
  <DayAppointments
    v-if="openDayAppointments"
    :appointments="appointments"
    :selectedDateIndex="0"
    :addAppointmentInput="addAppointmentInput"
    @close="openDayAppointments = false"
  />
</template>

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
  width: calc(100% - 40px);
  margin-top: 5px;
  margin-bottom: 10px;

}
.calendar-container {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  height: calc(var(--day-calender__height) - 20px);
  background: var(--background-color--secondary);
  padding: 10px 20px;
  border-radius: var(--border-radius__secondary-background);
  text-align: center;
  width: calc(100% - 40px);
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

/* Tag-Text */
.day-text {
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

.appointments-button-container {
  margin-top: 10px;
  text-align: center;
}

.appointments-button {
  background: #f1c40f;
  padding: 10px 15px;
  font-size: 16px;
  font-weight: bold;
  border-radius: 5px;
  color: var(--button__text-color);
  border: none;
  cursor: pointer;
}
</style>