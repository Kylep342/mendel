<script setup lang="ts">
defineProps<{
  id: string,
  bodyClasses?: Array<string>,
}>();
</script>

<template>
  <dialog
    :id="id"
    @close="$emit('exit')"
    :class="['modal', 'modal-bottom', 'sm:modal-middle']"
  >
    <div :class="['modal-box', 'p-0', 'flex', 'flex-col', 'min-w-40', 'max-w-fit']">
      <base-card
        :class="['overflow-hidden']"
        :body-classes="bodyClasses"
      >
        <template #cardTitle>
          <header class="navbar bg-secondary">
            <div class="flex-1">
              <slot name="header" />
            </div>
            <div class="modal-action flex-none mt-0">
              <menu>
                <slot name="headerActions" />
              </menu>
            </div>
          </header>
        </template>
        <template #cardBody>
          <slot name="body" />
        </template>
        <template #cardActions>
          <div class="card-actions justify-end p-4">
            <menu>
              <slot name="actions" />
            </menu>
          </div>
        </template>
      </base-card>
    </div>
  </dialog>
</template>

<style scoped>
.modal {
  background-color: rgba(0, 0, 0, 0.4);
}

.modal-box {
  background-color: transparent; /* The card will provide the background */
  padding: 0;
  border-radius: 0.75rem; /* Match the card's border-radius */
  max-width: 50rem; /* Adjust as needed */
  width: 90%;
}

.modal-box header.navbar {
  background-color: var(--secondary-color);
  color: var(--text-on-secondary);
  padding: 0.75rem 1.5rem;
  font-size: 1.25rem;
  font-weight: 600;
}

.modal-action {
  margin-top: 0;
}
</style>
