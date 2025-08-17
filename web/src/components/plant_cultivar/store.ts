import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

import {
  type PlantCultivar,
  type PlantCultivarRequest,
  usePlantCultivarAPI
} from './useAPI';

export default defineStore('plantCultivarAPI', () => {
  // --- HTTP composable ---
  const {
    // Create
    isCreatingPlantCultivar,
    createPlantCultivarError,
    createPlantCultivar,
    // Get All
    plantCultivarList,
    isLoadingPlantCultivarList,
    getPlantCultivarListError,
    fetchAllPlantCultivar,
  } = usePlantCultivarAPI();

  // --- State ---
  const plantCultivarFormActive = ref<boolean>(false);

  // --- Computed / Getters ---
  const plantCultivarFormTitle = computed<string>(() => 'Creating a Plant Cultivar');

  const plantCultivarIdentifiers = computed<Record<string, string>>(() => {
    if (!plantCultivarList.value) {
      return {};
    }

    const sortedCultivar = [...plantCultivarList.value].sort((a, b: PlantCultivar) => {
      return a.name.localeCompare(b.name);
    });

    return sortedCultivar.reduce(
      (acc: Record<string, string>, cultivar: PlantCultivar) => {
        acc[cultivar.name] = cultivar.id;
        return acc;
      },
      {},
    );
  });

  // --- Form lifecycle ---
  const showPlantCultivarForm = () => {
    plantCultivarFormActive.value = true;
  };

  const exitPlantCultivarForm = () => {
    plantCultivarFormActive.value = false;
    if (createPlantCultivarError.value) {
      createPlantCultivarError.value = null;
    }
  };

  // --- HTTP methods ---
  /**
   * 
   * submitNewPlantCultivar makes an HTTP POST Request
   *  for the creation of a new plant cultivar.
   *  Closes the form on success.
   * @param {PlantCultivarCultivarRequest} plantCultivarData The data for the new plantCultivar cultivar
   * @returns {Promise<PlantCultivar | null>} the created plant cultivar
   */
  const submitNewPlantCultivar = async (plantCultivarData: PlantCultivarRequest): Promise<PlantCultivar | null> => {
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
     * fetchAllPlantCultivarIfNeeded makes an HTTP GET Request
     *  for all plant cultivars
     * @param {boolean} force override flag to force fetching independent of caching logic
     * @returns {PlantCultivar[]} Plant cultivar records in the database
     */
  const fetchAllPlantCultivarIfNeeded = async (forceFetch: boolean = false) => {
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
    plantCultivarIdentifiers,
    plantCultivarList,
    showPlantCultivarForm,
    submitNewPlantCultivar,
  }
});
