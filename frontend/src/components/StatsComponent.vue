<script setup>
import { ref, computed } from "vue";
import { DoughnutChart, BarChart } from "vue-chart-3";
import { Chart, registerables } from "chart.js";

Chart.register(...registerables);

// Beispielwerte (später mit Backend ersetzen)
const consumedCalories = ref(577);
const calorieGoal = ref(2500);
const fat = ref(3);
const fatGoal = ref(54);
const carbs = ref(106);
const carbsGoal = ref(287);
const protein = ref(28);
const proteinGoal = ref(201);

// Berechnung der verbleibenden Kalorien
const remainingCalories = computed(() => calorieGoal.value - consumedCalories.value);

// Chart-Key für Neuladen der Diagramme
const chartKey = ref(0);

// Chart-Optionen
const chartOptions = ref({
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,
    },
  },
});

// Chart-Daten für Kalorienübersicht
const caloriesChartData = computed(() => ({
  labels: ["Verbraucht", "Übrig"],
  datasets: [
    {
      data: [consumedCalories.value, remainingCalories.value],
      backgroundColor: ["#4A90E2", "#E0E0E0"],
    },
  ],
}));

const macros = computed(() => [
  { name: "Fette", value: fat.value, goal: fatGoal.value, color: "red" },
  { name: "Kohlenhydrate", value: carbs.value, goal: carbsGoal.value, color: "orange" },
  { name: "Proteine", value: protein.value, goal: proteinGoal.value, color: "green" },
]);
</script>

<template>
  <div class="stats-container">
    <!-- Kalorienübersicht mit Kreisdiagramm -->
    <div class="calories-overview">
      <h2>Kalorienübersicht</h2>
      <div class="calories-flex">
        <DoughnutChart
          v-if="caloriesChartData"
          class="doughnut-chart"
          :chart-data="caloriesChartData"
          :chart-options="chartOptions"
          :key="chartKey"
        />
        <div class="calories-info">
          <p>Gegessen: <strong>{{ consumedCalories }}</strong> kcal</p>
          <p>Ziel: <strong>{{ calorieGoal }}</strong> kcal</p>
          <p>Übrig: <strong>{{ remainingCalories }}</strong> kcal</p>
        </div>
      </div>
    </div>
  
    <!-- Makronährstoffe mit Balkendiagramm -->      
    <div class="macros">
      <h2>Makronährstoffe</h2>
      <div class="macros-row">
        <div class="macro-bar" v-for="macro in macros" :key="macro.name">
          <p class="macro-label">{{ macro.name }}</p>
          <div class="progress-bar">
            <div class="progress-fill" :style="{ width: Math.min((macro.value / macro.goal) * 100, 100) + '%', backgroundColor: macro.color }"></div>
          </div>
          <small>{{ macro.value }}g / {{ macro.goal }}g</small>
        </div>
      </div>
    </div>
  </div>
</template>
  
<style scoped>
.stats-container {
  background: var(--background-color--secondary);
  padding: 20px;
  width: calc(100% - 40px);
  border-radius: var(--border-radius__secondary-background);
  text-align: center;
  box-sizing: border-box;
}

/* Kalorienübersicht */
.calories-overview h2 {
  margin-top: 0px;
}

.calories-flex {
  display: flex;
  align-items: center;  
  justify-content: center;
}

.calories-info {
  text-align: left;

}
  
.doughnut-chart {    
    height: 230px;
}
  
  /* Makronährstoffe */
  .macros {
  margin-top: 15px;
  max-width: 400px;
  margin-left: auto;
  margin-right: auto;
  text-align: left;
}

.macros h2 {
    text-align: center;
  margin-top: 20px;
}

.macros-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 20px;
}

.macro-bar {
    flex: 1;
  margin-bottom: 10px;
  min-width: 80px;
}

.macro-label {
  margin-bottom: 5px;
  text-align: center;
}

.progress-bar {
  width: 100%;
  height: 10px;
  background-color: #e0e0e0;
  border-radius: var(--border-radius__secondary-background);
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  border-radius: var(--border-radius__secondary-background);
}
</style>