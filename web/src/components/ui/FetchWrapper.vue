<script setup lang="ts">
import type { FetchWrappable } from '@/types/app';
const {
    isFetching,
    error,
    data,
} = defineProps<FetchWrappable>()
</script>

<template>
<div>
    <div v-if="isFetching">
        <slot name="isFetching">
            Loading
        </slot>
    </div>
    <div v-else-if="error">
        <slot name="error">
            Error
        </slot>
    </div>
    <div v-else-if="data">
        <slot name="data">
            Data
        </slot>
    </div>
</div>
</template>