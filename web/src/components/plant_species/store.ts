import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

import {
  type PlantSpecies,
  type PlantSpeciesRequest,
  usePlantSpeciesAPI
} from './useAPI';

export default defineStore('usePlantSpeciesStore', () => {
  // --- HTTP composable ---
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

  // --- State ---
  const plantSpeciesFormActive = ref<boolean>(false);

  // --- Computed / Getters ---
  const plantSpeciesFormTitle = computed<string>(() => 'Creating a Plant Species');

  const plantSpeciesIdentifiers = computed<Record<string, string>>(() => {
    if (!plantSpeciesList.value) {
      return {};
    }

    const sortedSpecies = [...plantSpeciesList.value].sort((a, b: PlantSpecies) => {
      return a.name.localeCompare(b.name);
    });

    return sortedSpecies.reduce(
      (acc: Record<string, string>, species: PlantSpecies) => {
        acc[species.name] = species.id;
        return acc;
      },
      {},
    );
  });


  // --- Form lifecycle ---
  const showPlantSpeciesForm = () => {
    plantSpeciesFormActive.value = true;
  };

  const exitPlantSpeciesForm = () => {
    plantSpeciesFormActive.value = false;
    if (createPlantSpeciesError.value) {
      createPlantSpeciesError.value = null;
    }
  };

  // --- HTTP methods ---
  /**
   * 
   * submitNewPlantSpecies makes an HTTP POST Request
   *  for the creation of a new plant species.
   *  Closes the form on success.
   * @param {PlantSpeciesRequest} speciesData The data for the new plant species.
   * @returns {Promise<PlantSpecies | null>} the created plant species
   */
  const submitNewPlantSpecies = async (speciesData: PlantSpeciesRequest): Promise<PlantSpecies | null> => {
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
   * fetchAllPlantSpeciesIfNeeded makes an HTTP GET Request
   *  for all plant species
   * @param {boolean} force override flag to force fetching independent of caching logic
   * @returns {PlantSpecies[]} Plant species records from the database
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
      createPlantSpecies,
      createPlantSpeciesError,
      exitPlantSpeciesForm,
      fetchAllPlantSpecies: fetchAllPlantSpeciesIfNeeded,
      getPlantSpeciesListError,
      isCreatingPlantSpecies,
      isLoadingPlantSpeciesList,
      plantSpeciesFormActive,
      plantSpeciesFormTitle,
      plantSpeciesIdentifiers,
      plantSpeciesList,
      showPlantSpeciesForm,
      submitNewPlantSpecies,
    }
});
