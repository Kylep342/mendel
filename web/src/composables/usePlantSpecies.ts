import { ref } from 'vue';

/**
 * Interface for the data required to create a new PlantSpecies.
 */
interface PlantSpeciesData {
  name: string;
  taxon: string;
}

/**
 * A composable function to handle Plant Species API interactions.
 * It encapsulates the loading and error states for the API calls.
 *
 * @returns An object containing the loading state, error state, and the function to create a species.
 */
export function usePlantSpeciesApi() {
  const isLoading = ref<boolean>(false);
  const error = ref<string | null>(null);

  /**
   * Performs an API POST request to create a new PlantSpecies.
   * @param {PlantSpeciesData} speciesData - The data for the new plant species.
   * @returns {Promise<boolean>} A promise that resolves to true if successful, false otherwise.
   */
  const createPlantSpecies = async (speciesData: PlantSpeciesData): Promise<boolean> => {
    // Set loading state and clear previous errors
    isLoading.value = true;
    error.value = null;

    try {
      // Replace '/api/plant-species' with your actual API endpoint
      const response = await fetch('/api/plant-species', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(speciesData),
      });

      // If the server returns a non-2xx status, handle it as an error
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to create plant species.');
      }

      // Log the success response
      console.log('Successfully created plant species:', await response.json());
      return true; // Indicate success

    } catch (e: any) {
      // Catch any errors and store the message
      console.error(e);
      error.value = e.message;
      return false; // Indicate failure
    } finally {
      // Always turn off loading state
      isLoading.value = false;
    }
  };

  // Expose the state and the function
  return {
    isLoading,
    error,
    createPlantSpecies,
  };
}
