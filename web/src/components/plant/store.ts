import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

import { type PlantRequest, usePlantAPI } from './useAPI';

export default defineStore('plantAPI', () => {
    const {
        isCreatingPlant,
        createPlantError,
        createPlant,
        plantList,
        isLoadingPlantList,
        getPlantListError,
        fetchAllPlant,
    } = usePlantAPI();

    const plantFormActive = ref<boolean>(false);

    const plantFormTitle = computed(() => 'Creating a Plant');

    const showPlantForm = () => {
        plantFormActive.value = true;
    };

    const exitPlantForm = () => {
        plantFormActive.value = false;
        if (createPlantError.value) {
            createPlantError.value = null;
        }
    };

    /**
   * Orchestrates the creation of a new plant cultivar.
   * Closes the form on success.
   * @param {PlantCultivarRequest} plantData The data for the new plant cultivar
   */
  const submitNewPlant = async (plantData: PlantRequest) => {
    const newPlant = await createPlant(plantData);
    if (newPlant) {
      exitPlantForm();
      // Append to existing local cache or create if empty
      if (plantList.value) {
        plantList.value.push(newPlant);
      } else {
        plantList.value = [newPlant];
      }
    }
    return newPlant;
  };

/**
   *
   * @param {boolean} force override flag to force fetching independent of caching logic
   * @returns
   */
  const fetchAllPlantIfNeeded = async (forceFetch: boolean=false) => {
    if (!forceFetch) {
      // If the list already has data, don't fetch again.
      if (plantList.value && plantList.value.length > 0) {
        console.log('Using cached Plant list.');
        return;
      }
    }
    // Otherwise, call the actual fetcher from the composable.
    return fetchAllPlant();
  };

  return {
    createPlant,
    createPlantError,
    exitPlantForm,
    fetchAllPlant: fetchAllPlantIfNeeded,
    getPlantListError,
    isCreatingPlant,
    isLoadingPlantList,
    plantFormActive,
    plantFormTitle,
    plantList,
    showPlantForm,
    submitNewPlant,
  }
});
