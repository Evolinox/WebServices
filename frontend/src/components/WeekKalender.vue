<template>
  <div class="calendar-view">
    <div class="calendar-container">
    <div class="calendar-title">Wochen Kalender</div>
    <div class="calendar-header">
      <button @click="prevWeek" class="nav-button">⬅</button>
      <span class="week-text">{{ formattedWeek }}</span>
      <button @click="nextWeek" class="nav-button">➡</button>
    </div>
    <button class="appointment-button">Tages Termine</button>
  </div>
  </div>
    
  
</template>

<script setup>
import { ref, computed } from "vue";
import { format, startOfWeek, endOfWeek, addWeeks, subWeeks } from "date-fns";

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

<style scoped>

select, input {
  border-radius: 6px;
}

/* Container über gesamte Breite */
.calendar-view {
  width: 100vw;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  box-sizing: border-box;
  border-radius: 12px;
  margin-top: 5px;
}
.calendar-container {
  width: calc(100% - 40px);
  background: #3498db;
  padding: 2px;
  border-radius: 5px;
  text-align:center;
  color: white;
  font-family: "Arial", sans-serif;
  margin-bottom: 10px;
  
}

/* Kalender-Titel */
.calendar-title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 5px;
  
}

/* Header mit Navigation */
.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #2980b9;
  padding: 2px;
  border-radius: 5px;
  
}

/* Wochen-Text */
.week-text {
  font-size: 20px;
  font-weight: bold;
}

/* Navigationstasten */
.nav-button {
  background: white;
  color: #2980b9;
  border: none;
  padding: 10px 15px;
  font-size: 18px;
  font-weight: bold;
  border-radius: 5px;
  cursor: pointer;
}

/* Button für Termine */
.appointment-button {
  background: #f1c40f;
  border: none;
  padding: 10px 15px;
  font-size: 16px;
  font-weight: bold;
  border-radius: 5px;
  cursor: pointer;
  margin-top: 2px;
}
</style>
