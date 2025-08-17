<script setup lang="ts">
import { computed, onMounted } from 'vue';

import constants from '@/constants/constants';
import usePlantCultivarStore from '@/components/plant_cultivar/store';
import PlantCultivarForm from '@/components/plant_cultivar/CreateForm.vue';
import { useModal } from '@/composables/useModal';

const state = usePlantCultivarStore();

useModal(computed(() => state.plantCultivarFormActive), constants.ID_PLANT_CULTIVAR_FORM);

// Fetch the data when the component is first mounted
onMounted(() => {
  state.fetchAllPlantCultivar();
});
</script>

<template>
  <header>
    <h3>{{ constants.TITLE_PLANT_CULTIVAR }}</h3>
    <base-button @click="state.showPlantCultivarForm">{{ constants.BTN_CREATE }}</base-button>
  </header>

  <main class="content-container">
    <!-- Loading State -->
    <div v-if="state.isLoadingPlantCultivarList" class="text-center p-8">
      <span class="loading loading-lg loading-spinner text-primary"></span>
    </div>

    <!-- Error State -->
    <div v-else-if="state.getPlantCultivarListError" class="p-4 my-4 text-red-700 bg-red-100 rounded-lg">
      <p><strong>Error:</strong> {{ state.getPlantCultivarListError }}</p>
    </div>

    <!-- Data Display -->
    <div v-else-if="state.plantCultivarList && state.plantCultivarList.length > 0" class="cultivar-grid">
      <div v-for="cultivar in state.plantCultivarList" :key="cultivar.id" class="card bg-base-200 shadow-xl">
        <div class="card-body">
          <h2 class="card-title">{{ cultivar.name }}</h2>
          <p class="italic text-neutral-500">{{ cultivar.cultivar }}</p>
          <div class="card-actions justify-end">
            <button class="btn btn-primary btn-sm">View</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="text-center p-8 text-neutral-500">
      <p>No plant cultivar found. Create one to get started!</p>
    </div>
  </main>

  <div id="plant-cultivar-forms">
    <PlantCultivarForm :id="constants.ID_PLANT_CULTIVAR_FORM" />
  </div>
</template>

<style scoped>
header {
  /* Core layout with Flexbox */
  display: flex;
  justify-content: space-between;
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

.content-container {
  padding-top: 8rem; /* Adjust based on header/navbar height */
  padding-left: 1.5rem;
  padding-right: 1.5rem;
}

.cultivar-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 1rem;
}
</style>
