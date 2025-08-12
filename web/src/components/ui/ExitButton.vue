<script setup lang="ts">
defineProps<{
  exit: () => void
}>();
</script>

<template>
  <base-button @click="exit" class="close-button">
    <svg
    xmlns="http://www.w3.org/2000/svg"
    class="icon"
    fill="none"
    viewBox="0 0 24 24"
    stroke="currentColor"
    >
    <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M6 18L18 6M6 6l12 12"
    />
    </svg>
  </base-button>
</template>

<style scoped>
.close-button {
  background-color: var(--urgent-color);
}
</style>
