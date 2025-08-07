import { useCreate } from '@/composables/crud/useCreate';
import { useGetAll } from '@/composables/crud/useGetAll';
import routes from '@/constants/routes';

export interface PlantCultivarData {
  name: string;
  cultivar: string;
  species_id: string;
  genetics: Record<string, any>;
}

export interface PlantCultivar extends PlantCultivarData {
  id: string;
  created_at: string;
  updated_at: string;
}

/**
 * A specific composable for Plant Cultivar API interactions.
 */
export function usePlantCultivarAPI() {
  // --- Create Logic ---
  const {
    isLoading: isCreatingPlantCultivar,
    error: createPlantCultivarError,
    create: createItem
  } = useCreate<PlantCultivar>(routes.ROUTE_PLANT_SPECIES);

  const createPlantCultivar = async (speciesData: PlantCultivarData): Promise<PlantCultivar | null> => {
    return createItem(speciesData);
  };

  // --- Get All Logic ---
  const {
    data: plantCultivarList,
    isLoading: isLoadingPlantCultivarList,
    error: getPlantCultivarListError,
    fetchAll: fetchAllPlantCultivar
  } = useGetAll<PlantCultivar>(routes.ROUTE_PLANT_SPECIES);

  return {
    // Create
    isCreatingPlantCultivar,
    createPlantCultivarError,
    createPlantCultivar,
    // Get All
    plantCultivarList,
    isLoadingPlantCultivarList,
    getPlantCultivarListError,
    fetchAllPlantCultivar,
  };
}
