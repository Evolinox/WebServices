import { ref } from 'vue'

export type ColorSchemeValue = 'light' | 'dark';

export const theme = ref<ColorSchemeValue>(localStorage.getItem('theme') as ColorSchemeValue || 'dark')

export function toggleTheme() {
  theme.value = theme.value === 'light' ? 'dark' : 'light'
  localStorage.setItem('theme', theme.value)
}