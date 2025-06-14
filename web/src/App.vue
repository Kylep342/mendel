<script setup lang="ts">
import { RouterView } from 'vue-router'
import HeaderBar from './components/HeaderBar.vue'
import FooterBar from './components/FooterBar.vue'
</script>

<template>
  <HeaderBar />
  <main class="content-wrapper">
    <RouterView />
  </main>
  <FooterBar />
</template>

<style scoped>
.content-wrapper {
  /* This padding-top should be the same as the height of your header */
  padding-top: 4rem;
  padding-left: 1.5rem;
  padding-right: 1.5rem;
}
</style>