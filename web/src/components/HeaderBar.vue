<script setup lang="ts">
import { useRouter } from "vue-router";

import type { Button } from "@/types/app";
import routes from "@/constants/routes";
import constants from "@/constants/constants";

// const state = useDebtonateCoreStore();

const router = useRouter();

// const copyStateToClipboard = () => navigator.clipboard.writeText(
//   JSON.stringify(state.exportState())
// );

const buttons: Array<Button> = [
  { text: constants.BTN_ADMIN, onClick: () => router.push(routes.ROUTE_INTERNAL)},
  { text: constants.BTN_PLANT, onClick: () => router.push(routes.ROUTE_PLANT_SPECIES)},
  { text: constants.BTN_PLANT_CULTIVAR, onClick: () => router.push(routes.ROUTE_PLANT_CULTIVAR)},
  { text: constants.BTN_PLANT_SPECIES, onClick: () => router.push(routes.ROUTE_PLANT) },
];
</script>

<template>
  <header>
    <div>
      <!-- <img src="/icon.png"> -->
      <h1>Mendel</h1>
    </div>
    <base-menu :text="constants.BTN_MENU" :buttons="buttons" />
  </header>
</template>

<style scoped>

</style>