<template>
  <button :class="['btn']">
    <slot />
  </button>
</template>

<style scoped>
.btn {
  background-color: var(--primary-color);
  color: var(--text-on-primary);
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  font-weight: 600;
  transition: background-color 0.3s ease, transform 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.btn:hover {
  filter: brightness(110%);
}

.btn:active {
  transform: scale(0.98);
  filter: brightness(100%);
}
</style>
