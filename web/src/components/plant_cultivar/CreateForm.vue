<script setup lang="ts">
import { computed, ref } from 'vue';

import constants from '../../constants/constants';
import useMendelCoreStore from '@/stores/core';
import type { PlantCultivarData } from '@/components/plant_cultivar/composables/useAPI';

const state = useMendelCoreStore();

const name = ref<string | null>(null);
const cultivar = ref<string | null>(null);
const species_id = ref<string | null>(null);
const genetics = ref<Record<string, any | null>>({});

const createButtonEnabled = computed<boolean>(
  () => [name.value, cultivar.value, species_id.value].every(
    (input) => input !== null && input !== '',
  ),
);

const clearForm = () => {
  name.value = null;
  cultivar.value = null;
  species_id.value = null;
  genetics.value = {};
};

const exit = () => {
  clearForm();
  state.exitPlantCultivarForm();
};

const createPlantCultivar = () => {
  const data = <PlantCultivarData>{
    name: name.value,
    cultivar: cultivar.value,
    species_id: species_id.value,
    genetics: genetics.value,
  }
  const resp = state.submitNewPlantCultivar(data)
  console.log(`Response was ${resp}`)
  exit();
};
</script>

<template>
  <base-modal @exit="exit">
    <template #header>
      <h2 :class="['pl-4']">
        {{ state.plantCultivarFormTitle }}
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
        <input :id="`${constants.ID_PLANT_CULTIVAR_FORM}-name`" v-model="name"
          :class="['input input-bordered input-secondary w-full max-ws']" type="string" label="Name">
        <div :class="['label']">
          <span :class="['label-text']">Cultivar</span>
        </div>
        <input :id="`${constants.ID_PLANT_CULTIVAR_FORM}-cultivar`" v-model.number="cultivar"
          :class="['input input-bordered input-secondary w-full max-ws']" type="string" label="Cultivar">
        <div :class="['label']">
          <span :class="['label-text']">Plant Species</span>
        </div>
        <select
            :id="`${constants.ID_PLANT_CULTIVAR_FORM}-species-id`"
            v-model="species_id"
            class="select select-bordered w-full max-w-xs"
          >
            <option
              v-for="(id, name) in state.plantSpeciesIdentifiers"
              :key="name"
              :value="id"
            >
              {{ name }}
            </option>
          </select>
          <div :class="['label']">
            <span :class="['label-text']">Genetics</span>
          </div>
          <input :id="`${constants.ID_PLANT_CULTIVAR_FORM}-genetics`" v-model="genetics"
            :class="['input input-bordered input-secondary w-full max-ws']" type="string" label="Genetics">
      </div>
    </template>
    <template #actions>
      <base-button :disabled="!createButtonEnabled" :class="'btn-success'" @click="createPlantCultivar">
        {{ constants.BTN_CREATE }}
      </base-button>
    </template>
  </base-modal>
</template>
