<template>
  <div class="appointments-overlay" @click.self="$emit('close')">
    <div class="appointments-card">
      <div class="appointments-header">
        <h2>Tages Termine</h2>
        <div class="appointments-close-button" v-html="plusSvg" @click="$emit('close')"></div>
      </div>
      <div class="appointments-content">
        <ul v-if="appointments.length > 0">
          <li v-for="(appointment, index) in appointments" :key="index">
            <span class="appointment-date"><strong>{{ formatDate(appointment.date) }}</strong>:</span> 
            <span class="appointment-description">{{ appointment.description }}</span>
            <div class="appointment-trash-icon" v-html="trashSvg" @click="$emit('delete', index)"></div>
          </li>
        </ul>
        <p v-else style="text-align: center;">Keine Termine vorhanden</p>
      </div>
      <div class="appointments-footer">
        <input type="date" v-model="date" placeholder="Datum">
        <input type="text" v-model="description" placeholder="Beschreibung">
        <button class="plus add-appointment__add-button" v-html="plusSvg" @click="addAppointment"></button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import plusSvg from '../assets/plus.svg?raw';
import trashSvg from '../assets/trash.svg?raw';
import { format, parseISO  } from 'date-fns';
import { ref } from "vue";

const props = defineProps<{ appointments: Array<{ date: string; description: string }> }>();
const emit = defineEmits(['close', 'add', 'delete']);

const addAppointmentInput = ref(false);
const date = ref("");
const description = ref("");


const formatDate = (dateString: string) => {
  return format(parseISO(dateString), "dd.MM.yyyy");
};

// Neuen Termin hinzufügen
const addAppointment = () => {
  if (description.value === '' || date.value === '') {
    console.log('Beschreibung oder Datum ist leer');
    return;
  }

  emit('add', date.value, description.value);

  // Eingaben zurücksetzen
  date.value = "";
  description.value = "";
  addAppointmentInput.value = false;
}
</script>

<style>
.appointments-overlay {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 100;
}

.appointments-card {
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

.appointments-header {
  display: flex;
  justify-content: space-between;
}

.appointments-footer {
  display: flex;
  justify-content: center;
}

.add-appointment__add-button {
  background: #4CAF50;
  color: white;
}

.appointment-trash-icon > svg{
  height: 20px;
  width: 20px;
  fill: var(--icon-color);
}

.appointment-description {
  margin-left: 10px;
}

.appointments-close-button > svg {
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

</style>
