<script setup>
import { ref, computed, onMounted, watch } from "vue";
import { DoughnutChart, BarChart } from "vue-chart-3";
import { Chart, registerables } from "chart.js";
import BASE_URL from "../baseUrl";
import currentDay from "../day.ts";
import { format } from "date-fns";
import { loadSettings, settings } from "../settings";
  
Chart.register(...registerables);

const consumedCalories = ref(577);
const calorieGoal = ref(2500);
const fat = ref(3);
const fatGoal = ref(54);
const carbs = ref(106);
const carbsGoal = ref(287);
const protein = ref(28);
const proteinGoal = ref(201);
let date = ref("1900-01-01");

onMounted (() => {
  patchGoals();
  patchStats();
})

watch(currentDay, () => {
  patchStats();
})

watch(settings, () => {
  patchGoals();
})

const remainingCalories = computed(() => Math.max(calorieGoal.value - consumedCalories.value, 0));
const balanceCalories = computed(() => calorieGoal.value - consumedCalories.value);

// Chart-Key for reload of diagramm
const chartKey = ref(0);

const chartOptions = ref({
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,      
    }, 
    labels: {
      color: "white",
    },   
  },
});

// Chart-Data for calories chart
const caloriesChartData = computed(() => ({
  labels: ["Gegessen", "Übrig"],
  datasets: [
    {
      data: [consumedCalories.value, remainingCalories.value],
      backgroundColor: ["#0787d4", "#E0E0E0"],
    },
  ],
}));

let macros = computed(() => [
  { name: "Fette", value: fat.value, goal: fatGoal.value, color: "red" },
  { name: "Kohlenhydrate", value: carbs.value, goal: carbsGoal.value, color: "orange" },
  { name: "Proteine", value: protein.value, goal: proteinGoal.value, color: "green" },
]);

function patchStats() {
  date = format(currentDay.value, "yyyy-MM-dd");
  fetch(BASE_URL + '/nutrition/' + date, {
    method: 'GET',
  })
  .then(response => response.json())
  .then((data) => {
    consumedCalories.value = data.ConsumedCalories.toFixed(2);
    fat.value = data.ConsumedFats.toFixed(2);
    carbs.value = data.ConsumedCarbs.toFixed(2);
    protein.value = data.ConsumedProteins.toFixed(2);
    loadSettings();
  })
  .catch((error) => {
    console.error('Error:', error);
  });
}

function patchGoals() {
  calorieGoal.value = settings.value.PlannedCalories;
  fatGoal.value = settings.value.FatsInGrams;
  carbsGoal.value = settings.value.CarbsInGrams;
  proteinGoal.value = settings.value.ProteinsInGrams;
}
</script>

<template>
  <div class="stats-container">
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
          <p v-if="balanceCalories>=0">Übrig: <strong>{{ remainingCalories }}</strong> kcal</p>
            <p v-else>Zuviel: <strong>{{ Math.abs(balanceCalories) }}</strong> kcal</p>
        </div>
      </div>
    </div>
  
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

/* calories overview */
.calories-overview h2 {
  margin-top: 0px;
  margin-bottom: 15px;
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

.legend-label {
  color: white;
}
  
/* makro nutrition */
.macros {
  margin-top: 30px;
  max-width: 400px;
  margin-left: auto;
  margin-right: auto;
  text-align: left;
}

.macros h2 {
  text-align: center;
  margin-bottom: 10px;
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