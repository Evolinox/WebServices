<script lang="ts" setup>
import { ref } from 'vue';
import plusSvg from '../assets/plus.svg?raw';
import trashSvg from '../assets/trash.svg?raw';
import BASE_URL from '../baseUrl';
import { format } from 'date-fns';
import currentDay from '../day';

interface Appointment {
  ID: number;
  Date: string;
  BeginTime: string;
  EndTime: string;
  Name: string;
  Description: string;
}

const props = defineProps<{
  appointments: Array<Appointment>;
}>();
const emit = defineEmits(['close', 'reload']);

const beginTime = ref('')
const endTime = ref('')
const name = ref('')

function addAppointment() {
  if (!beginTime.value || !endTime.value || !name.value) {
    console.error('Start or end or name is empty')
    return
  }
  fetch(BASE_URL + '/calendar/', {
    method: 'POST',
    body: JSON.stringify({
      Date: format(currentDay.value, 'yyyy-MM-dd'),
      BeginTime: beginTime.value,
      EndTime: endTime.value,
      Name: name.value,
      Description: ''
    })
  })
  .then(response => {
    if(response.status === 201) {
      beginTime.value = ''
      endTime.value = ''
      name.value = ''
      emit('reload')
    } else {
      console.error('Error:', response);
    }
  })
}

function deleteAppointment(index: number) {
  fetch(BASE_URL + '/calendar/' + props.appointments[index].ID, {
    method: 'DELETE'
  })
  .then(response => {
    if(response.status === 200) {
      emit('reload')
    } else {
      console.error('Error:', response);
    }
  })
}
</script>

<template>
  <div class="day-appointments__overlay" @click.self="$emit('close')" style="width: 100%;">
    <div class="day-appointments__card">
      <div class="day-appointments__card-header">
        <h2>Termine {{ format(currentDay, 'dd.MM.yyyy') }}</h2>
        <div class="day-appointments__close-button" v-html="plusSvg" @click="$emit('close')"></div>
      </div>

      <div class="day-appointments__card-content">
        <ul class="day-appointments__list" v-if="appointments.length > 0">
          <li class="day-appointments__list-item" v-for="(appointment, index) in appointments">
            <span class="appointment__name"> {{ appointment.Name }} </span>
            <span class="appointment__begin-end"> {{ appointment.BeginTime }} - {{ appointment.EndTime }} Uhr </span>
            <div class="appointment__trash-icon" v-html="trashSvg" @click="deleteAppointment(index)"></div>
          </li>
        </ul>
        <p v-else style="text-align: center;">Keine Termine vorhanden</p>
      </div>

      <div class="day-appointments__add-appointment">
        <div class="add-appointment__input-container">
          <input class="add-appointment__input-name" type="text" v-model="name" placeholder="Beschreibung" />
          <div class="add-appointment__input-time-container">
            <input type="time" step="300" v-model="beginTime" placeholder="Start" />
            <input type="time" step="300" v-model="endTime" placeholder="Ende" />
          </div>
        </div>
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
  z-index: 1;
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
  margin-bottom: 10px;
}
.day-appointments__close-button {
  margin-bottom: auto;
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

.day-appointments__list {
  padding: 0px 15px;
}
.day-appointments__list-item {
  display: flex;
  align-items: center;
}
.appointment__name {
  max-width: 175px;
  overflow: scroll;
}
.appointment__begin-end {
  margin-left: auto;
  margin-right: 8px;
}

.appointment__trash-icon {
  cursor: pointer;
}
.appointment__trash-icon > svg {
  height: 20px;
  width: 20px;
  fill: var(--icon-color);
}
.appointment__trash-icon:hover > svg {
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
  padding: 0px 35px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-width: 300px;
}
.add-appointment__input-container > input:first-child {
  height: 20px;
  padding: 5px;
}
.add-appointment__input-name {
  width: 237px;
}
.add-appointment__input-time-container {
  display:flex; 
  justify-content: space-between;
  width: 250px;
  margin-top: 12px;
}
.add-appointment__input-time-container > input {
  width: 106px;
  height: 25px;
}
</style>
