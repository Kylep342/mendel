import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

import { type PlantCultivarRequest, usePlantCultivarAPI } from './useAPI';

export default defineStore('plantCultivarAPI', () => {
    const {
        isCreatingPlantCultivar,
        createPlantCultivarError,
        createPlantCultivar,
        plantCultivarList,
        isLoadingPlantCultivarList,
        getPlantCultivarListError,
        fetchAllPlantCultivar,
    } = usePlantCultivarAPI();

    const plantCultivarFormActive = ref<boolean>(false);

    const plantCultivarFormTitle = computed(() => 'Creating a Plant Cultivar');

    const showPlantCultivarForm = () => {
        plantCultivarFormActive.value = true;
    };

    const exitPlantCultivarForm = () => {
        plantCultivarFormActive.value = false;
        if (createPlantCultivarError.value) {
            createPlantCultivarError.value = null;
        }
    };

    /**
   * Orchestrates the creation of a new plantCultivar cultivar.
   * Closes the form on success.
   * @param {PlantCultivarCultivarRequest} plantCultivarData The data for the new plantCultivar cultivar
   */
  const submitNewPlantCultivar = async (plantCultivarData: PlantCultivarRequest) => {
    const newPlantCultivar = await createPlantCultivar(plantCultivarData);
    if (newPlantCultivar) {
      exitPlantCultivarForm();
      // Append to existing local cache or create if empty
      if (plantCultivarList.value) {
        plantCultivarList.value.push(newPlantCultivar);
      } else {
        plantCultivarList.value = [newPlantCultivar];
      }
    }
    return newPlantCultivar;
  };

/**
   *
   * @param {boolean} force override flag to force fetching independent of caching logic
   * @returns
   */
  const fetchAllPlantCultivarIfNeeded = async (forceFetch: boolean=false) => {
    if (!forceFetch) {
      // If the list already has data, don't fetch again.
      if (plantCultivarList.value && plantCultivarList.value.length > 0) {
        console.log('Using cached Plant Cultivar list.');
        return;
      }
    }
    // Otherwise, call the actual fetcher from the composable.
    return fetchAllPlantCultivar();
  };

  return {
    createPlantCultivar,
    createPlantCultivarError,
    exitPlantCultivarForm,
    fetchAllPlantCultivar: fetchAllPlantCultivarIfNeeded,
    getPlantCultivarListError,
    isCreatingPlantCultivar,
    isLoadingPlantCultivarList,
    plantCultivarFormActive,
    plantCultivarFormTitle,
    plantCultivarList,
    showPlantCultivarForm,
    submitNewPlantCultivar,
  }
});
