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
    class="modal-container"
  >
    <div class="modal-content">
      <base-card>
        <template #cardTitle>
          <header class="modal-header">
            <div class="header-content">
              <slot name="header" />
            </div>
            <div class="header-actions">
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
          <div class="modal-footer">
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

.modal-container {
  background-color: transparent;
  padding: 0;
  border: none;
}

.modal-content {
  background-color: white;
  padding: 0;
  border-radius: 0.75rem;
  max-width: 50rem;
  width: 90%;
  overflow: hidden;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: var(--secondary-color, #f3f4f6);
  color: var(--text-on-secondary, #111827);
  padding: 0.75rem 1.5rem;
  font-size: 1.25rem;
  font-weight: 600;
}

.header-content {
  flex: 1;
}

.header-actions {
  flex-shrink: 0;
}

.header-actions menu,
.modal-footer menu {
  margin: 0;
  padding: 0;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  padding: 1rem 1.5rem;
  border-top: 1px solid #e5e7eb;
}
</style>
