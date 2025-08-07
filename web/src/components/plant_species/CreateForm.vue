<script setup lang="ts">
import { computed, ref } from 'vue';

import constants from '../../constants/constants';
import useMendelCoreStore from '@/stores/core';
import type { PlantSpeciesData } from '@/components/plant_species/composables/useAPI';

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
  const data = <PlantSpeciesData>{
    name: name.value,
    taxon: taxon.value,
  }
  const resp = state.submitNewPlantSpecies(data)
  console.log(`Response was ${resp}`)
  exit();
};
</script>

<template>
  <base-modal @exit="exit">
    <template #header>
      <h2 :class="['pl-4']">
        {{ state.plantSpeciesFormTitle }}
      </h2>
    </template>
    <template #headerActions>
      <base-button :class="['btn btn-circle btn-ghost']" @click="exit">
        x
      </base-button>
    </template>
    <template #body>
      <div :class="['formInputs']">
        <div :class="['label']">
          <span :class="['label-text']">Name</span>
        </div>
        <input :id="`${constants.ID_PLANT_SPECIES_FORM}-name`" v-model="name"
          :class="['input input-bordered input-secondary w-full max-ws']" type="string" label="Name">
        <div :class="['label']">
          <span :class="['label-text']">Taxon</span>
        </div>
        <input :id="`${constants.ID_PLANT_SPECIES_FORM}-taxon`" v-model.number="taxon"
          :class="['input input-bordered input-secondary w-full max-ws']" type="string" label="Taxon">
      </div>
    </template>
    <template #actions>
      <base-button :disabled="!createButtonEnabled" :class="'btn-success'" @click="createPlantSpecies">
        {{ constants.BTN_CREATE }}
      </base-button>
    </template>
  </base-modal>
</template>
