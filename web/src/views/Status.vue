<script setup lang="ts">
import { useFetch } from '@vueuse/core';
import { withServer } from '@/functions/withServer';
import BaseCard from '@/components/ui/BaseCard.vue'; // Make sure to import your BaseCard

// Fetch the data from your /env endpoint and instruct useFetch to parse it as JSON.
// The `data` ref will be populated with the JSON object.
const { isFetching, error, data: envConfig } = useFetch('/env').json();
</script>

<template>
  <div class="p-4 md:p-8">
    <base-card>
      <template #cardTitle>
        <header class="p-4">
          <h1 class="text-xl font-bold">Server Environment</h1>
        </header>
      </template>

      <template #cardBody>
        <div v-if="isFetching" class="p-8 text-center">
          <p>Loading configuration...</p>
        </div>

        <div v-else-if="error" class="p-8 text-red-600">
          <p class="font-bold">Failed to load environment variables.</p>
          <pre class="mt-2 text-sm bg-red-50 p-2 rounded">Error: {{ error.message }}</pre>
        </div>

        <div v-else-if="envConfig" class="overflow-x-auto">
          <table class="config-table">
            <thead>
              <tr>
                <th>Component</th>
                <th>Settings</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(value, key) in envConfig['data']" :key="key">
                <td>{{ key }}</td>
                <td>{{ value }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </template>
    </base-card>
  </div>
</template>

<style scoped>
.config-table {
  width: 100%;
  border-collapse: collapse; /* For clean lines */
  font-size: 0.9rem;
}

.config-table th,
.config-table td {
  padding: 0.75rem 1rem;
  border-bottom: 1px solid var(--secondary-color);
  text-align: left;
  vertical-align: top;
  white-space: pre-wrap; /* Allows long values to wrap */
  word-break: break-all;
}

.config-table th {
  font-weight: 600;
  background-color: var(--secondary-color);
}

/* Style for the variable names to make them stand out */
.config-table td:first-child {
  font-weight: bold;
}

/* Remove the border from the last row for a cleaner look */
.config-table tbody tr:last-child td {
  border-bottom: none;
}
</style>