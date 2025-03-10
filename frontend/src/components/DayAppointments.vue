<script lang="ts" setup>
import { ref } from 'vue';
import plusSvg from '../assets/plus.svg?raw';
import trashSvg from '../assets/trash.svg?raw';
import { format, parseISO } from 'date-fns';

const props = defineProps<{
  appointments: Array<{ date: string; description: string }>;
}>();

// Eingabefelder (Datum & Beschreibung)
const date = ref('')
const description = ref('')

// Funktion: Termin hinzufügen
function addAppointment() {
  if (!date.value || !description.value) {
    console.error('Date or description is empty')
    return
  }
  
  const dateFormatted = format(parseISO(date.value), 'dd.MM.yyyy')
  props.appointments.push({ date: dateFormatted, description: description.value })
  
  date.value = ''
  description.value = ''
}

// Termin löschen
function deleteAppointment(index: number) {
  props.appointments.splice(index, 1);
}
</script>

<template>
  <div class="day-appointments__overlay" @click.self="$emit('close')" style="width: 100%;">
    <div class="day-appointments__card">
      <div class="day-appointments__card-header">
        <h2>Tages Termine</h2>
        <div class="day-appointments__close-button" v-html="plusSvg" @click="$emit('close')"></div>
      </div>

      <div class="day-appointments__card-content">
        <ul v-if="appointments.length > 0">
          <li v-for="(appt, index) in appointments" :key="index">
            <span class="appt__date"> {{ appt.date }}:</span> 
            <span class="appt__desc"> {{ appt.description }} </span>
            <div class="appt__trash-icon" v-html="trashSvg" @click="deleteAppointment(index)"></div>
          </li>
        </ul>
        <p v-else style="text-align: center;">Keine Termine vorhanden</p>
      </div>

      <div class="day-appointments__add-appointment">
        <input type="date" v-model="date" placeholder="Datum" />
        <input type="text" v-model="description" placeholder="Beschreibung" />
        <button class="plus add-appointment__add-button" v-html="plusSvg" @click="addAppointment"></button>
      </div>
    </div>
  </div>
</template>

<style>
.day-appointments__overlay {
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  width: 100vw;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 999;
}

.day-appointments__card {
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

/* Header */
.day-appointments__card-header {
  align-items: center;
  display: flex;
  justify-content: space-between;
}
.day-appointments__card-header h2 {
  margin-left: 5px;
}
.day-appointments__close-button > svg {
  width: 20px;
  height: 20px;
  rotate: 45deg;
  cursor: pointer;
  margin-left: 30px;
  margin-right: 0;
  scale: 1.2;
  fill: var(--icon-color);
}

/* Content */
.day-appointments__card-content ul {
  padding: 0px 35px;
}
.day-appointments__card-content li {
  list-style-type: none;
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.day-appointments__card-content li:last-child {
  margin-bottom: 0;
}

.appt__desc {
  margin-left: 10px;
}

.appt__trash-icon {
  cursor: pointer;
}
.appt__trash-icon > svg {
  height: 20px;
  width: 20px;
  fill: var(--icon-color);
}
.appt__trash-icon:hover > svg {
  fill: var(--accent-color--primary);
}

/* Footer */
.day-appointments__card-footer {
  display: flex;
  justify-content: center;

}

/* Plus-Button */
.plus {
  height: 35px;
  width: 35px;
  border-radius: 90px;
  border: none;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--button__background-color);
  scale: 0.9;
}
.plus > svg {
  height: 50px;
  width: 50px;
  scale: 0.8;
  fill: var(--button__text-color);
}

/* Neues Termin-Eingabeformular */
.day-appointments__add-appointment {
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-width: 370px;
}
</style>
