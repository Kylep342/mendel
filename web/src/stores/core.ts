import { ref, computed } from 'vue';
import { defineStore } from 'pinia';
import { type PlantDTO, usePlantAPI } from '@/components/plant/composables/useAPI';
import { type PlantCultivarDTO, usePlantCultivarAPI } from '@/components/plant_cultivar/composables/useAPI';
import { type PlantSpeciesDTO, usePlantSpeciesAPI } from '@/components/plant_species/composables/useAPI';

export default defineStore('mendelCore', () => {
  // --- Composables ---
  const {
    // Create
    isCreatingPlant,
    createPlantError,
    createPlant,
    // Get All
    plantList,
    isLoadingPlantList,
    getPlantListError,
    fetchAllPlant
  } = usePlantAPI();

  const {
    // Create
    isCreatingPlantCultivar,
    createPlantCultivarError,
    createPlantCultivar,
    // Get All
    plantCultivarList,
    isLoadingPlantCultivarList,
    getPlantCultivarListError,
    fetchAllPlantCultivar
  } = usePlantCultivarAPI();

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
  const plantFormActive = ref<boolean>(false);
  const plantCultivarFormActive = ref<boolean>(false);
  const plantSpeciesFormActive = ref<boolean>(false);

  // --- Getters / Computed ---
  const plantFormTitle = computed(() => 'Creating a Plant');
  const plantCultivarFormTitle = computed(() => 'Creating a Plant Cultivar');
  const plantSpeciesFormTitle = computed(() => 'Creating a Plant Species');
  const plantSpeciesIdentifiers = computed(() => {
    if (!plantSpeciesList.value) {
      return {};
    }

    const sortedSpecies = [...plantSpeciesList.value].sort((a, b) => {
      return a.name.localeCompare(b.name);
    });

    return sortedSpecies.reduce((acc, species) => {
      acc[species.name] = species.id;
      return acc;
    }, {} as Record<string, string>);
  });

  // --- Actions ---

  const showPlantForm = () => {
    plantFormActive.value = true;
  };
  const showPlantCultivarForm = () => {
    plantCultivarFormActive.value = true;
  };
  const showPlantSpeciesForm = () => {
    plantSpeciesFormActive.value = true;
  };

  const exitPlantForm = () => {
    plantFormActive.value = false;
    if (createPlantError.value) {
      createPlantError.value = null;
    }
  };
  const exitPlantCultivarForm = () => {
    plantCultivarFormActive.value = false;
    if (createPlantCultivarError.value) {
      createPlantCultivarError.value = null;
    }
  };
  const exitPlantSpeciesForm = () => {
    plantSpeciesFormActive.value = false;
    if (createPlantSpeciesError.value) {
      createPlantSpeciesError.value = null;
    }
  };

  /**
   * Orchestrates the creation of a new plant cultivar.
   * Closes the form on success.
   * @param {PlantCultivarDTO} plantData The data for the new plant cultivar
   */
  const submitNewPlant = async (plantData: PlantDTO) => {
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
   * Orchestrates the creation of a new plant cultivar.
   * Closes the form on success.
   * @param {PlantCultivarDTO} cultivarData The data for the new plant cultivar
   */
  const submitNewPlantCultivar = async (cultivarData: PlantCultivarDTO) => {
    const newCultivar = await createPlantCultivar(cultivarData);
    if (newCultivar) {
      exitPlantCultivarForm();
      // Append to existing local cache or create if empty
      if (plantCultivarList.value) {
        plantCultivarList.value.push(newCultivar);
      } else {
        plantCultivarList.value = [newCultivar];
      }
    }
    return newCultivar;
  };

  /**
   * Orchestrates the creation of a new plant species.
   * Closes the form on success.
   * @param {PlantSpeciesDTO} speciesData The data for the new plant species.
   */
  const submitNewPlantSpecies = async (speciesData: PlantSpeciesDTO) => {
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
    plantFormActive,
    isCreatingPlant,
    createPlant,
    plantCultivarFormActive,
    isCreatingPlantCultivar,
    createPlantCultivarError,
    plantSpeciesFormActive,
    isCreatingPlantSpecies,
    createPlantSpeciesError,
    // Get all
    plantList,
    isLoadingPlantList,
    getPlantListError,
    fetchAllPlant: fetchAllPlantIfNeeded,
    plantSpeciesIdentifiers,
    plantSpeciesList,
    isLoadingPlantSpeciesList,
    getPlantSpeciesListError,
    fetchAllPlantSpecies: fetchAllPlantSpeciesIfNeeded,
    // plantCultivarIdentifiers,
    plantCultivarList,
    isLoadingPlantCultivarList,
    getPlantCultivarListError,
    fetchAllPlantCultivar: fetchAllPlantCultivarIfNeeded,

    // Getters
    plantFormTitle,
    plantCultivarFormTitle,
    plantSpeciesFormTitle,

    // Actions
    showPlantForm,
    exitPlantForm,
    submitNewPlant,
    showPlantCultivarForm,
    exitPlantCultivarForm,
    submitNewPlantCultivar,
    showPlantSpeciesForm,
    exitPlantSpeciesForm,
    submitNewPlantSpecies,
  };
});
