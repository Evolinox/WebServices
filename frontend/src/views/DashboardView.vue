<script lang="ts" setup>
  import StatsComponent from '../components/StatsComponent.vue'
</script>

<template>
  <div class="dashboard">
    <StatsComponent />
  </div>
</template>

<style scoped>

.dashboard {
  display: flex;
  flex-direction: column;
  align-items: center;
  /* justify-content: center; */
  height: 100%;
}

</style>