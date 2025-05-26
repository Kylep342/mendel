<script setup lang="ts">
import { Button } from '@/types/app';

defineProps<{
  text: string,
  buttons: Array<Button>
}>();
</script>

<template>
  <div :class="['dropdown', 'dropdown-bottom', 'dropdown-end']">
    <base-button>{{ text }}</base-button>
    <ul tabIndex="{0}"
      :class="['dropdown-content', 'menu', 'bg-base-100', 'rounded-box', 'z-[1]', 'w-fit', 'p-2', 'shadow']">
      <li v-for="(button) in buttons" :key="button.text" @click.prevent="button.onClick()">
        <a>{{ button.text }}</a>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.dropdown-content {
  background-color: white;
  border-radius: 0.5rem;
  box-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);
  padding: 0.5rem;
  min-width: 12rem;
}

.menu > li > a {
  padding: 0.5rem 1rem;
  border-radius: 0.25rem;
  transition: background-color 0.2s ease;
}

.menu > li > a:hover {
  background-color: var(--secondary-color);
}
</style>
