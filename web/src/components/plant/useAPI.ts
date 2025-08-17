import { useCreate } from '@/composables/crud/useCreate';
import { useGetAll } from '@/composables/crud/useGetAll';
import routes from '@/constants/routes';
import { type JSONable } from '@/types/app';

// Request interface | Front end -> Back end
export interface PlantRequest {
  cultivar_id: string;
  species_id: string;
  seed_id: string;
  pollen_id: string;
  genetics: Record<string, JSONable>;
  labels: Record<string, JSONable>;
}

// Response interface | Back end -> Front end
export interface Plant extends PlantRequest {
  id: string;
  generation: number;
  created_at: string;
  updated_at: string;
}

/**
 * A specific composable for Plant  API interactions.
 */
export function usePlantAPI() {
  const {
    isLoading: isCreatingPlant,
    error: createPlantError,
    create: createItem
  } = useCreate<Plant>(routes.ROUTE_PLANT);

  const createPlant = async (params: PlantRequest): Promise<Plant | null> => {
    return createItem(params);
  };

  const {
    data: plantList,
    isLoading: isLoadingPlantList,
    error: getPlantListError,
    fetchAll: fetchAllPlant
  } = useGetAll<Plant>(routes.ROUTE_PLANT);

  return {
    // Create
    // create : { isCreatingPlant, createPlantError, createPlant },
    isCreatingPlant,
    createPlantError,
    createPlant,
    // Get All
    plantList,
    isLoadingPlantList,
    getPlantListError,
    fetchAllPlant,
  };
}
