<script setup lang="ts">
import { formatCurrency } from '@/tools/currency';
import { computed } from 'vue'; 

defineProps<{
  date: Date
}>();

const netWorth = computed(() => {
  return 0;
});

const color = computed(() => {
  if (netWorth.value > 0) return 'text-green-700';
  if (netWorth.value == 0) return 'text-current';
  return 'text-red-700';
})

</script>

<template>
  <div class="flex flex-grow items-center border rounded-md border-black/30">
    <div class="flex justify-center px-4 border-r border-black/30">
      NET WORTH
    </div>

    <div class="flex justify-center flex-grow">
      <span class="font-black" :class="color">
        {{ formatCurrency(netWorth, true) }}
      </span>
    </div>
  </div>
</template>