<script setup lang="ts">
defineProps<{ bodyClasses?: Array<string> }>();
</script>

<template>
  <div :class="['card']">
    <slot name="cardTitle" />
    <div :class="['card-body', 'p-4', bodyClasses]">
      <slot name="cardBody" />
    </div>
    <slot name="cardActions" />
  </div>
</template>

<style scoped>
.card {
  background-color: white;
  border-radius: 0.75rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.card-body {
  padding: 1rem;
}
</style>
