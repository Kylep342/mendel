<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  getItemName: () => string,
  pivot: Array<any>,
  isViewedItemId: () => boolean,
  setViewedItemId: () => void,
}>();

const flexBasis = computed(() => props.pivot ? `basis-1/${props.pivot!.length}` : 'basis-1');
const tabStyle = (id) => props.isViewedItemId(id) ? 'btn-secondary' : 'btn-ghost';
</script>

<template>
  <div>
    <div :class="['tabs', 'flex', 'flex-row', 'join', 'w-full', 'flex-grow', 'p-4']">
      <div v-for="item in pivot" :key="item.id" :class="['join-item', flexBasis, 'w-full']">
        <base-button :class="[
          'w-full',
          tabStyle(item.id),
          { 'shadow-lg': isViewedItemId(item.id) },
        ]" @click="setViewedItemId(item.id)">
          {{ getItemName(item.id) }}
        </base-button>
      </div>
    </div>
    <slot name="tabContent" />
  </div>
</template>

<style scoped>
.tabs {
  padding: 1rem;
  gap: 0.5rem;
}

.tabs .btn {
  border-radius: 0.5rem;
  transition: all 0.3s ease;
  border: 1px solid transparent;
}

.tabs .btn-secondary {
  background-color: var(--secondary-color);
  color: var(--text-on-secondary);
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
  transform: translateY(-2px);
}

.tabs .btn-ghost {
  background-color: transparent;
  color: var(--text-on-secondary);
  border-color: #e5e7eb;
}

.tabs .btn-ghost:hover {
  background-color: rgba(0, 0, 0, 0.05);
}
</style>