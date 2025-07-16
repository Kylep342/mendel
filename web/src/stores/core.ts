import { ref, computed } from 'vue';
import { defineStore } from 'pinia';
import { type PlantSpeciesData, usePlantSpeciesApi } from '@/composables/crud/createPlantSpecies';

export default defineStore('mendelCore', () => {
  // --- Composables ---
  const {
    isLoading: isCreatingPlantSpecies,
    error: createPlantSpeciesError,
    createPlantSpecies: createPlantSpecies
  } = usePlantSpeciesApi();

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
      console.log('A new species was successfully created:', newSpecies);
      // On success, close the form.
      exitPlantSpeciesForm();
      // TODO: add to local state array of species
    }
    // On failure, the form remains open and the error is automatically set in `createPlantSpeciesError`.
  };

  return {
    // State
    plantSpeciesFormActive,
    isCreatingPlantSpecies,
    createPlantSpeciesError,

    // Getters
    plantSpeciesFormTitle,

    // Actions
    showPlantSpeciesForm,
    exitPlantSpeciesForm,
    submitNewPlantSpecies,
  };
});
