import { useCreate } from '@/composables/crud/useCreate';
import { useGetAll } from '@/composables/crud/useGetAll';
import routes from '@/constants/routes';

// Request interface | Front end -> Back end
export interface PlantSpeciesRequest {
  name: string;
  taxon: string;
}

// Response interface | Back end -> Front end
export interface PlantSpecies extends PlantSpeciesRequest {
  id: string;
  created_at: string;
  updated_at: string;
}

/**
 * A specific composable for Plant Species API interactions.
 */
export function usePlantSpeciesAPI() {
  const {
    isLoading: isCreatingPlantSpecies,
    error: createPlantSpeciesError,
    create: createItem
  } = useCreate<PlantSpecies>(routes.ROUTE_PLANT_SPECIES);

  const createPlantSpecies = async (params: PlantSpeciesRequest): Promise<PlantSpecies | null> => {
    return createItem(params);
  };

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
