import { ref, computed } from 'vue';
import { defineStore } from 'pinia';
import { type PlantSpeciesData, usePlantSpeciesAPI } from '@/composables/plantSpecies/usePlantSpecies';

export default defineStore('mendelCore', () => {
  // --- Composables ---
  const {
    // Create
    isCreatingPlantSpecies,
    createPlantSpeciesError,
    createPlantSpecies,
    // Get All
    plantSpeciesList,
    isLoadingPlantSpeciesList,
    getPlantSpeciesListError,
    fetchAllPlantSpecies
  } = usePlantSpeciesAPI();

  // --- Store State ---
  const plantSpeciesFormActive = ref<boolean>(false);

  // --- Getters / Computed ---
  const plantSpeciesFormTitle = computed(() => 'Creating a Plant Species');

  // --- Actions ---
  const showPlantSpeciesForm = () => {
    plantSpeciesFormActive.value = true;
  };

  const exitPlantSpeciesForm = () => {
    plantSpeciesFormActive.value = false;
    if (createPlantSpeciesError.value) {
      createPlantSpeciesError.value = null;
    }
  };

  /**
   * Orchestrates the creation of a new plant species.
   * Closes the form on success.
   * @param {PlantSpeciesData} speciesData The data for the new plant species.
   */
  const submitNewPlantSpecies = async (speciesData: PlantSpeciesData) => {
    const newSpecies = await createPlantSpecies(speciesData);
    if (newSpecies) {
      exitPlantSpeciesForm();
      // Append to existing local cache or create if empty
      if (plantSpeciesList.value) {
        plantSpeciesList.value.push(newSpecies);
      } else {
        plantSpeciesList.value = [newSpecies];
      }
    }
    return newSpecies;
  };


  /**
   * 
   * @param {boolean} force override flag to force fetching independent of caching logic
   * @returns 
   */
  const fetchAllPlantSpeciesIfNeeded = async (forceFetch: boolean=false) => {
    if (!forceFetch) {
      // If the list already has data, don't fetch again.
      if (plantSpeciesList.value && plantSpeciesList.value.length > 0) {
        console.log('Using cached Plant Species list.');
        return;
      }
    }
    // Otherwise, call the actual fetcher from the composable.
    return fetchAllPlantSpecies();
  };

  return {
    // State
    plantSpeciesFormActive,
    isCreatingPlantSpecies,
    createPlantSpeciesError,
    // Get all
    plantSpeciesList,
    isLoadingPlantSpeciesList,
    getPlantSpeciesListError,
    fetchAllPlantSpecies: fetchAllPlantSpeciesIfNeeded,

    // Getters
    plantSpeciesFormTitle,

    // Actions
    showPlantSpeciesForm,
    exitPlantSpeciesForm,
    submitNewPlantSpecies,
  };
});
