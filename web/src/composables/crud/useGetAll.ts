import { ref } from 'vue';

/**
 * A generic, reusable composable for fetching all items from a resource via a GET request.
 * @template T The expected type of a single item in the returned array.
 * @param {string} apiPath The specific API path for the resource (e.g., '/api/plant-species/').
 * @returns An object with data, isLoading, and error refs, and a `fetchAll` function.
 */
export function useGetAll<T>(apiPath: string) {
  const data = ref<T[] | null>(null);
  const isLoading = ref<boolean>(false);
  const error = ref<string | null>(null);

  // Get the base URL from environment variables
  const baseUrl = import.meta.env.VITE_API_BASE_URL;
  if (!baseUrl) {
    console.error("VITE_API_BASE_URL is not set in your environment variables.");
  }
  const fullApiUrl = `${baseUrl}${apiPath}`;

  /**
   * Performs the API GET request to fetch all items.
   * Populates the data, isLoading, and error refs.
   */
  const fetchAll = async () => {
    isLoading.value = true;
    error.value = null;
    try {
      const response = await fetch(fullApiUrl);
      const jsonResponse = await response.json();

      if (!response.ok) {
        const errorMessage = jsonResponse.error || 'An unknown server error occurred.';
        throw new Error(errorMessage);
      }

      // On success, unpack the 'data' key which should contain the array
      data.value = jsonResponse.data as T[];

    } catch (e: any) {
      console.error(`Failed to fetch items from ${fullApiUrl}:`, e);
      error.value = e.message;
      data.value = null; // Clear data on error
    } finally {
      isLoading.value = false;
    }
  };

  return {
    data,
    isLoading,
    error,
    fetchAll,
  };
}
