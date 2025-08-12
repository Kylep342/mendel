import { useCreate } from '@/composables/crud/useCreate';
import { useGetAll } from '@/composables/crud/useGetAll';
import routes from '@/constants/routes';
import { type JSONable } from '@/types/app';

// Request interface | Front end -> Back end
export interface PlantCultivarDTO {
  name: string;
  cultivar: string;
  species_id: string;
  genetics: Record<string, JSONable>;
}

// Response interface | Back end -> Front end
export interface PlantCultivar extends PlantCultivarDTO {
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
  } = useCreate<PlantCultivar>(routes.ROUTE_PLANT_CULTIVAR);

  const createPlantCultivar = async (speciesData: PlantCultivarDTO): Promise<PlantCultivar | null> => {
    return createItem(speciesData);
  };

  // --- Get All Logic ---
  const {
    data: plantCultivarList,
    isLoading: isLoadingPlantCultivarList,
    error: getPlantCultivarListError,
    fetchAll: fetchAllPlantCultivar
  } = useGetAll<PlantCultivar>(routes.ROUTE_PLANT_CULTIVAR);

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
