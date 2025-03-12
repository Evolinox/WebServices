import { ref } from 'vue';
import BASE_URL from './baseUrl';

interface Settings {
  PlannedCalories: number;
  ProteinsInGrams: number;
  FatsInGrams: number;
  CarbsInGrams: number;
  Gender: string;
  WeightInKg: number;
}

export const settings = ref<Settings>({
  PlannedCalories: 0,
  ProteinsInGrams: 0,
  FatsInGrams: 0,
  CarbsInGrams: 0,
  Gender: '',
  WeightInKg: 0
});

export async function loadSettings() {
  try {
    const response = await fetch(BASE_URL + '/settings/', {
      method: 'GET',
    });
    if (response.status === 401) {
      window.location.href = '/login';
    } else if (response.status === 200) {
      const data = await response.json();
      settings.value = data;
    }
  } catch (error) {
    console.error('Error fetching settings:', error);
  }
}