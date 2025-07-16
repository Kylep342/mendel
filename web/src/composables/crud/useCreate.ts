import { ref } from 'vue';

import constants from '@/constants/constants'

/**
 * A generic, reusable composable for creating an item via a POST request.
 * It handles the API endpoint, loading state, error handling, and response parsing.
 * @template T The expected type of the data returned upon successful creation.
 * @param {string} apiPath The specific API path for the resource (e.g., '/api/plant-species').
 * @returns An object with isLoading and error refs, and a `create` function.
 */
export function useCreate<T>(apiPath: string) {
  const isLoading = ref<boolean>(false);
  const error = ref<string | null>(null);

  // Get the base URL from environment variables.
  // Make sure you have a .env file in your project root with VITE_API_BASE_URL=http://localhost:8080 (or your server's address)
  const fullApiUrl = `${constants.BASE_URL}${apiPath}`;

  /**
   * Performs the API POST request to create a new item.
   * @template U The type of the payload being sent.
   * @param {U} payload The data for the new item to be created.
   * @returns {Promise<T | null>} A promise that resolves with the created item data, or null on failure.
   */
  const create = async <U>(payload: U): Promise<T | null> => {
    isLoading.value = true;
    error.value = null;

    try {
        const response = await fetch(fullApiUrl, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload),
        });

        const jsonResponse = await response.json();

        if (!response.ok) {
            const errorMessage = jsonResponse.error || 'An unknown server error occurred.';
            throw new Error(errorMessage);
        }

        return jsonResponse.data as T;

    } catch (e: any) {
        console.error('Failed to create item', e);
        error.value = e.message;
        return null;
    } finally {
        isLoading.value = false;
    }
  };

  return {
    isLoading,
    error,
    create,
  };
}
