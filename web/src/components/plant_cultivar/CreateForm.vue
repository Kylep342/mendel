<script setup lang="ts">
import { computed, ref } from 'vue';

import constants from '../../constants/constants';
import usePlantCultivarStore from './store';
import usePlantSpeciesStore from '@/components/plant_species/store';
import type { PlantCultivarRequest } from '@/components/plant_cultivar/useAPI';
import { type JSONable } from '@/types/app';

const species = usePlantSpeciesStore();
const state = usePlantCultivarStore();

const name = ref<string | null>(null);
const cultivar = ref<string | null>(null);
const species_id = ref<string | null>(null);
const genetics = ref<Record<string, JSONable> | null>(null);

const createButtonEnabled = computed<boolean>(
  () => [name.value, cultivar.value, species_id.value].every(
    (input) => input !== null && input !== '',
  ),
);

const clearForm = () => {
  name.value = null;
  cultivar.value = null;
  species_id.value = null;
  genetics.value = null;
};

const exit = () => {
  clearForm();
  state.exitPlantCultivarForm();
};

const createPlantCultivar = () => {
  const data = <PlantCultivarRequest>{
    name: name.value!,
    cultivar: cultivar.value!,
    species_id: species_id.value!,
    genetics: genetics.value! || {},
  }
  state.submitNewPlantCultivar(data)
};
</script>

<template>
  <base-modal class="modal-panel" @exit="exit">
    <template #header>
      <div class="form-header">
        <h2>{{ state.plantCultivarFormTitle }}</h2>
        <exit-button :exit="exit" />
      </div>
    </template>

    <template #body>
      <div class="form-body">
        <div class="form-group">
          <label :for="`${constants.ID_PLANT_CULTIVAR_FORM}-name`" class="form-label">Name</label>
          <input
            :id="`${constants.ID_PLANT_CULTIVAR_FORM}-name`"
            v-model="name"
            type="text"
            class="form-input"
            placeholder="e.g., Jalapeño"
          />
        </div>

        <div class="form-group">
          <label :for="`${constants.ID_PLANT_CULTIVAR_FORM}-cultivar`" class="form-label">Cultivar</label>
          <input
            :id="`${constants.ID_PLANT_CULTIVAR_FORM}-cultivar`"
            v-model="cultivar"
            type="text"
            class="form-input"
            placeholder="e.g., Capsicum annuum 'Jalapeño'"
          />
        </div>

        <div class="form-group">
          <label :for="`${constants.ID_PLANT_CULTIVAR_FORM}-species-id`" class="form-label">Plant Species</label>
          <select
            :id="`${constants.ID_PLANT_CULTIVAR_FORM}-species-id`"
            v-model="species_id"
            class="form-select"
          >
            <option disabled :value="null">Please select one</option>
            <option v-for="(id, name) in species.plantSpeciesIdentifiers" :key="name" :value="id">
              {{ name }}
            </option>
          </select>
        </div>

        <div class="form-group">
          <label :for="`${constants.ID_PLANT_CULTIVAR_FORM}-genetics`" class="form-label">Genetics (JSON)</label>
          <input
            :id="`${constants.ID_PLANT_CULTIVAR_FORM}-genetics`"
            v-model="genetics"
            type="text"
            class="form-input"
            placeholder='e.g., {"color": "red"}'
          />
        </div>
      </div>
    </template>

    <template #actions>
      <div class="form-actions">
        <button :disabled="!createButtonEnabled" @click="createPlantCultivar" class="submit-button">
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
  background-color: #f9fafb;
  border-top: 1px solid #e5e7eb;
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
