import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

import {
  type Plant,
  type PlantRequest,
  usePlantAPI
} from './useAPI';

export default defineStore('plantAPI', () => {
  // --- HTTP composable ---
  const {
    // Create
    isCreatingPlant,
    createPlantError,
    createPlant,
    // Get All
    plantList,
    isLoadingPlantList,
    getPlantListError,
    fetchAllPlant,
  } = usePlantAPI();

  // --- State ---
  const plantFormActive = ref<boolean>(false);

  // --- Computed / Getters
  const plantFormTitle = computed(() => 'Creating a Plant');

  // --- Form lifecycle
  const showPlantForm = () => {
    plantFormActive.value = true;
  };

  const exitPlantForm = () => {
    plantFormActive.value = false;
    if (createPlantError.value) {
      createPlantError.value = null;
    }
  };

  // --- HTTP methods ---
  /**
   * 
   * submitNewPlant makes an HTTP POST Request
   *  for the creation of a new plant.
   *  Closes the form on success.
   * @param {PlantRequest} plantData The data for the new plant 
   * @returns {Promise<Plant | null>} the created plant
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
     * fetchAllPlantIfNeeded makes an HTTP GET Request
     *  for all plants
     * @param {boolean} force override flag to force fetching independent of caching logic
     * @returns {Plant{}} Plant records in the database
     */
  const fetchAllPlantIfNeeded = async (forceFetch: boolean = false) => {
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
