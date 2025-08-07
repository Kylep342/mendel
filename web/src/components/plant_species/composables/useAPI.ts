import { useCreate } from '@/composables/crud/useCreate';
import { useGetAll } from '@/composables/crud/useGetAll';
import routes from '@/constants/routes';

// Interfaces remain the same
export interface PlantSpeciesData {
  name: string;
  taxon: string;
}

export interface PlantSpecies extends PlantSpeciesData {
  id: string; // Assuming UUID from your screenshot
  created_at: string;
  updated_at: string;
}

/**
 * A specific composable for Plant Species API interactions.
 */
export function usePlantSpeciesAPI() {
  // --- Create Logic (existing) ---
  const {
    isLoading: isCreatingPlantSpecies,
    error: createPlantSpeciesError,
    create: createItem
  } = useCreate<PlantSpecies>(routes.ROUTE_PLANT_SPECIES);

  const createPlantSpecies = async (speciesData: PlantSpeciesData): Promise<PlantSpecies | null> => {
    return createItem(speciesData);
  };

  // --- Get All Logic (new) ---
  const {
    data: plantSpeciesList,
    isLoading: isLoadingPlantSpeciesList,
    error: getPlantSpeciesListError,
    fetchAll: fetchAllPlantSpecies
  } = useGetAll<PlantSpecies>(routes.ROUTE_PLANT_SPECIES);

  return {
    // Create
    isCreatingPlantSpecies,
    createPlantSpeciesError,
    createPlantSpecies,
    // Get All
    plantSpeciesList,
    isLoadingPlantSpeciesList,
    getPlantSpeciesListError,
    fetchAllPlantSpecies,
  };
}
