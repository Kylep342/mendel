import { useCreate } from './useCreate'; // Adjust path if needed
import routes from '@/constants/routes'; // Adjust path if needed

/**
 * PlantSpeciesData is the JSON for the create Post request
 */
export interface PlantSpeciesData {
  name: string;
  taxon: string;
}

/**
 * PlantSpecies is the JSON for the create Post response
 */
export interface PlantSpecies extends PlantSpeciesData {
  id: number; // or string, depending on your database
}

/**
 * It uses the generic `useCreate` composable to handle the actual API call.
 */
export function usePlantSpeciesApi() {
  // Instantiate the generic composable with the specific type and API path.
  const {
    isLoading,
    error,
    create: createItem
  } = useCreate<PlantSpecies>(routes.ROUTE_PLANT_SPECIES);

  /**
   * Creates a new PlantSpecies by calling the generic create function.
   * This provides a type-safe wrapper for the specific entity.
   * @param {PlantSpeciesData} speciesData The data for the new plant species.
   * @returns {Promise<PlantSpecies | null>} The created plant species or null on failure.
   */
  const createPlantSpecies = async (speciesData: PlantSpeciesData): Promise<PlantSpecies | null> => {
    return createItem(speciesData);
  };

  return {
    // Expose the state and the specific, typed function
    isLoading,
    error,
    createPlantSpecies,
  };
}
