<script lang="ts" setup>
import { computed, onMounted, ref, watch } from "vue";
import { format, addDays, subDays } from "date-fns";
import leftArrow from '../assets/left-arrow.svg?raw';
import currentDay from "../day";
import DayAppointments from "./DayAppointments.vue";
import BASE_URL from "../baseUrl";
import arrowSvg from '../assets/left-arrow.svg?raw';

interface Appointment {
  ID: number;
  Date: string;
  BeginTime: string;
  EndTime: string;
  Name: string;
  Description: string;
}

const formattedDay = computed(() => {
  return format(currentDay.value, "dd.MM.yyyy");
});

const nextDay = () => {
  currentDay.value = addDays(currentDay.value, 1);
};
const prevDay = () => {
  currentDay.value = subDays(currentDay.value, 1);
};

const openDayAppointments = ref(false);

const dayAppointments = ref<Appointment[]>([
  { ID: 999, Date: "01.05.2025", BeginTime: "15:00 Uhr", EndTime: "16:00 Uhr", Name: "Tag der Arbeit", Description: "Beschreibung" },
]);

onMounted(() => {
  loadAppointments();
});

watch(currentDay, () => {
  loadAppointments();
});

function loadAppointments() {
  fetch( BASE_URL + '/calendar/' + format(currentDay.value, 'yyyy-MM-dd') )
  .then(response => {
    if(response.status === 200) {
      return response.json();
    } else {
      console.error('Error:', response);
    }
  })
  .then(data => {
    console.log(data)
    dayAppointments.value = data
  })
}
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
      <div class="appointments-container" @click="openDayAppointments = true" >
        <button class="appointments-button">
          <span class="appointments-badge">{{ dayAppointments.length }}</span>
          <div class="appointments-button" v-html="arrowSvg"></div>
        </button>
      </div>
    </div>
  </div>
  <DayAppointments
    v-if="openDayAppointments"
    :appointments="dayAppointments"
    @close="openDayAppointments = false"
    @reload="loadAppointments"
  />
</template>

<style>
select, input {
  border-radius: 6px;
}

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
  justify-content: space-between;
  height: calc(var(--day-calender__height) - 20px);
  background: var(--background-color--secondary);
  padding: 10px 20px;
  border-radius: var(--border-radius__secondary-background);
  text-align: center;
  width: calc(100% - 40px);
}

.calendar-title {
  font-size: 24px;
  font-weight: bold;
}

.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--accent-color--primary);
  padding: 4px;
  border-radius: 5px;
  color: var(--button__text-color);
}

.day-text {
  font-size: 20px;
  font-weight: bold;
}

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

.appointments-container {
  display: flex;
  justify-content: center;
  cursor: pointer;
}

.appointments-button {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  background: transparent;
  position: relative;
  top: -8px;
  padding: 10px 15px;
  border-radius: 5px;
  border: none;
  width: 50px;
  height: 25px;
}
.appointments-button > svg {
  width: 20px;
  height: 20px;
  fill: var(--icon-color);
  rotate: -90deg;
  scale: 1.6;
}
.appointments-badge {
  display: inline;
  position: relative;
  top: 12px;
  right: -20px;
  background-color: var(--accent-color--primary);
  color: white;
  border-radius: 50%;
  padding: 5px 10px;
  font-size: 12px;
  font-weight: bold;
  z-index: 1;
  scale: 0.7;
}
</style>