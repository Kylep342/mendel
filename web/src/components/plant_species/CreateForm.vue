<script setup lang="ts">
import { computed, ref } from 'vue';
import constants from '../../constants/constants';
import useMendelCoreStore from '@/stores/core';
import type { PlantSpeciesDTO } from '@/components/plant_species/composables/useAPI';

const state = useMendelCoreStore();

const name = ref<string | null>(null);
const taxon = ref<string | null>(null);

const createButtonEnabled = computed<boolean>(
  () => [name.value, taxon.value].every(
    (input) => input !== null && input !== '',
  ),
);

const clearForm = () => {
  name.value = null;
  taxon.value = null;
};

const exit = () => {
  clearForm();
  state.exitPlantSpeciesForm();
};

const createPlantSpecies = () => {
  const data: PlantSpeciesDTO = {
    name: name.value!,
    taxon: taxon.value!,
  };
  state.submitNewPlantSpecies(data);
};
</script>

<template>
  <base-modal @exit="exit">
    <template #header>
      <div class="form-header">
        <h2>{{ state.plantSpeciesFormTitle }}</h2>
        <exit-button :exit="exit" />
      </div>
    </template>

    <template #body>
      <div class="form-body">
        <div class="form-group">
          <label :for="`${constants.ID_PLANT_SPECIES_FORM}-name`" class="form-label">
            Name
          </label>
          <input
            :id="`${constants.ID_PLANT_SPECIES_FORM}-name`"
            v-model="name"
            type="text"
            class="form-input"
            placeholder="e.g., Capsicum annuum"
          />
        </div>

        <div class="form-group">
          <label :for="`${constants.ID_PLANT_SPECIES_FORM}-taxon`" class="form-label">
            Taxon
          </label>
          <input
            :id="`${constants.ID_PLANT_SPECIES_FORM}-taxon`"
            v-model="taxon"
            type="text"
            class="form-input"
            placeholder="e.g., Solanaceae"
          />
        </div>
      </div>
    </template>

    <template #actions>
      <div class="form-actions">
        <button
          :disabled="!createButtonEnabled"
          @click="createPlantSpecies"
          class="submit-button"
        >
          {{ constants.BTN_CREATE }}
        </button>
      </div>
    </template>
  </base-modal>
</template>

<style scoped>
.form-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  border-bottom: 1px solid #e5e7eb;
}

.form-header h2 {
  font-size: 1.25rem;
  font-weight: 600;
  color: #111827;
}

.form-body {
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-actions {
  padding: 1rem 1.5rem;
  display: flex;
  justify-content: flex-end;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  margin-bottom: 0.25rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
}

.form-input {
  display: block;
  width: 100%;
  padding: 0.5rem 0.75rem;
  font-size: 1rem;
  color: #111827;
  background-color: #fff;
  border: 1px solid #d1d5db;
  border-radius: 0.375rem;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  transition: border-color 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.form-input::placeholder {
  color: #9ca3af;
}

.form-input:focus {
  outline: none;
  border-color: #22c55e;
  box-shadow: 0 0 0 3px rgba(34, 197, 94, 0.2);
}

/* Buttons */
.close-button {
  background: none;
  border: none;
  padding: 0.25rem;
  cursor: pointer;
  color: #6b7280;
  border-radius: 50%;
  transition: background-color 0.2s, color 0.2s;
}

.close-button:hover {
  background-color: #f3f4f6;
  color: #111827;
}

.close-button .icon {
  width: 1.5rem;
  height: 1.5rem;
}

.submit-button {
  display: inline-flex;
  justify-content: center;
  padding: 0.5rem 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: white;
  background-color: #22c55e;
  border: 1px solid transparent;
  border-radius: 0.375rem;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  cursor: pointer;
  transition: background-color 0.2s;
}

.submit-button:hover {
  background-color: #16a34a;
}

.submit-button:disabled {
  background-color: #9ca3af;
  cursor: not-allowed;
  opacity: 0.7;
}
</style>
