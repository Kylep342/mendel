<script setup lang="ts">
import { computed } from 'vue';

import constants from '@/constants/constants';
import useMendelCoreStore from '@/stores/core';
import PlantSpeciesForm from '@/components/forms/PlantSpeciesForm.vue';
import { useModal } from '@/composables/useModal';

const state = useMendelCoreStore();

useModal(computed(() => state.plantSpeciesFormActive), constants.ID_PLANT_SPECIES_FORM)
</script>

<template>
  <header>
    <h3>{{ constants.TITLE_PLANT_SPECIES }}</h3>
    <base-button @click="state.showPlantSpeciesForm">{{ constants.BTN_CREATE }}</base-button>
  </header>
  <div id="forms">
    <PlantSpeciesForm :id="constants.ID_PLANT_SPECIES_FORM"/>
  </div>
</template>

<style scoped>
header {
  /* Core layout with Flexbox */
  display: flex;
  justify-content: space-between; /* This pushes Mendel and Menu to opposite ends */
  align-items: center;

  /* Positioning */
  position: fixed;
  top: 4rem;
  left: 0;
  width: 100%;
  z-index: 900;

  /* Styling */
  background-color: var(--secondary-color);
  color: var(--text-on-secondary);
  padding: 0 1.5rem;
  height: 3rem;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

</style>